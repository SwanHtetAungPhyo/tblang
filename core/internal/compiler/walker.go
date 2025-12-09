package compiler

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tblang/core/internal/ast"
	"github.com/tblang/core/parser"
)

// ASTWalker walks the parse tree and builds the internal AST
type ASTWalker struct {
	*parser.BasetblangListener
	compiler          *Compiler
	variables         map[string]interface{}
	processedContexts map[interface{}]bool // Track contexts we've manually processed
	inManualExecution bool                 // Flag to indicate we're manually executing (e.g., in a loop)
}

// EnterBlockDeclaration handles cloud_vendor blocks
func (w *ASTWalker) EnterBlockDeclaration(ctx *parser.BlockDeclarationContext) {
	blockType := ctx.IDENTIFIER().GetText()
	blockName := strings.Trim(ctx.STRING_LITERAL().GetText(), `"'`)

	if blockType == "cloud_vendor" {
		// Extract properties from the block
		properties := make(map[string]interface{})

		for _, prop := range ctx.AllProperty() {
			propCtx := prop.(*parser.PropertyContext)
			key := propCtx.IDENTIFIER().GetText()
			value := w.evaluateExpression(propCtx.Expression())
			properties[key] = value
		}

		// Create cloud vendor
		cloudVendor := &ast.CloudVendor{
			Name:       blockName,
			Properties: properties,
		}

		w.compiler.cloudVendors[blockName] = cloudVendor
		fmt.Printf("Registered cloud vendor: %s\n", blockName)
	}
}

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

// EnterForLoop handles for loops
func (w *ASTWalker) EnterForLoop(ctx *parser.ForLoopContext) {
	if w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

	// Mark this loop and all its children as processed
	if w.processedContexts == nil {
		w.processedContexts = make(map[interface{}]bool)
	}
	w.processedContexts[ctx] = true

	for _, stmt := range ctx.AllStatement() {
		w.markStatementAsProcessed(stmt)
	}

	iterator := ctx.IDENTIFIER().GetText()
	collectionExpr := ctx.Expression()

	fmt.Printf("Processing for loop: %s in collection\n", iterator)

	var collectionName string
	if collectionExpr.IDENTIFIER() != nil {
		collectionName = collectionExpr.IDENTIFIER().GetText()
	}

	var items []interface{}
	if w.compiler.variables != nil {
		if val, exists := w.compiler.variables[collectionName]; exists {
			if arr, ok := val.Value.([]interface{}); ok {
				items = arr
				fmt.Printf("  Found collection '%s' with %d items\n", collectionName, len(items))
			} else {
				fmt.Printf("  Collection '%s' is not an array: %T\n", collectionName, val.Value)
			}
		} else {
			fmt.Printf("  Collection '%s' not found in variables\n", collectionName)
		}
	} else {
		fmt.Printf("  No variables available\n")
	}

	// Save current variables state
	savedVars := w.variables

	// Execute loop body for each item
	statements := ctx.AllStatement()

	for _, item := range items {
		// Create new scope with iterator variable
		if w.variables == nil {
			w.variables = make(map[string]interface{})
		} else {
			// Copy parent scope
			newVars := make(map[string]interface{})
			for k, v := range savedVars {
				newVars[k] = v
			}
			w.variables = newVars
		}

		// Set iterator variable in current scope
		w.variables[iterator] = item

		// Execute each statement in the loop body
		w.inManualExecution = true
		for _, stmt := range statements {
			w.executeStatement(stmt)
		}
		w.inManualExecution = false
	}

	// Restore variables
	w.variables = savedVars
}

// markStatementAsProcessed marks a statement and all its children as processed
func (w *ASTWalker) markStatementAsProcessed(stmt parser.IStatementContext) {
	stmtCtx := stmt.(*parser.StatementContext)

	if stmtCtx.VariableDeclaration() != nil {
		varDecl := stmtCtx.VariableDeclaration().(*parser.VariableDeclarationContext)
		w.processedContexts[varDecl] = true

		// Also mark any function calls within the variable declaration
		expr := varDecl.Expression()
		if exprCtx, ok := expr.(*parser.ExpressionContext); ok {
			if exprCtx.FunctionCall() != nil {
				w.processedContexts[exprCtx.FunctionCall()] = true
			}
		}
	} else if stmtCtx.FunctionCall() != nil {
		w.processedContexts[stmtCtx.FunctionCall()] = true
	} else if stmtCtx.BlockDeclaration() != nil {
		w.processedContexts[stmtCtx.BlockDeclaration()] = true
	} else if stmtCtx.ForLoop() != nil {
		w.processedContexts[stmtCtx.ForLoop()] = true
	}
}

// executeStatement processes a single statement
func (w *ASTWalker) executeStatement(stmt parser.IStatementContext) {
	stmtCtx := stmt.(*parser.StatementContext)

	// DON'T mark as processed here - we want to allow multiple executions in loops
	// Instead, we'll mark the loop itself as processed

	if stmtCtx.VariableDeclaration() != nil {
		ctx := stmtCtx.VariableDeclaration().(*parser.VariableDeclarationContext)
		w.EnterVariableDeclaration(ctx)
	} else if stmtCtx.FunctionCall() != nil {
		ctx := stmtCtx.FunctionCall().(*parser.FunctionCallContext)
		w.EnterFunctionCall(ctx)
	} else if stmtCtx.BlockDeclaration() != nil {
		ctx := stmtCtx.BlockDeclaration().(*parser.BlockDeclarationContext)
		w.EnterBlockDeclaration(ctx)
	} else if stmtCtx.ForLoop() != nil {
		ctx := stmtCtx.ForLoop().(*parser.ForLoopContext)
		w.EnterForLoop(ctx)
	}
}

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

// Helper methods

func (w *ASTWalker) isResourceType(funcName string) bool {
	resourceTypes := []string{"vpc", "subnet", "security_group", "ec2", "internet_gateway", "route_table", "eip", "nat_gateway"}
	for _, rt := range resourceTypes {
		if rt == funcName {
			return true
		}
	}
	return false
}

func (w *ASTWalker) isDataSourceType(funcName string) bool {
	dataSourceTypes := []string{"data_ami", "data_vpc", "data_subnet", "data_availability_zones", "data_caller_identity"}
	for _, dt := range dataSourceTypes {
		if dt == funcName {
			return true
		}
	}
	return false
}

func (w *ASTWalker) extractArguments(argList parser.IArgumentListContext) []interface{} {
	if argList == nil {
		return []interface{}{}
	}

	var args []interface{}
	for _, expr := range argList.(*parser.ArgumentListContext).AllExpression() {
		val := w.evaluateExpression(expr)
		args = append(args, val)
	}
	return args
}

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

func (w *ASTWalker) extractStringValue(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}
	return fmt.Sprintf("%v", value)
}

func (w *ASTWalker) convertToMap(value interface{}) map[string]interface{} {
	if m, ok := value.(map[string]interface{}); ok {
		return m
	}
	return make(map[string]interface{})
}
