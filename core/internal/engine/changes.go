package engine

import (
	"github.com/tblang/core/internal/compiler"
	"github.com/tblang/core/internal/state"
)

func (e *Engine) calculateChanges(program *compiler.Program, currentState *state.State) *PlanChanges {
	changes := &PlanChanges{
		Create: make([]*state.ResourceState, 0),
		Update: make([]*state.ResourceState, 0),
		Delete: make([]*state.ResourceState, 0),
	}

	for _, resource := range program.Resources {
		if _, exists := currentState.Resources[resource.Name]; !exists {
			changes.Create = append(changes.Create, &state.ResourceState{
				Name:       resource.Name,
				Type:       resource.Type,
				Status:     "planned",
				Attributes: resource.Properties,
			})
		}
	}

	programResources := make(map[string]bool)
	for _, resource := range program.Resources {
		programResources[resource.Name] = true
	}

	for name, resource := range currentState.Resources {
		if !programResources[name] {
			changes.Delete = append(changes.Delete, resource)
		}
	}

	return changes
}
