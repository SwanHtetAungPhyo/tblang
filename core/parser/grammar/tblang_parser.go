
package parser
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type tblangParser struct {
	*antlr.BaseParser
}

var TblangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func tblangParserInit() {
	staticData := &TblangParserStaticData
	staticData.LiteralNames = []string{
		"", "'declare'", "'for'", "'in'", "", "", "", "", "'='", "':'", "';'",
		"','", "'.'", "'('", "')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "DECLARE", "FOR", "IN", "STRING_LITERAL", "NUMBER", "BOOLEAN", "IDENTIFIER",
		"ASSIGN", "COLON", "SEMICOLON", "COMMA", "DOT", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "LBRACKET", "RBRACKET", "LINE_COMMENT", "BLOCK_COMMENT",
		"WS",
	}
	staticData.RuleNames = []string{
		"program", "statement", "blockDeclaration", "variableDeclaration", "forLoop",
		"property", "functionCall", "argumentList", "expression", "objectLiteral",
		"objectProperty", "arrayLiteral",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 21, 144, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 1, 0, 5, 0, 26, 8, 0, 10, 0, 12, 0, 29, 9, 0, 1, 0, 1,
		0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 38, 8, 1, 1, 2, 1, 2, 1, 2, 1, 2,
		5, 2, 44, 8, 2, 10, 2, 12, 2, 47, 9, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1,
		3, 1, 3, 3, 3, 56, 8, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 5, 4, 64,
		8, 4, 10, 4, 12, 4, 67, 9, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		75, 8, 5, 1, 6, 1, 6, 1, 6, 3, 6, 80, 8, 6, 1, 6, 1, 6, 3, 6, 84, 8, 6,
		1, 7, 1, 7, 1, 7, 5, 7, 89, 8, 7, 10, 7, 12, 7, 92, 9, 7, 1, 8, 1, 8, 1,
		8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 1, 8, 3, 8, 106, 8,
		8, 1, 8, 1, 8, 1, 8, 5, 8, 111, 8, 8, 10, 8, 12, 8, 114, 9, 8, 1, 9, 1,
		9, 5, 9, 118, 8, 9, 10, 9, 12, 9, 121, 9, 9, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 129, 8, 10, 1, 11, 1, 11, 1, 11, 1, 11, 5, 11, 135,
		8, 11, 10, 11, 12, 11, 138, 9, 11, 3, 11, 140, 8, 11, 1, 11, 1, 11, 1,
		11, 0, 1, 16, 12, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 0, 1, 1, 0,
		8, 9, 155, 0, 27, 1, 0, 0, 0, 2, 37, 1, 0, 0, 0, 4, 39, 1, 0, 0, 0, 6,
		50, 1, 0, 0, 0, 8, 57, 1, 0, 0, 0, 10, 70, 1, 0, 0, 0, 12, 76, 1, 0, 0,
		0, 14, 85, 1, 0, 0, 0, 16, 105, 1, 0, 0, 0, 18, 115, 1, 0, 0, 0, 20, 124,
		1, 0, 0, 0, 22, 130, 1, 0, 0, 0, 24, 26, 3, 2, 1, 0, 25, 24, 1, 0, 0, 0,
		26, 29, 1, 0, 0, 0, 27, 25, 1, 0, 0, 0, 27, 28, 1, 0, 0, 0, 28, 30, 1,
		0, 0, 0, 29, 27, 1, 0, 0, 0, 30, 31, 5, 0, 0, 1, 31, 1, 1, 0, 0, 0, 32,
		38, 3, 4, 2, 0, 33, 38, 3, 6, 3, 0, 34, 38, 3, 8, 4, 0, 35, 38, 3, 12,
		6, 0, 36, 38, 5, 10, 0, 0, 37, 32, 1, 0, 0, 0, 37, 33, 1, 0, 0, 0, 37,
		34, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 36, 1, 0, 0, 0, 38, 3, 1, 0, 0,
		0, 39, 40, 5, 7, 0, 0, 40, 41, 5, 4, 0, 0, 41, 45, 5, 15, 0, 0, 42, 44,
		3, 10, 5, 0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0,
		45, 46, 1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5,
		16, 0, 0, 49, 5, 1, 0, 0, 0, 50, 51, 5, 1, 0, 0, 51, 52, 5, 7, 0, 0, 52,
		53, 5, 8, 0, 0, 53, 55, 3, 16, 8, 0, 54, 56, 5, 10, 0, 0, 55, 54, 1, 0,
		0, 0, 55, 56, 1, 0, 0, 0, 56, 7, 1, 0, 0, 0, 57, 58, 5, 2, 0, 0, 58, 59,
		5, 7, 0, 0, 59, 60, 5, 3, 0, 0, 60, 61, 3, 16, 8, 0, 61, 65, 5, 15, 0,
		0, 62, 64, 3, 2, 1, 0, 63, 62, 1, 0, 0, 0, 64, 67, 1, 0, 0, 0, 65, 63,
		1, 0, 0, 0, 65, 66, 1, 0, 0, 0, 66, 68, 1, 0, 0, 0, 67, 65, 1, 0, 0, 0,
		68, 69, 5, 16, 0, 0, 69, 9, 1, 0, 0, 0, 70, 71, 5, 7, 0, 0, 71, 72, 5,
		8, 0, 0, 72, 74, 3, 16, 8, 0, 73, 75, 5, 10, 0, 0, 74, 73, 1, 0, 0, 0,
		74, 75, 1, 0, 0, 0, 75, 11, 1, 0, 0, 0, 76, 77, 5, 7, 0, 0, 77, 79, 5,
		13, 0, 0, 78, 80, 3, 14, 7, 0, 79, 78, 1, 0, 0, 0, 79, 80, 1, 0, 0, 0,
		80, 81, 1, 0, 0, 0, 81, 83, 5, 14, 0, 0, 82, 84, 5, 10, 0, 0, 83, 82, 1,
		0, 0, 0, 83, 84, 1, 0, 0, 0, 84, 13, 1, 0, 0, 0, 85, 90, 3, 16, 8, 0, 86,
		87, 5, 11, 0, 0, 87, 89, 3, 16, 8, 0, 88, 86, 1, 0, 0, 0, 89, 92, 1, 0,
		0, 0, 90, 88, 1, 0, 0, 0, 90, 91, 1, 0, 0, 0, 91, 15, 1, 0, 0, 0, 92, 90,
		1, 0, 0, 0, 93, 94, 6, 8, -1, 0, 94, 106, 5, 4, 0, 0, 95, 106, 5, 5, 0,
		0, 96, 106, 5, 6, 0, 0, 97, 106, 5, 7, 0, 0, 98, 106, 3, 18, 9, 0, 99,
		106, 3, 22, 11, 0, 100, 106, 3, 12, 6, 0, 101, 102, 5, 13, 0, 0, 102, 103,
		3, 16, 8, 0, 103, 104, 5, 14, 0, 0, 104, 106, 1, 0, 0, 0, 105, 93, 1, 0,
		0, 0, 105, 95, 1, 0, 0, 0, 105, 96, 1, 0, 0, 0, 105, 97, 1, 0, 0, 0, 105,
		98, 1, 0, 0, 0, 105, 99, 1, 0, 0, 0, 105, 100, 1, 0, 0, 0, 105, 101, 1,
		0, 0, 0, 106, 112, 1, 0, 0, 0, 107, 108, 10, 2, 0, 0, 108, 109, 5, 12,
		0, 0, 109, 111, 5, 7, 0, 0, 110, 107, 1, 0, 0, 0, 111, 114, 1, 0, 0, 0,
		112, 110, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 17, 1, 0, 0, 0, 114, 112,
		1, 0, 0, 0, 115, 119, 5, 15, 0, 0, 116, 118, 3, 20, 10, 0, 117, 116, 1,
		0, 0, 0, 118, 121, 1, 0, 0, 0, 119, 117, 1, 0, 0, 0, 119, 120, 1, 0, 0,
		0, 120, 122, 1, 0, 0, 0, 121, 119, 1, 0, 0, 0, 122, 123, 5, 16, 0, 0, 123,
		19, 1, 0, 0, 0, 124, 125, 5, 7, 0, 0, 125, 126, 7, 0, 0, 0, 126, 128, 3,
		16, 8, 0, 127, 129, 5, 11, 0, 0, 128, 127, 1, 0, 0, 0, 128, 129, 1, 0,
		0, 0, 129, 21, 1, 0, 0, 0, 130, 139, 5, 17, 0, 0, 131, 136, 3, 16, 8, 0,
		132, 133, 5, 11, 0, 0, 133, 135, 3, 16, 8, 0, 134, 132, 1, 0, 0, 0, 135,
		138, 1, 0, 0, 0, 136, 134, 1, 0, 0, 0, 136, 137, 1, 0, 0, 0, 137, 140,
		1, 0, 0, 0, 138, 136, 1, 0, 0, 0, 139, 131, 1, 0, 0, 0, 139, 140, 1, 0,
		0, 0, 140, 141, 1, 0, 0, 0, 141, 142, 5, 18, 0, 0, 142, 23, 1, 0, 0, 0,
		15, 27, 37, 45, 55, 65, 74, 79, 83, 90, 105, 112, 119, 128, 136, 139,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

func TblangParserInit() {
	staticData := &TblangParserStaticData
	staticData.once.Do(tblangParserInit)
}

func NewtblangParser(input antlr.TokenStream) *tblangParser {
	TblangParserInit()
	this := new(tblangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &TblangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "tblang.g4"

	return this
}

const (
	tblangParserEOF            = antlr.TokenEOF
	tblangParserDECLARE        = 1
	tblangParserFOR            = 2
	tblangParserIN             = 3
	tblangParserSTRING_LITERAL = 4
	tblangParserNUMBER         = 5
	tblangParserBOOLEAN        = 6
	tblangParserIDENTIFIER     = 7
	tblangParserASSIGN         = 8
	tblangParserCOLON          = 9
	tblangParserSEMICOLON      = 10
	tblangParserCOMMA          = 11
	tblangParserDOT            = 12
	tblangParserLPAREN         = 13
	tblangParserRPAREN         = 14
	tblangParserLBRACE         = 15
	tblangParserRBRACE         = 16
	tblangParserLBRACKET       = 17
	tblangParserRBRACKET       = 18
	tblangParserLINE_COMMENT   = 19
	tblangParserBLOCK_COMMENT  = 20
	tblangParserWS             = 21
)

const (
	tblangParserRULE_program             = 0
	tblangParserRULE_statement           = 1
	tblangParserRULE_blockDeclaration    = 2
	tblangParserRULE_variableDeclaration = 3
	tblangParserRULE_forLoop             = 4
	tblangParserRULE_property            = 5
	tblangParserRULE_functionCall        = 6
	tblangParserRULE_argumentList        = 7
	tblangParserRULE_expression          = 8
	tblangParserRULE_objectLiteral       = 9
	tblangParserRULE_objectProperty      = 10
	tblangParserRULE_arrayLiteral        = 11
)

type IProgramContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(tblangParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (s *ProgramContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitProgram(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, tblangParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(27)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1158) != 0 {
		{
			p.SetState(24)
			p.Statement()
		}

		p.SetState(29)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(30)
		p.Match(tblangParserEOF)
		if p.HasError() {

			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IStatementContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	BlockDeclaration() IBlockDeclarationContext
	VariableDeclaration() IVariableDeclarationContext
	ForLoop() IForLoopContext
	FunctionCall() IFunctionCallContext
	SEMICOLON() antlr.TerminalNode

	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) BlockDeclaration() IBlockDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockDeclarationContext)
}

func (s *StatementContext) VariableDeclaration() IVariableDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableDeclarationContext)
}

func (s *StatementContext) ForLoop() IForLoopContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForLoopContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForLoopContext)
}

