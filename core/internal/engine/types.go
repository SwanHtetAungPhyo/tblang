package engine

import (
	"github.com/fatih/color"
	"github.com/tblang/core/internal/compiler"
	"github.com/tblang/core/internal/state"
)

type Engine struct {
	compiler      *compiler.Compiler
	stateManager  *state.Manager
	pluginManager *PluginManager
	workingDir    string
}

type PlanChanges struct {
	Create []*state.ResourceState
	Update []*state.ResourceState
	Delete []*state.ResourceState
}

var (

	successColor = color.New(color.FgGreen, color.Bold)
	errorColor   = color.New(color.FgRed, color.Bold)
	warningColor = color.New(color.FgYellow, color.Bold)
	infoColor    = color.New(color.FgCyan, color.Bold)
	headerColor  = color.New(color.FgMagenta, color.Bold)
	createColor  = color.New(color.FgGreen)
	updateColor  = color.New(color.FgYellow)
	deleteColor  = color.New(color.FgRed)
)
