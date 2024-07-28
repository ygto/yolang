// Code generated from parser/Lang.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Lang
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type LangParser struct {
	*antlr.BaseParser
}

var LangParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func langParserInit() {
	staticData := &LangParserStaticData
	staticData.LiteralNames = []string{
		"", "'print'", "'int'", "'varchar'", "'$'", "'var'", "'return'", "'#'",
		"','", "", "", "", "", "'null'", "'='", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "", "INT", "VARCHAR", "ID", "WS", "NULL",
		"ASSIGN", "LPAREN", "RPAREN", "LSCOPE", "RSCOPE",
	}
	staticData.RuleNames = []string{
		"prog", "expr", "reservedFunctions", "functions", "types", "valExpr",
		"variableExpr", "varExpr", "returnExpr", "assignExpr", "block", "callFuncExpr",
		"funcExpr", "staticValues", "exprList",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 18, 112, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 1, 0, 4, 0,
		32, 8, 0, 11, 0, 12, 0, 33, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 3, 1, 45, 8, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 4,
		1, 4, 1, 5, 1, 5, 1, 5, 3, 5, 59, 8, 5, 1, 6, 1, 6, 1, 6, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 8, 1, 8, 1, 8, 1, 9, 1, 9, 1, 9, 1, 9, 1, 10, 1, 10, 4, 10,
		77, 8, 10, 11, 10, 12, 10, 78, 1, 10, 1, 10, 1, 10, 1, 10, 3, 10, 85, 8,
		10, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 12, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 5,
		14, 107, 8, 14, 10, 14, 12, 14, 110, 9, 14, 1, 14, 0, 0, 15, 0, 2, 4, 6,
		8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 0, 2, 1, 0, 2, 3, 2, 0, 9, 10,
		13, 13, 108, 0, 31, 1, 0, 0, 0, 2, 44, 1, 0, 0, 0, 4, 46, 1, 0, 0, 0, 6,
		48, 1, 0, 0, 0, 8, 53, 1, 0, 0, 0, 10, 58, 1, 0, 0, 0, 12, 60, 1, 0, 0,
		0, 14, 63, 1, 0, 0, 0, 16, 67, 1, 0, 0, 0, 18, 70, 1, 0, 0, 0, 20, 84,
		1, 0, 0, 0, 22, 86, 1, 0, 0, 0, 24, 92, 1, 0, 0, 0, 26, 101, 1, 0, 0, 0,
		28, 103, 1, 0, 0, 0, 30, 32, 3, 2, 1, 0, 31, 30, 1, 0, 0, 0, 32, 33, 1,
		0, 0, 0, 33, 31, 1, 0, 0, 0, 33, 34, 1, 0, 0, 0, 34, 35, 1, 0, 0, 0, 35,
		36, 5, 0, 0, 1, 36, 1, 1, 0, 0, 0, 37, 45, 3, 18, 9, 0, 38, 45, 3, 24,
		12, 0, 39, 45, 3, 22, 11, 0, 40, 45, 3, 6, 3, 0, 41, 45, 3, 10, 5, 0, 42,
		45, 3, 14, 7, 0, 43, 45, 3, 16, 8, 0, 44, 37, 1, 0, 0, 0, 44, 38, 1, 0,
		0, 0, 44, 39, 1, 0, 0, 0, 44, 40, 1, 0, 0, 0, 44, 41, 1, 0, 0, 0, 44, 42,
		1, 0, 0, 0, 44, 43, 1, 0, 0, 0, 45, 3, 1, 0, 0, 0, 46, 47, 5, 1, 0, 0,
		47, 5, 1, 0, 0, 0, 48, 49, 3, 4, 2, 0, 49, 50, 5, 15, 0, 0, 50, 51, 3,
		10, 5, 0, 51, 52, 5, 16, 0, 0, 52, 7, 1, 0, 0, 0, 53, 54, 7, 0, 0, 0, 54,
		9, 1, 0, 0, 0, 55, 59, 3, 12, 6, 0, 56, 59, 3, 22, 11, 0, 57, 59, 3, 26,
		13, 0, 58, 55, 1, 0, 0, 0, 58, 56, 1, 0, 0, 0, 58, 57, 1, 0, 0, 0, 59,
		11, 1, 0, 0, 0, 60, 61, 5, 4, 0, 0, 61, 62, 5, 11, 0, 0, 62, 13, 1, 0,
		0, 0, 63, 64, 5, 5, 0, 0, 64, 65, 3, 12, 6, 0, 65, 66, 3, 8, 4, 0, 66,
		15, 1, 0, 0, 0, 67, 68, 5, 6, 0, 0, 68, 69, 3, 10, 5, 0, 69, 17, 1, 0,
		0, 0, 70, 71, 3, 12, 6, 0, 71, 72, 5, 14, 0, 0, 72, 73, 3, 10, 5, 0, 73,
		19, 1, 0, 0, 0, 74, 76, 5, 17, 0, 0, 75, 77, 3, 2, 1, 0, 76, 75, 1, 0,
		0, 0, 77, 78, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 78, 79, 1, 0, 0, 0, 79, 80,
		1, 0, 0, 0, 80, 81, 5, 18, 0, 0, 81, 85, 1, 0, 0, 0, 82, 83, 5, 17, 0,
		0, 83, 85, 5, 18, 0, 0, 84, 74, 1, 0, 0, 0, 84, 82, 1, 0, 0, 0, 85, 21,
		1, 0, 0, 0, 86, 87, 5, 7, 0, 0, 87, 88, 5, 11, 0, 0, 88, 89, 5, 15, 0,
		0, 89, 90, 3, 12, 6, 0, 90, 91, 5, 16, 0, 0, 91, 23, 1, 0, 0, 0, 92, 93,
		5, 7, 0, 0, 93, 94, 5, 11, 0, 0, 94, 95, 5, 15, 0, 0, 95, 96, 3, 8, 4,
		0, 96, 97, 3, 12, 6, 0, 97, 98, 5, 16, 0, 0, 98, 99, 3, 8, 4, 0, 99, 100,
		3, 20, 10, 0, 100, 25, 1, 0, 0, 0, 101, 102, 7, 1, 0, 0, 102, 27, 1, 0,
		0, 0, 103, 108, 3, 2, 1, 0, 104, 105, 5, 8, 0, 0, 105, 107, 3, 2, 1, 0,
		106, 104, 1, 0, 0, 0, 107, 110, 1, 0, 0, 0, 108, 106, 1, 0, 0, 0, 108,
		109, 1, 0, 0, 0, 109, 29, 1, 0, 0, 0, 110, 108, 1, 0, 0, 0, 6, 33, 44,
		58, 78, 84, 108,
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

// LangParserInit initializes any static state used to implement LangParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewLangParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func LangParserInit() {
	staticData := &LangParserStaticData
	staticData.once.Do(langParserInit)
}

// NewLangParser produces a new parser instance for the optional input antlr.TokenStream.
func NewLangParser(input antlr.TokenStream) *LangParser {
	LangParserInit()
	this := new(LangParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &LangParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "Lang.g4"

	return this
}

// LangParser tokens.
const (
	LangParserEOF     = antlr.TokenEOF
	LangParserT__0    = 1
	LangParserT__1    = 2
	LangParserT__2    = 3
	LangParserT__3    = 4
	LangParserT__4    = 5
	LangParserT__5    = 6
	LangParserT__6    = 7
	LangParserT__7    = 8
	LangParserINT     = 9
	LangParserVARCHAR = 10
	LangParserID      = 11
	LangParserWS      = 12
	LangParserNULL    = 13
	LangParserASSIGN  = 14
	LangParserLPAREN  = 15
	LangParserRPAREN  = 16
	LangParserLSCOPE  = 17
	LangParserRSCOPE  = 18
)

// LangParser rules.
const (
	LangParserRULE_prog              = 0
	LangParserRULE_expr              = 1
	LangParserRULE_reservedFunctions = 2
	LangParserRULE_functions         = 3
	LangParserRULE_types             = 4
	LangParserRULE_valExpr           = 5
	LangParserRULE_variableExpr      = 6
	LangParserRULE_varExpr           = 7
	LangParserRULE_returnExpr        = 8
	LangParserRULE_assignExpr        = 9
	LangParserRULE_block             = 10
	LangParserRULE_callFuncExpr      = 11
	LangParserRULE_funcExpr          = 12
	LangParserRULE_staticValues      = 13
	LangParserRULE_exprList          = 14
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_prog
	return p
}

func InitEmptyProgContext(p *ProgContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_prog
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) EOF() antlr.TerminalNode {
	return s.GetToken(LangParserEOF, 0)
}

func (s *ProgContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ProgContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitProg(s)
	}
}

func (p *LangParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LangParserRULE_prog)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(31)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9970) != 0) {
		{
			p.SetState(30)
			p.Expr()
		}

		p.SetState(33)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(35)
		p.Match(LangParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignExpr() IAssignExprContext
	FuncExpr() IFuncExprContext
	CallFuncExpr() ICallFuncExprContext
	Functions() IFunctionsContext
	ValExpr() IValExprContext
	VarExpr() IVarExprContext
	ReturnExpr() IReturnExprContext

	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) AssignExpr() IAssignExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignExprContext)
}

