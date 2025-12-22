package engine

import (
	"fmt"
)

func (e *Engine) Show() error {
	fmt.Println("Current infrastructure state:")

	currentState, err := e.stateManager.LoadState()
	if err != nil {
		fmt.Println("No state found")
		return nil
	}

	if len(currentState.Resources) == 0 {
		fmt.Println("No resources found")
		return nil
	}

	for name, resource := range currentState.Resources {
		fmt.Printf("\nResource: %s\n", name)
		fmt.Printf("   Type: %s\n", resource.Type)
		fmt.Printf("   Status: %s\n", resource.Status)
		if len(resource.Attributes) > 0 {
			fmt.Println("   Attributes:")
			for key, value := range resource.Attributes {
				fmt.Printf("     %s: %v\n", key, value)
			}
		}
	}

	return nil
}
