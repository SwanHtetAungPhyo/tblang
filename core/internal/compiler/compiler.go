package compiler

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/internal/graph"
	"github.com/tblang/core/parser"
)

type Compiler struct {
	resources        map[string]*ast.Resource
	depGraph         *graph.DependencyGraph
	orderedResources []*ast.Resource
	cloudVendors     map[string]*ast.CloudVendor
	variables        map[string]*ast.Variable
}

type Program struct {
	CloudVendors map[string]*ast.CloudVendor
	Variables    map[string]*ast.Variable
	Resources    []*ast.Resource
}

func New() *Compiler {
	return &Compiler{
		resources:    make(map[string]*ast.Resource),
		depGraph:     graph.NewDependencyGraph(),
		cloudVendors: make(map[string]*ast.CloudVendor),
		variables:    make(map[string]*ast.Variable),
	}
}

func (c *Compiler) CompileFile(filename string) (*Program, error) {

	input, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	inputStream := antlr.NewInputStream(string(input))
	lexer := parser.NewtblangLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewtblangParser(stream)

	tree := p.Program()

	walker := &ASTWalker{compiler: c}
	antlr.ParseTreeWalkerDefault.Walk(walker, tree)

	if err := c.buildDependencyGraph(); err != nil {
		return nil, fmt.Errorf("failed to build dependency graph: %w", err)
	}

	program := &Program{
		CloudVendors: c.cloudVendors,
		Variables:    c.variables,
		Resources:    c.orderedResources,
	}

	return program, nil
}

func (c *Compiler) buildDependencyGraph() error {
	fmt.Println("Building dependency graph...")

	for _, resource := range c.resources {
		c.depGraph.AddResource(resource)
	}

	if err := c.depGraph.AnalyzeDependencies(); err != nil {
		return err
	}

	c.depGraph.PrintGraph()

	orderedResources, err := c.depGraph.TopologicalSort()
	if err != nil {
		return err
	}

	c.orderedResources = orderedResources

	fmt.Printf("Dependency graph built successfully. Resource order: ")
	for i, resource := range orderedResources {
		if i > 0 {
			fmt.Print(" -> ")
		}
		fmt.Print(resource.Name)
	}
	fmt.Println()

	return nil
}