func (s *ExprContext) FuncExpr() IFuncExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncExprContext)
}

func (s *ExprContext) CallFuncExpr() ICallFuncExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFuncExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFuncExprContext)
}

func (s *ExprContext) Functions() IFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionsContext)
}

func (s *ExprContext) ValExpr() IValExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValExprContext)
}

func (s *ExprContext) VarExpr() IVarExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarExprContext)
}

func (s *ExprContext) ReturnExpr() IReturnExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnExprContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterExpr(s)
	}
}

func (s *ExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitExpr(s)
	}
}

func (p *LangParser) Expr() (localctx IExprContext) {
	localctx = NewExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LangParserRULE_expr)
	p.SetState(44)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(37)
			p.AssignExpr()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(38)
			p.FuncExpr()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(39)
			p.CallFuncExpr()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(40)
			p.Functions()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(41)
			p.ValExpr()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(42)
			p.VarExpr()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(43)
			p.ReturnExpr()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReservedFunctionsContext is an interface to support dynamic dispatch.
type IReservedFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsReservedFunctionsContext differentiates from other interfaces.
	IsReservedFunctionsContext()
}

type ReservedFunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReservedFunctionsContext() *ReservedFunctionsContext {
	var p = new(ReservedFunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_reservedFunctions
	return p
}

func InitEmptyReservedFunctionsContext(p *ReservedFunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_reservedFunctions
}

func (*ReservedFunctionsContext) IsReservedFunctionsContext() {}

func NewReservedFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReservedFunctionsContext {
	var p = new(ReservedFunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_reservedFunctions

	return p
}

func (s *ReservedFunctionsContext) GetParser() antlr.Parser { return s.parser }
func (s *ReservedFunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReservedFunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReservedFunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterReservedFunctions(s)
	}
}

func (s *ReservedFunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitReservedFunctions(s)
	}
}

