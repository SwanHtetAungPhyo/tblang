package compiler

import (
	"fmt"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

// EnterVariableDeclaration handles declare statements
func (w *ASTWalker) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {
	// Skip if already processed (but not if we're in manual execution mode)
	if !w.inManualExecution && w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

	varName := ctx.IDENTIFIER().GetText()

	// Check if the expression is a function call (resource declaration)
	expr := ctx.Expression()
	if exprCtx, ok := expr.(*parser.ExpressionContext); ok {
		if exprCtx.FunctionCall() != nil {
			funcCallCtx := exprCtx.FunctionCall().(*parser.FunctionCallContext)
			funcName := funcCallCtx.IDENTIFIER().GetText()

			// If it's a resource type, process it as a resource
			if w.isResourceType(funcName) {
				args := w.extractArguments(funcCallCtx.ArgumentList())

				if len(args) >= 2 {
					resourceName := w.extractStringValue(args[0])
					resourceConfig := args[1]

					// Create resource
					props := w.convertToMap(resourceConfig)
					resource := &ast.Resource{
						Name:       resourceName,
						Type:       funcName,
						Properties: props,
						DependsOn:  []string{},
					}

					w.compiler.resources[resourceName] = resource
					fmt.Printf("Created resource: %s (%s)\n", resourceName, funcName)

					// Store variable reference to the resource
					if w.variables == nil {
						w.variables = make(map[string]interface{})
					}
					w.variables[varName] = resourceName

					variable := &ast.Variable{
						Name:  varName,
						Value: resourceName,
					}
					w.compiler.variables[varName] = variable

					fmt.Printf("Declared variable: %s\n", varName)
					return
				}
			}
		}
	}

	// Regular variable declaration
	value := w.evaluateExpression(ctx.Expression())

	if w.variables == nil {
		w.variables = make(map[string]interface{})
	}
	w.variables[varName] = value

	// Store in compiler
	variable := &ast.Variable{
		Name:  varName,
		Value: value,
	}
	w.compiler.variables[varName] = variable

	fmt.Printf("Declared variable: %s\n", varName)
}
