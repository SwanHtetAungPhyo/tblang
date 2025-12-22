
package parser
import "github.com/antlr4-go/antlr/v4"

type BasetblangVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasetblangVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitBlockDeclaration(ctx *BlockDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitForLoop(ctx *ForLoopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitProperty(ctx *PropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitArgumentList(ctx *ArgumentListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitObjectLiteral(ctx *ObjectLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitObjectProperty(ctx *ObjectPropertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasetblangVisitor) VisitArrayLiteral(ctx *ArrayLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}
