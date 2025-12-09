package engine

import (
	"context"
	"fmt"
)

// Initialize initializes the engine and discovers plugins
func (e *Engine) Initialize(ctx context.Context) error {
	if err := e.pluginManager.DiscoverPlugins(); err != nil {
		return fmt.Errorf("failed to discover plugins: %w", err)
	}

	fmt.Printf("Discovered plugins: %v\n", e.pluginManager.ListPlugins())
	return nil
}

// Shutdown gracefully shuts down the engine
func (e *Engine) Shutdown() error {
	return e.pluginManager.ShutdownAll()
}

// ListPlugins returns available plugins
func (e *Engine) ListPlugins() []string {
	return e.pluginManager.ListPlugins()
}
