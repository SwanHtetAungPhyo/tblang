package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
	"github.com/tblang/core/pkg/plugin"
)

// createResourceWithPlugin creates a resource using the plugin
func (e *Engine) createResourceWithPlugin(ctx context.Context, resource *state.ResourceState) (interface{}, error) {
	// Get the AWS plugin
	pluginInstance, err := e.pluginManager.GetPlugin("aws")
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS plugin: %w", err)
	}

	// Resolve resource references in attributes
	resolvedAttrs := e.resolveResourceReferences(resource.Attributes)

	// Call ApplyResourceChange to create the resource
	req := &plugin.ApplyResourceChangeRequest{
		TypeName:     resource.Type,
		PriorState:   nil, // nil for new resources
		PlannedState: resolvedAttrs,
		Config:       resolvedAttrs,
	}

	resp, err := pluginInstance.Client.ApplyResourceChange(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("plugin error: %w", err)
	}

	// Check for diagnostics
	if len(resp.Diagnostics) > 0 {
		for _, diag := range resp.Diagnostics {
			if diag.Severity == "error" {
				return nil, fmt.Errorf("%s: %s", diag.Summary, diag.Detail)
			} else if diag.Severity == "warning" {
				warningColor.Printf("  ⚠ Warning: %s\n", diag.Summary)
			}
		}
	}

	return resp.NewState, nil
}

// destroyResourceWithPlugin destroys a resource using the plugin
func (e *Engine) destroyResourceWithPlugin(ctx context.Context, resource *state.ResourceState) error {
	// Get the AWS plugin
	pluginInstance, err := e.pluginManager.GetPlugin("aws")
	if err != nil {
		return fmt.Errorf("failed to get AWS plugin: %w", err)
	}

	// Import the plugin package types
	req := &plugin.ApplyResourceChangeRequest{
		TypeName:     resource.Type,
		PriorState:   resource.Attributes,
		PlannedState: nil, // nil indicates destroy
		Config:       resource.Attributes,
	}

	resp, err := pluginInstance.Client.ApplyResourceChange(ctx, req)
	if err != nil {
		return fmt.Errorf("plugin error: %w", err)
	}

	// Check for diagnostics
	if len(resp.Diagnostics) > 0 {
		for _, diag := range resp.Diagnostics {
			if diag.Severity == "error" {
				return fmt.Errorf("%s: %s", diag.Summary, diag.Detail)
			}
		}
	}

	return nil
}
