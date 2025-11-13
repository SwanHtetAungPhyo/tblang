// Code generated from grammar/tblang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // tblang
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by tblangParser.
type tblangVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by tblangParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by tblangParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by tblangParser#blockDeclaration.
	VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{}

	// Visit a parse tree produced by tblangParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by tblangParser#forLoop.
	VisitForLoop(ctx *ForLoopContext) interface{}

	// Visit a parse tree produced by tblangParser#property.
	VisitProperty(ctx *PropertyContext) interface{}

	// Visit a parse tree produced by tblangParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by tblangParser#argumentList.
	VisitArgumentList(ctx *ArgumentListContext) interface{}

	// Visit a parse tree produced by tblangParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by tblangParser#objectLiteral.
	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	// Visit a parse tree produced by tblangParser#objectProperty.
	VisitObjectProperty(ctx *ObjectPropertyContext) interface{}

	// Visit a parse tree produced by tblangParser#arrayLiteral.
	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}
}