func (p *LangParser) ReservedFunctions() (localctx IReservedFunctionsContext) {
	localctx = NewReservedFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LangParserRULE_reservedFunctions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(46)
		p.Match(LangParserT__0)
		if p.HasError() {
			// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionsContext is an interface to support dynamic dispatch.
type IFunctionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ReservedFunctions() IReservedFunctionsContext
	LPAREN() antlr.TerminalNode
	ValExpr() IValExprContext
	RPAREN() antlr.TerminalNode

	// IsFunctionsContext differentiates from other interfaces.
	IsFunctionsContext()
}

type FunctionsContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionsContext() *FunctionsContext {
	var p = new(FunctionsContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_functions
	return p
}

func InitEmptyFunctionsContext(p *FunctionsContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_functions
}

func (*FunctionsContext) IsFunctionsContext() {}

func NewFunctionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionsContext {
	var p = new(FunctionsContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_functions

	return p
}

func (s *FunctionsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionsContext) ReservedFunctions() IReservedFunctionsContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReservedFunctionsContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReservedFunctionsContext)
}

func (s *FunctionsContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserLPAREN, 0)
}

func (s *FunctionsContext) ValExpr() IValExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValExprContext)
}

func (s *FunctionsContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserRPAREN, 0)
}

func (s *FunctionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterFunctions(s)
	}
}

func (s *FunctionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitFunctions(s)
	}
}

