package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
)

// Apply creates/updates the infrastructure
func (e *Engine) Apply(ctx context.Context, filename string) error {
	infoColor.Println("Applying infrastructure changes...")

	program, err := e.compiler.CompileFile(filename)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	if err := e.loadAndConfigurePlugins(ctx, program); err != nil {
		return fmt.Errorf("failed to load plugins: %w", err)
	}

	currentState, err := e.stateManager.LoadState()
	if err != nil {
		currentState = &state.State{Resources: make(map[string]*state.ResourceState)}
	}

	changes := e.calculateChanges(program, currentState)

	e.displayPlan(changes)

	fmt.Print("\nDo you want to perform these actions? (yes/no): ")
	var response string
	fmt.Scanln(&response)

	if response != "yes" && response != "y" {
		warningColor.Println("Apply cancelled.")
		return nil
	}

	if err := e.applyChanges(ctx, changes, currentState); err != nil {
		return fmt.Errorf("apply failed: %w", err)
	}

	successColor.Println("\nApply complete!")
	return nil
}

// applyChanges applies the planned changes
func (e *Engine) applyChanges(ctx context.Context, changes *PlanChanges, currentState *state.State) error {
	// Apply creates
	for _, resource := range changes.Create {
		resourceColor := e.getResourceColor(resource.Type)
		resourceColor.Printf("\nCreating %s (%s)...\n", resource.Name, resource.Type)

		// Use plugin to create resource
		newState, err := e.createResourceWithPlugin(ctx, resource)
		if err != nil {
			errorColor.Printf("  ✗ Failed to create %s: %v\n", resource.Name, err)
			return fmt.Errorf("failed to create %s: %w", resource.Name, err)
		}

		// Update resource with new state from plugin
		if newState != nil {
			if stateMap, ok := newState.(map[string]interface{}); ok {
				resource.Attributes = stateMap
			}
		}

		// Update state
		resource.Status = "created"
		currentState.Resources[resource.Name] = resource
		if err := e.stateManager.SaveState(currentState); err != nil {
			return fmt.Errorf("failed to save state: %w", err)
		}

		successColor.Printf("  ✓ Created %s (%s)\n", resource.Name, resource.Type)
	}

	// Apply deletes
	for _, resource := range changes.Delete {
		warningColor.Printf("\nDeleting %s (%s)...\n", resource.Name, resource.Type)

		// Use plugin to delete resource
		if err := e.destroyResourceWithPlugin(ctx, resource); err != nil {
			errorColor.Printf("  ✗ Failed to delete %s: %v\n", resource.Name, err)
		} else {
			successColor.Printf("  ✓ Deleted %s (%s)\n", resource.Name, resource.Type)
		}

		// Remove from state
		delete(currentState.Resources, resource.Name)
		if err := e.stateManager.SaveState(currentState); err != nil {
			return fmt.Errorf("failed to save state: %w", err)
		}
	}

	return nil
}
