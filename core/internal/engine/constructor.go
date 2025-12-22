package engine

import (
	"os"
	"path/filepath"

	"github.com/tblang/core/internal/compiler"
	"github.com/tblang/core/internal/state"
)

func New() *Engine {
	workingDir, _ := os.Getwd()

	pluginDirs := []string{
		"/usr/local/lib/tblang/plugins",
		"/opt/homebrew/opt/tblang/lib/tblang/plugins",
		"/usr/local/opt/tblang/lib/tblang/plugins",
		filepath.Join(workingDir, ".tblang", "plugins"),
	}

	var pluginDir string
	for _, dir := range pluginDirs {
		if _, err := os.Stat(dir); err == nil {
			pluginDir = dir
			break
		}
	}

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