func (s *StatementContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *StatementContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(tblangParserSEMICOLON, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, tblangParserRULE_statement)
	p.SetState(37)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(32)
			p.BlockDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.VariableDeclaration()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(34)
			p.ForLoop()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(35)
			p.FunctionCall()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(36)
			p.Match(tblangParserSEMICOLON)
			if p.HasError() {

				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IBlockDeclarationContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IDENTIFIER() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllProperty() []IPropertyContext
	Property(i int) IPropertyContext

	IsBlockDeclarationContext()
}

type BlockDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockDeclarationContext() *BlockDeclarationContext {
	var p = new(BlockDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_blockDeclaration
	return p
}

func InitEmptyBlockDeclarationContext(p *BlockDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_blockDeclaration
}

func (*BlockDeclarationContext) IsBlockDeclarationContext() {}

func NewBlockDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockDeclarationContext {
	var p = new(BlockDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_blockDeclaration

	return p
}

func (s *BlockDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *BlockDeclarationContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(tblangParserSTRING_LITERAL, 0)
}

func (s *BlockDeclarationContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserLBRACE, 0)
}

func (s *BlockDeclarationContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserRBRACE, 0)
}

func (s *BlockDeclarationContext) AllProperty() []IPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPropertyContext); ok {
			len++
		}
	}

	tst := make([]IPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPropertyContext); ok {
			tst[i] = t.(IPropertyContext)
			i++
		}
	}

	return tst
}

