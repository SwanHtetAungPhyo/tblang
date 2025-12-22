
package parser

import "github.com/antlr4-go/antlr/v4"

type BasetblangListener struct{}

var _ tblangListener = &BasetblangListener{}

func (s *BasetblangListener) VisitTerminal(node antlr.TerminalNode) {}

func (s *BasetblangListener) VisitErrorNode(node antlr.ErrorNode) {}

func (s *BasetblangListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

func (s *BasetblangListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

func (s *BasetblangListener) EnterProgram(ctx *ProgramContext) {}

func (s *BasetblangListener) ExitProgram(ctx *ProgramContext) {}

func (s *BasetblangListener) EnterStatement(ctx *StatementContext) {}

func (s *BasetblangListener) ExitStatement(ctx *StatementContext) {}

func (s *BasetblangListener) EnterBlockDeclaration(ctx *BlockDeclarationContext) {}

func (s *BasetblangListener) ExitBlockDeclaration(ctx *BlockDeclarationContext) {}

func (s *BasetblangListener) EnterVariableDeclaration(ctx *VariableDeclarationContext) {}

func (s *BasetblangListener) ExitVariableDeclaration(ctx *VariableDeclarationContext) {}

func (s *BasetblangListener) EnterForLoop(ctx *ForLoopContext) {}

func (s *BasetblangListener) ExitForLoop(ctx *ForLoopContext) {}

func (s *BasetblangListener) EnterProperty(ctx *PropertyContext) {}

func (s *BasetblangListener) ExitProperty(ctx *PropertyContext) {}

func (s *BasetblangListener) EnterFunctionCall(ctx *FunctionCallContext) {}

func (s *BasetblangListener) ExitFunctionCall(ctx *FunctionCallContext) {}

func (s *BasetblangListener) EnterArgumentList(ctx *ArgumentListContext) {}

func (s *BasetblangListener) ExitArgumentList(ctx *ArgumentListContext) {}

func (s *BasetblangListener) EnterExpression(ctx *ExpressionContext) {}

func (s *BasetblangListener) ExitExpression(ctx *ExpressionContext) {}

func (s *BasetblangListener) EnterObjectLiteral(ctx *ObjectLiteralContext) {}

func (s *BasetblangListener) ExitObjectLiteral(ctx *ObjectLiteralContext) {}

func (s *BasetblangListener) EnterObjectProperty(ctx *ObjectPropertyContext) {}

func (s *BasetblangListener) ExitObjectProperty(ctx *ObjectPropertyContext) {}

func (s *BasetblangListener) EnterArrayLiteral(ctx *ArrayLiteralContext) {}

func (s *BasetblangListener) ExitArrayLiteral(ctx *ArrayLiteralContext) {}
