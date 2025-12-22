package compiler

import (
	"fmt"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

func (w *ASTWalker) EnterVariableDeclaration(ctx *parser.VariableDeclarationContext) {

	if !w.inManualExecution && w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

	varName := ctx.IDENTIFIER().GetText()

	expr := ctx.Expression()
	if exprCtx, ok := expr.(*parser.ExpressionContext); ok {
		if exprCtx.FunctionCall() != nil {
			funcCallCtx := exprCtx.FunctionCall().(*parser.FunctionCallContext)
			funcName := funcCallCtx.IDENTIFIER().GetText()

			if w.isResourceType(funcName) {
				args := w.extractArguments(funcCallCtx.ArgumentList())

				if len(args) >= 2 {
					resourceName := w.extractStringValue(args[0])
					resourceConfig := args[1]

					props := w.convertToMap(resourceConfig)
					resource := &ast.Resource{
						Name:       resourceName,
						Type:       funcName,
						Properties: props,
						DependsOn:  []string{},
					}

					w.compiler.resources[resourceName] = resource
					fmt.Printf("Created resource: %s (%s)\n", resourceName, funcName)

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

	value := w.evaluateExpression(ctx.Expression())

	if w.variables == nil {
		w.variables = make(map[string]interface{})
	}
	w.variables[varName] = value

	variable := &ast.Variable{
		Name:  varName,
		Value: value,
	}
	w.compiler.variables[varName] = variable

	fmt.Printf("Declared variable: %s\n", varName)
}