func (s *BlockDeclarationContext) Property(i int) IPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *BlockDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterBlockDeclaration(s)
	}
}

func (s *BlockDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitBlockDeclaration(s)
	}
}

func (s *BlockDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitBlockDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) BlockDeclaration() (localctx IBlockDeclarationContext) {
	localctx = NewBlockDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, tblangParserRULE_blockDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(39)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(40)
		p.Match(tblangParserSTRING_LITERAL)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(41)
		p.Match(tblangParserLBRACE)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tblangParserIDENTIFIER {
		{
			p.SetState(42)
			p.Property()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(tblangParserRBRACE)
		if p.HasError() {

			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IVariableDeclarationContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	DECLARE() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	IsVariableDeclarationContext()
}

type VariableDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDeclarationContext() *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_variableDeclaration
	return p
}

func InitEmptyVariableDeclarationContext(p *VariableDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_variableDeclaration
}

func (*VariableDeclarationContext) IsVariableDeclarationContext() {}

func NewVariableDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDeclarationContext {
	var p = new(VariableDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_variableDeclaration

	return p
}

func (s *VariableDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDeclarationContext) DECLARE() antlr.TerminalNode {
	return s.GetToken(tblangParserDECLARE, 0)
}

func (s *VariableDeclarationContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *VariableDeclarationContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(tblangParserASSIGN, 0)
}

func (s *VariableDeclarationContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableDeclarationContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(tblangParserSEMICOLON, 0)
}

func (s *VariableDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitVariableDeclaration(s)
	}
}

func (s *VariableDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitVariableDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) VariableDeclaration() (localctx IVariableDeclarationContext) {
	localctx = NewVariableDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, tblangParserRULE_variableDeclaration)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(50)
		p.Match(tblangParserDECLARE)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(52)
		p.Match(tblangParserASSIGN)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(53)
		p.expression(0)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(54)
			p.Match(tblangParserSEMICOLON)
			if p.HasError() {

				goto errorExit
			}
		}

	} else if p.HasError() {
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IForLoopContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	FOR() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	IN() antlr.TerminalNode
	Expression() IExpressionContext
	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	IsForLoopContext()
}

type ForLoopContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForLoopContext() *ForLoopContext {
	var p = new(ForLoopContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_forLoop
	return p
}

func InitEmptyForLoopContext(p *ForLoopContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_forLoop
}

func (*ForLoopContext) IsForLoopContext() {}

func NewForLoopContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForLoopContext {
	var p = new(ForLoopContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_forLoop

	return p
}

func (s *ForLoopContext) GetParser() antlr.Parser { return s.parser }

func (s *ForLoopContext) FOR() antlr.TerminalNode {
	return s.GetToken(tblangParserFOR, 0)
}

func (s *ForLoopContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *ForLoopContext) IN() antlr.TerminalNode {
	return s.GetToken(tblangParserIN, 0)
}

func (s *ForLoopContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForLoopContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserLBRACE, 0)
}

func (s *ForLoopContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserRBRACE, 0)
}

func (s *ForLoopContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ForLoopContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ForLoopContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForLoopContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForLoopContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterForLoop(s)
	}
}

func (s *ForLoopContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitForLoop(s)
	}
}

func (s *ForLoopContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitForLoop(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) ForLoop() (localctx IForLoopContext) {
	localctx = NewForLoopContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, tblangParserRULE_forLoop)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(57)
		p.Match(tblangParserFOR)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(58)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(59)
		p.Match(tblangParserIN)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(60)
		p.expression(0)
	}
	{
		p.SetState(61)
		p.Match(tblangParserLBRACE)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(65)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1158) != 0 {
		{
			p.SetState(62)
			p.Statement()
		}

		p.SetState(67)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(68)
		p.Match(tblangParserRBRACE)
		if p.HasError() {

			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IPropertyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IDENTIFIER() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext
	SEMICOLON() antlr.TerminalNode

	IsPropertyContext()
}

type PropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_property
	return p
}

