
package parser

import "github.com/antlr4-go/antlr/v4"

type tblangListener interface {
	antlr.ParseTreeListener

	EnterProgram(c *ProgramContext)

	EnterStatement(c *StatementContext)

	EnterBlockDeclaration(c *BlockDeclarationContext)

	EnterVariableDeclaration(c *VariableDeclarationContext)

	EnterForLoop(c *ForLoopContext)

	EnterProperty(c *PropertyContext)

	EnterFunctionCall(c *FunctionCallContext)

	EnterArgumentList(c *ArgumentListContext)

	EnterExpression(c *ExpressionContext)

	EnterObjectLiteral(c *ObjectLiteralContext)

	EnterObjectProperty(c *ObjectPropertyContext)

	EnterArrayLiteral(c *ArrayLiteralContext)

	ExitProgram(c *ProgramContext)

	ExitStatement(c *StatementContext)

	ExitBlockDeclaration(c *BlockDeclarationContext)

	ExitVariableDeclaration(c *VariableDeclarationContext)

	ExitForLoop(c *ForLoopContext)

	ExitProperty(c *PropertyContext)

	ExitFunctionCall(c *FunctionCallContext)

	ExitArgumentList(c *ArgumentListContext)

	ExitExpression(c *ExpressionContext)

	ExitObjectLiteral(c *ObjectLiteralContext)

	ExitObjectProperty(c *ObjectPropertyContext)

	ExitArrayLiteral(c *ArrayLiteralContext)
}
