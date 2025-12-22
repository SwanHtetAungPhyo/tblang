package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/fatih/color"
	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/internal/compiler"
)

func (e *Engine) Graph(ctx context.Context, filename string) error {
	program, err := e.compiler.CompileFile(filename)
	if err != nil {
		return fmt.Errorf("compilation failed: %w", err)
	}

	e.displayVisualGraph(program)
	return nil
}

func (e *Engine) displayVisualGraph(program *compiler.Program) {
	headerColor.Println("\nDependency Graph & Deployment Order:")
	headerColor.Println(strings.Repeat("=", 50))

	dependencies := make(map[string][]string)

	for _, resource := range program.Resources {
		var deps []string
		for _, value := range resource.Properties {
			if refs := e.findResourceReferences(value, program.Resources); len(refs) > 0 {
				for _, ref := range refs {
					if ref != resource.Name {
						deps = append(deps, ref)
					}
				}
			}
		}
		dependencies[resource.Name] = e.removeDuplicates(deps)
	}

	fmt.Println()
	for i, resource := range program.Resources {

		resourceColor := e.getResourceColor(resource.Type)
		resourceColor.Printf("[%d] %s", i+1, resource.Name)
		fmt.Printf(" (%s)\n", resource.Type)

		if deps := dependencies[resource.Name]; len(deps) > 0 {
			infoColor.Print("    Dependencies: ")
			for j, dep := range deps {
				if j > 0 {
					fmt.Print(", ")
				}
				warningColor.Print(dep)
			}
			fmt.Println()
		} else {
			infoColor.Println("    No dependencies")
		}

		if i < len(program.Resources)-1 {
			fmt.Println("    |")
			fmt.Println("    v")
		}
	}

	fmt.Println()
	headerColor.Println("Deployment Flow:")
	headerColor.Println(strings.Repeat("-", 30))

	for i, resource := range program.Resources {
		resourceColor := e.getResourceColor(resource.Type)
		if i > 0 {
			infoColor.Print(" --> ")
		}
		resourceColor.Print(resource.Name)
	}
	fmt.Println("\n")
}

func (e *Engine) findResourceReferences(value interface{}, resources []*ast.Resource) []string {
	var refs []string
	resourceNames := make(map[string]bool)

	for _, res := range resources {
		resourceNames[res.Name] = true
	}

	switch v := value.(type) {
	case string:

		if resourceNames[v] {
			refs = append(refs, v)
		}
	case map[string]interface{}:
		for _, val := range v {
			refs = append(refs, e.findResourceReferences(val, resources)...)
		}
	case []interface{}:
		for _, val := range v {
			refs = append(refs, e.findResourceReferences(val, resources)...)
		}
	}

	return refs
}

func (e *Engine) removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string
	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}
	return result
}

func (e *Engine) getResourceColor(resourceType string) *color.Color {
	switch resourceType {
	case "vpc":
		return color.New(color.FgBlue, color.Bold)
	case "subnet":
		return color.New(color.FgGreen, color.Bold)
	case "security_group":
		return color.New(color.FgYellow, color.Bold)
	case "ec2":
		return color.New(color.FgRed, color.Bold)
	case "internet_gateway":
		return color.New(color.FgCyan, color.Bold)
	case "route_table":
		return color.New(color.FgMagenta, color.Bold)
	case "eip":
		return color.New(color.FgHiGreen, color.Bold)
	case "nat_gateway":
		return color.New(color.FgHiYellow, color.Bold)

	case "data_ami", "data_vpc", "data_subnet", "data_availability_zones", "data_caller_identity":
		return color.New(color.FgHiCyan)
	default:
		return color.New(color.FgWhite, color.Bold)
	}
}