func InitEmptyPropertyContext(p *PropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_property
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *PropertyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(tblangParserASSIGN, 0)
}

func (s *PropertyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PropertyContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(tblangParserSEMICOLON, 0)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, tblangParserRULE_property)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(71)
		p.Match(tblangParserASSIGN)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.expression(0)
	}
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tblangParserSEMICOLON {
		{
			p.SetState(73)
			p.Match(tblangParserSEMICOLON)
			if p.HasError() {

				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IFunctionCallContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IDENTIFIER() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	RPAREN() antlr.TerminalNode
	ArgumentList() IArgumentListContext
	SEMICOLON() antlr.TerminalNode

	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *FunctionCallContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tblangParserLPAREN, 0)
}

func (s *FunctionCallContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tblangParserRPAREN, 0)
}

func (s *FunctionCallContext) ArgumentList() IArgumentListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgumentListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgumentListContext)
}

func (s *FunctionCallContext) SEMICOLON() antlr.TerminalNode {
	return s.GetToken(tblangParserSEMICOLON, 0)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, tblangParserRULE_functionCall)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(76)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(77)
		p.Match(tblangParserLPAREN)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(79)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&172272) != 0 {
		{
			p.SetState(78)
			p.ArgumentList()
		}

	}
	{
		p.SetState(81)
		p.Match(tblangParserRPAREN)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(83)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 7, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(82)
			p.Match(tblangParserSEMICOLON)
			if p.HasError() {

				goto errorExit
			}
		}

	} else if p.HasError() {
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IArgumentListContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	IsArgumentListContext()
}

type ArgumentListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentListContext() *ArgumentListContext {
	var p = new(ArgumentListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_argumentList
	return p
}

func InitEmptyArgumentListContext(p *ArgumentListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_argumentList
}

func (*ArgumentListContext) IsArgumentListContext() {}

func NewArgumentListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentListContext {
	var p = new(ArgumentListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_argumentList

	return p
}

func (s *ArgumentListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgumentListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgumentListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tblangParserCOMMA)
}

func (s *ArgumentListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tblangParserCOMMA, i)
}

func (s *ArgumentListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterArgumentList(s)
	}
}

func (s *ArgumentListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitArgumentList(s)
	}
}

func (s *ArgumentListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitArgumentList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) ArgumentList() (localctx IArgumentListContext) {
	localctx = NewArgumentListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, tblangParserRULE_argumentList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(85)
		p.expression(0)
	}
	p.SetState(90)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tblangParserCOMMA {
		{
			p.SetState(86)
			p.Match(tblangParserCOMMA)
			if p.HasError() {

				goto errorExit
			}
		}
		{
			p.SetState(87)
			p.expression(0)
		}

		p.SetState(92)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IExpressionContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	STRING_LITERAL() antlr.TerminalNode
	NUMBER() antlr.TerminalNode
	BOOLEAN() antlr.TerminalNode
	IDENTIFIER() antlr.TerminalNode
	ObjectLiteral() IObjectLiteralContext
	ArrayLiteral() IArrayLiteralContext
	FunctionCall() IFunctionCallContext
	LPAREN() antlr.TerminalNode
	Expression() IExpressionContext
	RPAREN() antlr.TerminalNode
	DOT() antlr.TerminalNode

	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(tblangParserSTRING_LITERAL, 0)
}

func (s *ExpressionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(tblangParserNUMBER, 0)
}

