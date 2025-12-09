package compiler

import (
	"fmt"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

// EnterFunctionCall handles function calls (resource declarations)
func (w *ASTWalker) EnterFunctionCall(ctx *parser.FunctionCallContext) {
	// Skip if already processed (but not if we're in manual execution mode)
	if !w.inManualExecution && w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

	funcName := ctx.IDENTIFIER().GetText()

	// Handle print function
	if funcName == "print" {
		args := w.extractArguments(ctx.ArgumentList())
		w.handlePrint(args)
		return
	}

	// Handle output function (alias for print with formatting)
	if funcName == "output" {
		args := w.extractArguments(ctx.ArgumentList())
		w.handleOutput(args)
		return
	}

	// Check if this is a data source function call
	if w.isDataSourceType(funcName) {
		args := w.extractArguments(ctx.ArgumentList())

		if len(args) >= 2 {
			dataSourceName := w.extractStringValue(args[0])
			dataSourceConfig := args[1]

			// Create data source (stored as a special resource type)
			dataSource := &ast.Resource{
				Name:       dataSourceName,
				Type:       funcName,
				Properties: w.convertToMap(dataSourceConfig),
				DependsOn:  []string{},
			}

			w.compiler.resources[dataSourceName] = dataSource
			fmt.Printf("Created data source: %s (%s)\n", dataSourceName, funcName)
		}
		return
	}

	// Check if this is a resource function call
	if w.isResourceType(funcName) {
		args := w.extractArguments(ctx.ArgumentList())

		if len(args) >= 2 {
			resourceName := w.extractStringValue(args[0])
			resourceConfig := args[1]

			// Create resource
			resource := &ast.Resource{
				Name:       resourceName,
				Type:       funcName,
				Properties: w.convertToMap(resourceConfig),
				DependsOn:  []string{},
			}

			w.compiler.resources[resourceName] = resource
			fmt.Printf("Created resource: %s (%s)\n", resourceName, funcName)
		}
	}
}
