package compiler

import (
	"fmt"
)

// handlePrint handles the print() function for debugging output
func (w *ASTWalker) handlePrint(args []interface{}) {
	for i, arg := range args {
		if i > 0 {
			fmt.Print(" ")
		}
		w.printValue(arg)
	}
	fmt.Println()
}

// handleOutput handles the output() function with formatted output
func (w *ASTWalker) handleOutput(args []interface{}) {
	if len(args) == 0 {
		return
	}

	// First argument is the label/name
	if len(args) >= 1 {
		label := w.extractStringValue(args[0])
		fmt.Printf("\033[1;36m[OUTPUT]\033[0m %s", label)

		if len(args) >= 2 {
			fmt.Print(" = ")
			w.printValue(args[1])
		}
		fmt.Println()
	}
}

// printValue prints a value with proper formatting
func (w *ASTWalker) printValue(value interface{}) {
	switch v := value.(type) {
	case string:
		if w.variables != nil {
			if resolved, exists := w.variables[v]; exists {
				w.printValue(resolved)
				return
			}
		}
		fmt.Printf("\033[32m\"%s\"\033[0m", v)
	case float64:
		fmt.Printf("\033[33m%v\033[0m", v)
	case bool:
		fmt.Printf("\033[35m%v\033[0m", v)
	case map[string]interface{}:
		fmt.Print("{\n")
		for key, val := range v {
			fmt.Printf("  \033[34m%s\033[0m: ", key)
			w.printValue(val)
			fmt.Println(",")
		}
		fmt.Print("}")
	case []interface{}:
		fmt.Print("[")
		for i, item := range v {
			if i > 0 {
				fmt.Print(", ")
			}
			w.printValue(item)
		}
		fmt.Print("]")
	default:
		fmt.Printf("%v", v)
	}
}