func (s *ExpressionContext) BOOLEAN() antlr.TerminalNode {
	return s.GetToken(tblangParserBOOLEAN, 0)
}

func (s *ExpressionContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *ExpressionContext) ObjectLiteral() IObjectLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectLiteralContext)
}

func (s *ExpressionContext) ArrayLiteral() IArrayLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArrayLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArrayLiteralContext)
}

func (s *ExpressionContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(tblangParserLPAREN, 0)
}

func (s *ExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(tblangParserRPAREN, 0)
}

func (s *ExpressionContext) DOT() antlr.TerminalNode {
	return s.GetToken(tblangParserDOT, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *tblangParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx
	_startState := 16
	p.EnterRecursionRule(localctx, 16, tblangParserRULE_expression, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(105)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(94)
			p.Match(tblangParserSTRING_LITERAL)
			if p.HasError() {

				goto errorExit
			}
		}

	case 2:
		{
			p.SetState(95)
			p.Match(tblangParserNUMBER)
			if p.HasError() {

				goto errorExit
			}
		}

	case 3:
		{
			p.SetState(96)
			p.Match(tblangParserBOOLEAN)
			if p.HasError() {

				goto errorExit
			}
		}

	case 4:
		{
			p.SetState(97)
			p.Match(tblangParserIDENTIFIER)
			if p.HasError() {

				goto errorExit
			}
		}

	case 5:
		{
			p.SetState(98)
			p.ObjectLiteral()
		}

	case 6:
		{
			p.SetState(99)
			p.ArrayLiteral()
		}

	case 7:
		{
			p.SetState(100)
			p.FunctionCall()
		}

	case 8:
		{
			p.SetState(101)
			p.Match(tblangParserLPAREN)
			if p.HasError() {

				goto errorExit
			}
		}
		{
			p.SetState(102)
			p.expression(0)
		}
		{
			p.SetState(103)
			p.Match(tblangParserRPAREN)
			if p.HasError() {

				goto errorExit
			}
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(112)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, tblangParserRULE_expression)
			p.SetState(107)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				goto errorExit
			}
			{
				p.SetState(108)
				p.Match(tblangParserDOT)
				if p.HasError() {

					goto errorExit
				}
			}
			{
				p.SetState(109)
				p.Match(tblangParserIDENTIFIER)
				if p.HasError() {

					goto errorExit
				}
			}

		}
		p.SetState(114)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit
}

type IObjectLiteralContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	LBRACE() antlr.TerminalNode
	RBRACE() antlr.TerminalNode
	AllObjectProperty() []IObjectPropertyContext
	ObjectProperty(i int) IObjectPropertyContext

	IsObjectLiteralContext()
}

type ObjectLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectLiteralContext() *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_objectLiteral
	return p
}

func InitEmptyObjectLiteralContext(p *ObjectLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_objectLiteral
}

func (*ObjectLiteralContext) IsObjectLiteralContext() {}

func NewObjectLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectLiteralContext {
	var p = new(ObjectLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_objectLiteral

	return p
}

func (s *ObjectLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectLiteralContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserLBRACE, 0)
}

func (s *ObjectLiteralContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(tblangParserRBRACE, 0)
}

func (s *ObjectLiteralContext) AllObjectProperty() []IObjectPropertyContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectPropertyContext); ok {
			len++
		}
	}

	tst := make([]IObjectPropertyContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectPropertyContext); ok {
			tst[i] = t.(IObjectPropertyContext)
			i++
		}
	}

	return tst
}

func (s *ObjectLiteralContext) ObjectProperty(i int) IObjectPropertyContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectPropertyContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectPropertyContext)
}

func (s *ObjectLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitObjectLiteral(s)
	}
}