func (p *LangParser) Functions() (localctx IFunctionsContext) {
	localctx = NewFunctionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LangParserRULE_functions)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(48)
		p.ReservedFunctions()
	}
	{
		p.SetState(49)
		p.Match(LangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.ValExpr()
	}
	{
		p.SetState(51)
		p.Match(LangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_types
	return p
}

func InitEmptyTypesContext(p *TypesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_types
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }
func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *LangParser) Types() (localctx ITypesContext) {
	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LangParserRULE_types)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		_la = p.GetTokenStream().LA(1)

		if !(_la == LangParserT__1 || _la == LangParserT__2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValExprContext is an interface to support dynamic dispatch.
type IValExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableExpr() IVariableExprContext
	CallFuncExpr() ICallFuncExprContext
	StaticValues() IStaticValuesContext

	// IsValExprContext differentiates from other interfaces.
	IsValExprContext()
}

type ValExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValExprContext() *ValExprContext {
	var p = new(ValExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_valExpr
	return p
}

func InitEmptyValExprContext(p *ValExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_valExpr
}

func (*ValExprContext) IsValExprContext() {}

func NewValExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValExprContext {
	var p = new(ValExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_valExpr

	return p
}

func (s *ValExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ValExprContext) VariableExpr() IVariableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableExprContext)
}

func (s *ValExprContext) CallFuncExpr() ICallFuncExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICallFuncExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICallFuncExprContext)
}

func (s *ValExprContext) StaticValues() IStaticValuesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStaticValuesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStaticValuesContext)
}

func (s *ValExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterValExpr(s)
	}
}

func (s *ValExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitValExpr(s)
	}
}

func (p *LangParser) ValExpr() (localctx IValExprContext) {
	localctx = NewValExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LangParserRULE_valExpr)
	p.SetState(58)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case LangParserT__3:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(55)
			p.VariableExpr()
		}

	case LangParserT__6:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.CallFuncExpr()
		}

	case LangParserINT, LangParserVARCHAR, LangParserNULL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(57)
			p.StaticValues()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVariableExprContext is an interface to support dynamic dispatch.
type IVariableExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode

	// IsVariableExprContext differentiates from other interfaces.
	IsVariableExprContext()
}

type VariableExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableExprContext() *VariableExprContext {
	var p = new(VariableExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_variableExpr
	return p
}

func InitEmptyVariableExprContext(p *VariableExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_variableExpr
}

func (*VariableExprContext) IsVariableExprContext() {}

func NewVariableExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableExprContext {
	var p = new(VariableExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_variableExpr

	return p
}

func (s *VariableExprContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LangParserID, 0)
}

func (s *VariableExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterVariableExpr(s)
	}
}

func (s *VariableExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitVariableExpr(s)
	}
}

func (p *LangParser) VariableExpr() (localctx IVariableExprContext) {
	localctx = NewVariableExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LangParserRULE_variableExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(60)
		p.Match(LangParserT__3)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(61)
		p.Match(LangParserID)
		if p.HasError() {
			// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarExprContext is an interface to support dynamic dispatch.
type IVarExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableExpr() IVariableExprContext
	Types() ITypesContext

	// IsVarExprContext differentiates from other interfaces.
	IsVarExprContext()
}

type VarExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarExprContext() *VarExprContext {
	var p = new(VarExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_varExpr
	return p
}

func InitEmptyVarExprContext(p *VarExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_varExpr
}

func (*VarExprContext) IsVarExprContext() {}

func NewVarExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarExprContext {
	var p = new(VarExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_varExpr

	return p
}

func (s *VarExprContext) GetParser() antlr.Parser { return s.parser }

func (s *VarExprContext) VariableExpr() IVariableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableExprContext)
}

func (s *VarExprContext) Types() ITypesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *VarExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterVarExpr(s)
	}
}

func (s *VarExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitVarExpr(s)
	}
}

func (p *LangParser) VarExpr() (localctx IVarExprContext) {
	localctx = NewVarExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LangParserRULE_varExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(63)
		p.Match(LangParserT__4)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(64)
		p.VariableExpr()
	}
	{
		p.SetState(65)
		p.Types()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnExprContext is an interface to support dynamic dispatch.
type IReturnExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ValExpr() IValExprContext

	// IsReturnExprContext differentiates from other interfaces.
	IsReturnExprContext()
}

type ReturnExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnExprContext() *ReturnExprContext {
	var p = new(ReturnExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_returnExpr
	return p
}

func InitEmptyReturnExprContext(p *ReturnExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_returnExpr
}

func (*ReturnExprContext) IsReturnExprContext() {}

func NewReturnExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnExprContext {
	var p = new(ReturnExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_returnExpr

	return p
}

func (s *ReturnExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnExprContext) ValExpr() IValExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValExprContext)
}

func (s *ReturnExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterReturnExpr(s)
	}
}

func (s *ReturnExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitReturnExpr(s)
	}
}

