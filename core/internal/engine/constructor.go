package engine

import (
	"os"
	"path/filepath"

	"github.com/tblang/core/internal/compiler"
	"github.com/tblang/core/internal/state"
)

// New creates a new Engine instance
func New() *Engine {
	workingDir, _ := os.Getwd()

	pluginDirs := []string{
		"/usr/local/lib/tblang/plugins",                 // Manual install
		"/opt/homebrew/opt/tblang/lib/tblang/plugins",   // Homebrew (Apple Silicon)
		"/usr/local/opt/tblang/lib/tblang/plugins",      // Homebrew (Intel)
		filepath.Join(workingDir, ".tblang", "plugins"), // Local project
	}

	var pluginDir string
	for _, dir := range pluginDirs {
		if _, err := os.Stat(dir); err == nil {
			pluginDir = dir
			break
		}
	}

	// Default to local if none found
	if pluginDir == "" {
		pluginDir = filepath.Join(workingDir, ".tblang", "plugins")
	}

	return &Engine{
		compiler:      compiler.New(),
		stateManager:  state.NewManager(filepath.Join(workingDir, ".tblang")),
		pluginManager: NewPluginManager(pluginDir),
		workingDir:    workingDir,
	}
}
