package compiler

import (
	"fmt"

	"github.com/tblang/core/parser"
)

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