func (p *LangParser) ReturnExpr() (localctx IReturnExprContext) {
	localctx = NewReturnExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LangParserRULE_returnExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(67)
		p.Match(LangParserT__5)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(68)
		p.ValExpr()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignExprContext is an interface to support dynamic dispatch.
type IAssignExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VariableExpr() IVariableExprContext
	ASSIGN() antlr.TerminalNode
	ValExpr() IValExprContext

	// IsAssignExprContext differentiates from other interfaces.
	IsAssignExprContext()
}

type AssignExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignExprContext() *AssignExprContext {
	var p = new(AssignExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_assignExpr
	return p
}

func InitEmptyAssignExprContext(p *AssignExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_assignExpr
}

func (*AssignExprContext) IsAssignExprContext() {}

func NewAssignExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignExprContext {
	var p = new(AssignExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_assignExpr

	return p
}

func (s *AssignExprContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignExprContext) VariableExpr() IVariableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableExprContext)
}

func (s *AssignExprContext) ASSIGN() antlr.TerminalNode {
	return s.GetToken(LangParserASSIGN, 0)
}

func (s *AssignExprContext) ValExpr() IValExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValExprContext)
}

func (s *AssignExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterAssignExpr(s)
	}
}

func (s *AssignExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitAssignExpr(s)
	}
}

func (p *LangParser) AssignExpr() (localctx IAssignExprContext) {
	localctx = NewAssignExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LangParserRULE_assignExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(70)
		p.VariableExpr()
	}
	{
		p.SetState(71)
		p.Match(LangParserASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(72)
		p.ValExpr()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LSCOPE() antlr.TerminalNode
	RSCOPE() antlr.TerminalNode
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LSCOPE() antlr.TerminalNode {
	return s.GetToken(LangParserLSCOPE, 0)
}

func (s *BlockContext) RSCOPE() antlr.TerminalNode {
	return s.GetToken(LangParserRSCOPE, 0)
}

func (s *BlockContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *LangParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LangParserRULE_block)
	var _la int

	p.SetState(84)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(74)
			p.Match(LangParserLSCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(76)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9970) != 0) {
			{
				p.SetState(75)
				p.Expr()
			}

			p.SetState(78)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(80)
			p.Match(LangParserRSCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(LangParserLSCOPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(83)
			p.Match(LangParserRSCOPE)
			if p.HasError() {
				// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICallFuncExprContext is an interface to support dynamic dispatch.
type ICallFuncExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	VariableExpr() IVariableExprContext
	RPAREN() antlr.TerminalNode

	// IsCallFuncExprContext differentiates from other interfaces.
	IsCallFuncExprContext()
}

type CallFuncExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCallFuncExprContext() *CallFuncExprContext {
	var p = new(CallFuncExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_callFuncExpr
	return p
}

func InitEmptyCallFuncExprContext(p *CallFuncExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_callFuncExpr
}

func (*CallFuncExprContext) IsCallFuncExprContext() {}

func NewCallFuncExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CallFuncExprContext {
	var p = new(CallFuncExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_callFuncExpr

	return p
}

func (s *CallFuncExprContext) GetParser() antlr.Parser { return s.parser }

func (s *CallFuncExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LangParserID, 0)
}

func (s *CallFuncExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserLPAREN, 0)
}

func (s *CallFuncExprContext) VariableExpr() IVariableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableExprContext)
}

func (s *CallFuncExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserRPAREN, 0)
}

func (s *CallFuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CallFuncExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CallFuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterCallFuncExpr(s)
	}
}

func (s *CallFuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitCallFuncExpr(s)
	}
}

