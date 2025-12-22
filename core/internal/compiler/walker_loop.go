package compiler

import (
	"fmt"

	"github.com/tblang/core/parser"
)

func (w *ASTWalker) EnterForLoop(ctx *parser.ForLoopContext) {
	if w.processedContexts != nil && w.processedContexts[ctx] {
		return
	}

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

	savedVars := w.variables

	statements := ctx.AllStatement()

	for _, item := range items {

		if w.variables == nil {
			w.variables = make(map[string]interface{})
		} else {

			newVars := make(map[string]interface{})
			for k, v := range savedVars {
				newVars[k] = v
			}
			w.variables = newVars
		}

		w.variables[iterator] = item

		w.inManualExecution = true
		for _, stmt := range statements {
			w.executeStatement(stmt)
		}
		w.inManualExecution = false
	}

	w.variables = savedVars
}

func (w *ASTWalker) markStatementAsProcessed(stmt parser.IStatementContext) {
	stmtCtx := stmt.(*parser.StatementContext)

	if stmtCtx.VariableDeclaration() != nil {
		varDecl := stmtCtx.VariableDeclaration().(*parser.VariableDeclarationContext)
		w.processedContexts[varDecl] = true

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

func (w *ASTWalker) executeStatement(stmt parser.IStatementContext) {
	stmtCtx := stmt.(*parser.StatementContext)

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