func (s *ObjectLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitObjectLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) ObjectLiteral() (localctx IObjectLiteralContext) {
	localctx = NewObjectLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, tblangParserRULE_objectLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(115)
		p.Match(tblangParserLBRACE)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(119)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == tblangParserIDENTIFIER {
		{
			p.SetState(116)
			p.ObjectProperty()
		}

		p.SetState(121)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(122)
		p.Match(tblangParserRBRACE)
		if p.HasError() {

			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IObjectPropertyContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	IDENTIFIER() antlr.TerminalNode
	Expression() IExpressionContext
	COLON() antlr.TerminalNode
	ASSIGN() antlr.TerminalNode
	COMMA() antlr.TerminalNode

	IsObjectPropertyContext()
}

type ObjectPropertyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectPropertyContext() *ObjectPropertyContext {
	var p = new(ObjectPropertyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_objectProperty
	return p
}

func InitEmptyObjectPropertyContext(p *ObjectPropertyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_objectProperty
}

func (*ObjectPropertyContext) IsObjectPropertyContext() {}

func NewObjectPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectPropertyContext {
	var p = new(ObjectPropertyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_objectProperty

	return p
}

func (s *ObjectPropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectPropertyContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(tblangParserIDENTIFIER, 0)
}

func (s *ObjectPropertyContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ObjectPropertyContext) COLON() antlr.TerminalNode {
	return s.GetToken(tblangParserCOLON, 0)
}

func (s *ObjectPropertyContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(tblangParserASSIGN, 0)
}

func (s *ObjectPropertyContext) COMMA() antlr.TerminalNode {
	return s.GetToken(tblangParserCOMMA, 0)
}

func (s *ObjectPropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectPropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectPropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterObjectProperty(s)
	}
}

func (s *ObjectPropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitObjectProperty(s)
	}
}

func (s *ObjectPropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitObjectProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) ObjectProperty() (localctx IObjectPropertyContext) {
	localctx = NewObjectPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, tblangParserRULE_objectProperty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(124)
		p.Match(tblangParserIDENTIFIER)
		if p.HasError() {

			goto errorExit
		}
	}
	{
		p.SetState(125)
		_la = p.GetTokenStream().LA(1)

		if !(_la == tblangParserASSIGN || _la == tblangParserCOLON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(126)
		p.expression(0)
	}
	p.SetState(128)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == tblangParserCOMMA {
		{
			p.SetState(127)
			p.Match(tblangParserCOMMA)
			if p.HasError() {

				goto errorExit
			}
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

type IArrayLiteralContext interface {
	antlr.ParserRuleContext

	GetParser() antlr.Parser

	LBRACKET() antlr.TerminalNode
	RBRACKET() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	IsArrayLiteralContext()
}

type ArrayLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArrayLiteralContext() *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_arrayLiteral
	return p
}

func InitEmptyArrayLiteralContext(p *ArrayLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = tblangParserRULE_arrayLiteral
}

func (*ArrayLiteralContext) IsArrayLiteralContext() {}

func NewArrayLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayLiteralContext {
	var p = new(ArrayLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = tblangParserRULE_arrayLiteral

	return p
}

func (s *ArrayLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayLiteralContext) LBRACKET() antlr.TerminalNode {
	return s.GetToken(tblangParserLBRACKET, 0)
}

func (s *ArrayLiteralContext) RBRACKET() antlr.TerminalNode {
	return s.GetToken(tblangParserRBRACKET, 0)
}

func (s *ArrayLiteralContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArrayLiteralContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArrayLiteralContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(tblangParserCOMMA)
}

func (s *ArrayLiteralContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(tblangParserCOMMA, i)
}

func (s *ArrayLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.EnterArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(tblangListener); ok {
		listenerT.ExitArrayLiteral(s)
	}
}

func (s *ArrayLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case tblangVisitor:
		return t.VisitArrayLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *tblangParser) ArrayLiteral() (localctx IArrayLiteralContext) {
	localctx = NewArrayLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, tblangParserRULE_arrayLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(130)
		p.Match(tblangParserLBRACKET)
		if p.HasError() {

			goto errorExit
		}
	}
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&172272) != 0 {
		{
			p.SetState(131)
			p.expression(0)
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == tblangParserCOMMA {
			{
				p.SetState(132)
				p.Match(tblangParserCOMMA)
				if p.HasError() {

					goto errorExit
				}
			}
			{
				p.SetState(133)
				p.expression(0)
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(141)
		p.Match(tblangParserRBRACKET)
		if p.HasError() {

			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit
}

func (p *tblangParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 8:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *tblangParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
