
package parser
import "github.com/antlr4-go/antlr/v4"

type tblangVisitor interface {
	antlr.ParseTreeVisitor

	VisitProgram(ctx *ProgramContext) interface{}

	VisitStatement(ctx *StatementContext) interface{}

	VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{}

	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	VisitForLoop(ctx *ForLoopContext) interface{}

	VisitProperty(ctx *PropertyContext) interface{}

	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	VisitArgumentList(ctx *ArgumentListContext) interface{}

	VisitExpression(ctx *ExpressionContext) interface{}

	VisitObjectLiteral(ctx *ObjectLiteralContext) interface{}

	VisitObjectProperty(ctx *ObjectPropertyContext) interface{}

	VisitArrayLiteral(ctx *ArrayLiteralContext) interface{}
}
