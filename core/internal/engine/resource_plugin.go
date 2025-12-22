package engine

import (
	"context"
	"fmt"

	"github.com/tblang/core/internal/state"
	"github.com/tblang/core/pkg/plugin"
)

func (e *Engine) createResourceWithPlugin(ctx context.Context, resource *state.ResourceState) (interface{}, error) {

	pluginInstance, err := e.pluginManager.GetPlugin("aws")
	if err != nil {
		return nil, fmt.Errorf("failed to get AWS plugin: %w", err)
	}

	resolvedAttrs := e.resolveResourceReferences(resource.Attributes)

	req := &plugin.ApplyResourceChangeRequest{
		TypeName:     resource.Type,
		PriorState:   nil,
		PlannedState: resolvedAttrs,
		Config:       resolvedAttrs,
	}

	resp, err := pluginInstance.Client.ApplyResourceChange(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("plugin error: %w", err)
	}

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

func (e *Engine) destroyResourceWithPlugin(ctx context.Context, resource *state.ResourceState) error {

	pluginInstance, err := e.pluginManager.GetPlugin("aws")
	if err != nil {
		return fmt.Errorf("failed to get AWS plugin: %w", err)
	}

	req := &plugin.ApplyResourceChangeRequest{
		TypeName:     resource.Type,
		PriorState:   resource.Attributes,
		PlannedState: nil,
		Config:       resource.Attributes,
	}

	resp, err := pluginInstance.Client.ApplyResourceChange(ctx, req)
	if err != nil {
		return fmt.Errorf("plugin error: %w", err)
	}

	if len(resp.Diagnostics) > 0 {
		for _, diag := range resp.Diagnostics {
			if diag.Severity == "error" {
				return fmt.Errorf("%s: %s", diag.Summary, diag.Detail)
			}
		}
	}

	return nil
}
