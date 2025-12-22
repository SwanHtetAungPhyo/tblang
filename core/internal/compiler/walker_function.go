package compiler

import (
	"fmt"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

func (w *ASTWalker) EnterFunctionCall(ctx *parser.FunctionCallContext) {

	if !w.inManualExecution && w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

	funcName := ctx.IDENTIFIER().GetText()

	if funcName == "print" {
		args := w.extractArguments(ctx.ArgumentList())
		w.handlePrint(args)
		return
	}

	if funcName == "output" {
		args := w.extractArguments(ctx.ArgumentList())
		w.handleOutput(args)
		return
	}

	if w.isDataSourceType(funcName) {
		args := w.extractArguments(ctx.ArgumentList())

		if len(args) >= 2 {
			dataSourceName := w.extractStringValue(args[0])
			dataSourceConfig := args[1]

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

	if w.isResourceType(funcName) {
		args := w.extractArguments(ctx.ArgumentList())

		if len(args) >= 2 {
			resourceName := w.extractStringValue(args[0])
			resourceConfig := args[1]

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
