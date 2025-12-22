package engine

import (
	"context"
	"fmt"
)

func (e *Engine) Initialize(ctx context.Context) error {
	if err := e.pluginManager.DiscoverPlugins(); err != nil {
		return fmt.Errorf("failed to discover plugins: %w", err)
	}

	fmt.Printf("Discovered plugins: %v\n", e.pluginManager.ListPlugins())
	return nil
}

func (e *Engine) Shutdown() error {
	return e.pluginManager.ShutdownAll()
}

func (e *Engine) ListPlugins() []string {
	return e.pluginManager.ListPlugins()
}
