package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
)

func (e *Engine) Plan(ctx context.Context, filename string) error {
	fmt.Println("Planning infrastructure changes...")

	program, err := e.compiler.CompileFile(filename)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	if err := e.loadRequiredPlugins(ctx, program); err != nil {
		return fmt.Errorf("failed to load plugins: %w", err)
	}

	currentState, err := e.stateManager.LoadState()
	if err != nil {
		fmt.Println("No existing state found, will create new infrastructure")
		currentState = &state.State{Resources: make(map[string]*state.ResourceState)}
	}

	changes := e.calculateChanges(program, currentState)

	e.displayPlan(changes)

	return nil
}

func (e *Engine) displayPlan(changes *PlanChanges) {
	headerColor.Println("\nPlan Summary:")

	if len(changes.Create) > 0 {
		createColor.Printf("\nResources to create (%d):\n", len(changes.Create))
		for _, resource := range changes.Create {
			createColor.Printf("  + %s ", resource.Name)
			fmt.Printf("(%s)\n", resource.Type)
		}
	}

	if len(changes.Update) > 0 {
		updateColor.Printf("\nResources to update (%d):\n", len(changes.Update))
		for _, resource := range changes.Update {
			updateColor.Printf("  ~ %s ", resource.Name)
			fmt.Printf("(%s)\n", resource.Type)
		}
	}

	if len(changes.Delete) > 0 {
		deleteColor.Printf("\nResources to delete (%d):\n", len(changes.Delete))
		for _, resource := range changes.Delete {
			deleteColor.Printf("  - %s ", resource.Name)
			fmt.Printf("(%s)\n", resource.Type)
		}
	}

	if len(changes.Create) == 0 && len(changes.Update) == 0 && len(changes.Delete) == 0 {
		infoColor.Println("\nNo changes. Infrastructure is up-to-date.")
	}
}
