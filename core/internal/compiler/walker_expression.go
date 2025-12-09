package compiler

import (
	"strconv"
	"strings"

	"github.com/tblang/core/parser"
)

// evaluateExpression evaluates an expression and returns its value
func (w *ASTWalker) evaluateExpression(expr parser.IExpressionContext) interface{} {
	if expr == nil {
		return nil
	}

	switch e := expr.(type) {
	case *parser.ExpressionContext:
		if e.DOT() != nil && e.Expression() != nil && e.IDENTIFIER() != nil {
			obj := w.evaluateExpression(e.Expression())
			propName := e.IDENTIFIER().GetText()

			if objMap, ok := obj.(map[string]interface{}); ok {
				if val, exists := objMap[propName]; exists {
					return val
				}
			}
			return nil
		}

		if e.STRING_LITERAL() != nil {
			return strings.Trim(e.STRING_LITERAL().GetText(), `"'`)
		}
		if e.NUMBER() != nil {
			if val, err := strconv.ParseFloat(e.NUMBER().GetText(), 64); err == nil {
				return val
			}
		}
		if e.BOOLEAN() != nil {
			return e.BOOLEAN().GetText() == "true"
		}
		if e.IDENTIFIER() != nil {
			varName := e.IDENTIFIER().GetText()
			if w.variables != nil {
				if val, exists := w.variables[varName]; exists {
					return val
				}
			}
			return varName // Return as string reference
		}
		if e.ObjectLiteral() != nil {
			return w.evaluateObjectLiteral(e.ObjectLiteral())
		}
		if e.ArrayLiteral() != nil {
			return w.evaluateArrayLiteral(e.ArrayLiteral())
		}
		if e.FunctionCall() != nil {
			return w.evaluateFunctionCall(e.FunctionCall())
		}
		// Handle parenthesized expressions
		if e.LPAREN() != nil && e.Expression() != nil {
			return w.evaluateExpression(e.Expression())
		}
	}

	return nil
}

func (w *ASTWalker) evaluateObjectLiteral(obj parser.IObjectLiteralContext) map[string]interface{} {
	result := make(map[string]interface{})

	if objCtx, ok := obj.(*parser.ObjectLiteralContext); ok {
		for _, prop := range objCtx.AllObjectProperty() {
			key := prop.IDENTIFIER().GetText()
			value := w.evaluateExpression(prop.Expression())
			result[key] = value
		}
	}

	return result
}

func (w *ASTWalker) evaluateArrayLiteral(arr parser.IArrayLiteralContext) []interface{} {
	var result []interface{}

	if arrCtx, ok := arr.(*parser.ArrayLiteralContext); ok {
		for _, expr := range arrCtx.AllExpression() {
			result = append(result, w.evaluateExpression(expr))
		}
	}

	return result
}

func (w *ASTWalker) evaluateFunctionCall(funcCall parser.IFunctionCallContext) interface{} {
	if funcCtx, ok := funcCall.(*parser.FunctionCallContext); ok {
		funcName := funcCtx.IDENTIFIER().GetText()
		args := w.extractArguments(funcCtx.ArgumentList())

		// For resource references, return the resource name
		if w.isResourceType(funcName) && len(args) > 0 {
			return w.extractStringValue(args[0])
		}
	}

	return nil
}
