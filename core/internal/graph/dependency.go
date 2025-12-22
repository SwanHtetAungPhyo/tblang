package graph

import (
	"fmt"
	"sort"

	"github.com/tblang/core/internal/ast"
)

type DependencyGraph struct {
	nodes map[string]*Node
	edges map[string][]string
}

type Node struct {
	Resource     *ast.Resource
	Dependencies []string
	Dependents   []string
}

func NewDependencyGraph() *DependencyGraph {
	return &DependencyGraph{
		nodes: make(map[string]*Node),
		edges: make(map[string][]string),
	}
}

func (dg *DependencyGraph) AddResource(resource *ast.Resource) {
	node := &Node{
		Resource:     resource,
		Dependencies: make([]string, 0),
		Dependents:   make([]string, 0),
	}

	dg.nodes[resource.Name] = node
	dg.edges[resource.Name] = make([]string, 0)
}

func (dg *DependencyGraph) AnalyzeDependencies() error {
	for _, node := range dg.nodes {
		deps := dg.extractDependencies(node.Resource)
		for _, dep := range deps {
			if err := dg.AddDependency(node.Resource.Name, dep); err != nil {
				return err
			}
		}
	}
	return nil
}

func (dg *DependencyGraph) extractDependencies(resource *ast.Resource) []string {
	var dependencies []string

	for _, value := range resource.Properties {
		deps := dg.findResourceReferences(value)
		dependencies = append(dependencies, deps...)
	}

	return dependencies
}

func (dg *DependencyGraph) findResourceReferences(value interface{}) []string {
	var refs []string

	switch v := value.(type) {
	case string:

		if _, exists := dg.nodes[v]; exists {
			refs = append(refs, v)
		}

	case map[string]interface{}:
		for _, val := range v {
			refs = append(refs, dg.findResourceReferences(val)...)
		}

	case []interface{}:
		for _, val := range v {
			refs = append(refs, dg.findResourceReferences(val)...)
		}
	}

	return refs
}

func (dg *DependencyGraph) AddDependency(resource, dependency string) error {

	if resource == dependency {
		return nil
	}

	if _, exists := dg.nodes[resource]; !exists {
		return fmt.Errorf("resource %s not found", resource)
	}
	if _, exists := dg.nodes[dependency]; !exists {

		return nil
	}

	for _, existingDep := range dg.edges[resource] {
		if existingDep == dependency {
			return nil
		}
	}

	dg.edges[resource] = append(dg.edges[resource], dependency)

	dg.nodes[resource].Dependencies = append(dg.nodes[resource].Dependencies, dependency)
	dg.nodes[dependency].Dependents = append(dg.nodes[dependency].Dependents, resource)

	return nil
}

func (dg *DependencyGraph) TopologicalSort() ([]*ast.Resource, error) {

	if dg.hasCycle() {
		return nil, fmt.Errorf("circular dependency detected")
	}

	var result []*ast.Resource
	visited := make(map[string]bool)
	temp := make(map[string]bool)

	var visit func(string) error
	visit = func(name string) error {
		if temp[name] {
			return fmt.Errorf("circular dependency detected at %s", name)
		}
		if visited[name] {
			return nil
		}

		temp[name] = true

		for _, dep := range dg.edges[name] {
			if err := visit(dep); err != nil {
				return err
			}
		}

		temp[name] = false
		visited[name] = true

		if node, exists := dg.nodes[name]; exists {
			result = append(result, node.Resource)
		}

		return nil
	}

	for name := range dg.nodes {
		if !visited[name] {
			if err := visit(name); err != nil {
				return nil, err
			}
		}
	}

	return result, nil
}

func (dg *DependencyGraph) hasCycle() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var hasCycleUtil func(string) bool
	hasCycleUtil = func(name string) bool {
		visited[name] = true
		recStack[name] = true

		for _, dep := range dg.edges[name] {
			if !visited[dep] && hasCycleUtil(dep) {
				return true
			} else if recStack[dep] {
				return true
			}
		}

		recStack[name] = false
		return false
	}

	for name := range dg.nodes {
		if !visited[name] && hasCycleUtil(name) {
			return true
		}
	}

	return false
}

func (dg *DependencyGraph) GetDependencies(resourceName string) []string {
	if node, exists := dg.nodes[resourceName]; exists {
		return node.Dependencies
	}
	return nil
}

func (dg *DependencyGraph) GetDependents(resourceName string) []string {
	if node, exists := dg.nodes[resourceName]; exists {
		return node.Dependents
	}
	return nil
}

func (dg *DependencyGraph) PrintGraph() {
	fmt.Println("=== Dependency Graph ===")

	var names []string
	for name := range dg.nodes {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		node := dg.nodes[name]
		fmt.Printf("Resource: %s (%s)\n", name, node.Resource.Type)

		if len(node.Dependencies) > 0 {
			fmt.Printf("  Dependencies: %v\n", node.Dependencies)
		}

		if len(node.Dependents) > 0 {
			fmt.Printf("  Dependents: %v\n", node.Dependents)
		}

		fmt.Println()
	}
}