func (p *LangParser) CallFuncExpr() (localctx ICallFuncExprContext) {
	localctx = NewCallFuncExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LangParserRULE_callFuncExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(86)
		p.Match(LangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(87)
		p.Match(LangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(88)
		p.Match(LangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(89)
		p.VariableExpr()
	}
	{
		p.SetState(90)
		p.Match(LangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFuncExprContext is an interface to support dynamic dispatch.
type IFuncExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	LPAREN() antlr.TerminalNode
	AllTypes() []ITypesContext
	Types(i int) ITypesContext
	VariableExpr() IVariableExprContext
	RPAREN() antlr.TerminalNode
	Block() IBlockContext

	// IsFuncExprContext differentiates from other interfaces.
	IsFuncExprContext()
}

type FuncExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncExprContext() *FuncExprContext {
	var p = new(FuncExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_funcExpr
	return p
}

func InitEmptyFuncExprContext(p *FuncExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_funcExpr
}

func (*FuncExprContext) IsFuncExprContext() {}

func NewFuncExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncExprContext {
	var p = new(FuncExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_funcExpr

	return p
}

func (s *FuncExprContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncExprContext) ID() antlr.TerminalNode {
	return s.GetToken(LangParserID, 0)
}

func (s *FuncExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserLPAREN, 0)
}

func (s *FuncExprContext) AllTypes() []ITypesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITypesContext); ok {
			len++
		}
	}

	tst := make([]ITypesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITypesContext); ok {
			tst[i] = t.(ITypesContext)
			i++
		}
	}

	return tst
}

func (s *FuncExprContext) Types(i int) ITypesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypesContext); ok {
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

	return t.(ITypesContext)
}

func (s *FuncExprContext) VariableExpr() IVariableExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVariableExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVariableExprContext)
}

func (s *FuncExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LangParserRPAREN, 0)
}

func (s *FuncExprContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FuncExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterFuncExpr(s)
	}
}

func (s *FuncExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitFuncExpr(s)
	}
}

func (p *LangParser) FuncExpr() (localctx IFuncExprContext) {
	localctx = NewFuncExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LangParserRULE_funcExpr)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(LangParserT__6)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(93)
		p.Match(LangParserID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(94)
		p.Match(LangParserLPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(95)
		p.Types()
	}
	{
		p.SetState(96)
		p.VariableExpr()
	}
	{
		p.SetState(97)
		p.Match(LangParserRPAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Types()
	}
	{
		p.SetState(99)
		p.Block()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStaticValuesContext is an interface to support dynamic dispatch.
type IStaticValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INT() antlr.TerminalNode
	VARCHAR() antlr.TerminalNode
	NULL() antlr.TerminalNode

	// IsStaticValuesContext differentiates from other interfaces.
	IsStaticValuesContext()
}

type StaticValuesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStaticValuesContext() *StaticValuesContext {
	var p = new(StaticValuesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_staticValues
	return p
}

func InitEmptyStaticValuesContext(p *StaticValuesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_staticValues
}

func (*StaticValuesContext) IsStaticValuesContext() {}

func NewStaticValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StaticValuesContext {
	var p = new(StaticValuesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_staticValues

	return p
}

func (s *StaticValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *StaticValuesContext) INT() antlr.TerminalNode {
	return s.GetToken(LangParserINT, 0)
}

func (s *StaticValuesContext) VARCHAR() antlr.TerminalNode {
	return s.GetToken(LangParserVARCHAR, 0)
}

func (s *StaticValuesContext) NULL() antlr.TerminalNode {
	return s.GetToken(LangParserNULL, 0)
}

func (s *StaticValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StaticValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StaticValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterStaticValues(s)
	}
}

func (s *StaticValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitStaticValues(s)
	}
}

func (p *LangParser) StaticValues() (localctx IStaticValuesContext) {
	localctx = NewStaticValuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LangParserRULE_staticValues)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(101)
		_la = p.GetTokenStream().LA(1)

		if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&9728) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprListContext is an interface to support dynamic dispatch.
type IExprListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpr() []IExprContext
	Expr(i int) IExprContext

	// IsExprListContext differentiates from other interfaces.
	IsExprListContext()
}

type ExprListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprListContext() *ExprListContext {
	var p = new(ExprListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_exprList
	return p
}

func InitEmptyExprListContext(p *ExprListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = LangParserRULE_exprList
}

func (*ExprListContext) IsExprListContext() {}

func NewExprListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprListContext {
	var p = new(ExprListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = LangParserRULE_exprList

	return p
}

func (s *ExprListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprListContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *ExprListContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
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

	return t.(IExprContext)
}

func (s *ExprListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExprListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.EnterExprList(s)
	}
}

func (s *ExprListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LangListener); ok {
		listenerT.ExitExprList(s)
	}
}

func (p *LangParser) ExprList() (localctx IExprListContext) {
	localctx = NewExprListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LangParserRULE_exprList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(103)
		p.Expr()
	}
	p.SetState(108)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == LangParserT__7 {
		{
			p.SetState(104)
			p.Match(LangParserT__7)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(105)
			p.Expr()
		}

		p.SetState(110)
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
	goto errorExit // Trick to prevent compiler error if the label is not used
}
