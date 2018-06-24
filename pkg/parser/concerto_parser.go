// Code generated from Concerto.G4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Concerto

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 16, 54, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 3, 2, 6, 2, 18, 10, 2, 13, 2, 14, 2, 19, 3, 3, 3, 3, 5, 3, 24,
	10, 3, 3, 4, 3, 4, 5, 4, 28, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 7, 6, 40, 10, 6, 12, 6, 14, 6, 43, 11, 6, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 2, 2, 9, 2, 4, 6,
	8, 10, 12, 14, 2, 3, 3, 2, 6, 7, 2, 50, 2, 17, 3, 2, 2, 2, 4, 23, 3, 2,
	2, 2, 6, 27, 3, 2, 2, 2, 8, 29, 3, 2, 2, 2, 10, 33, 3, 2, 2, 2, 12, 47,
	3, 2, 2, 2, 14, 49, 3, 2, 2, 2, 16, 18, 5, 4, 3, 2, 17, 16, 3, 2, 2, 2,
	18, 19, 3, 2, 2, 2, 19, 17, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 3, 3, 2,
	2, 2, 21, 24, 5, 6, 4, 2, 22, 24, 5, 14, 8, 2, 23, 21, 3, 2, 2, 2, 23,
	22, 3, 2, 2, 2, 24, 5, 3, 2, 2, 2, 25, 28, 5, 8, 5, 2, 26, 28, 5, 10, 6,
	2, 27, 25, 3, 2, 2, 2, 27, 26, 3, 2, 2, 2, 28, 7, 3, 2, 2, 2, 29, 30, 7,
	6, 2, 2, 30, 31, 7, 6, 2, 2, 31, 32, 7, 9, 2, 2, 32, 9, 3, 2, 2, 2, 33,
	34, 7, 6, 2, 2, 34, 35, 7, 6, 2, 2, 35, 36, 7, 3, 2, 2, 36, 41, 5, 12,
	7, 2, 37, 38, 7, 4, 2, 2, 38, 40, 5, 12, 7, 2, 39, 37, 3, 2, 2, 2, 40,
	43, 3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 44, 3, 2, 2,
	2, 43, 41, 3, 2, 2, 2, 44, 45, 7, 5, 2, 2, 45, 46, 7, 9, 2, 2, 46, 11,
	3, 2, 2, 2, 47, 48, 9, 2, 2, 2, 48, 13, 3, 2, 2, 2, 49, 50, 7, 6, 2, 2,
	50, 51, 7, 6, 2, 2, 51, 52, 7, 9, 2, 2, 52, 15, 3, 2, 2, 2, 6, 19, 23,
	27, 41,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'('", "','", "')'",
}
var symbolicNames = []string{
	"", "", "", "", "IDENTIFIER", "STRING_LIT", "WS", "NEWLINE", "LITTLE_U_VALUE",
	"BIG_U_VALUE", "BINARY_OP", "INT_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "RUNE_LIT",
}

var ruleNames = []string{
	"prog", "statement", "declaration", "simpleDeclaration", "declarationConstructor",
	"arg", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ConcertoParser struct {
	*antlr.BaseParser
}

func NewConcertoParser(input antlr.TokenStream) *ConcertoParser {
	this := new(ConcertoParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Concerto.G4"

	return this
}

// ConcertoParser tokens.
const (
	ConcertoParserEOF            = antlr.TokenEOF
	ConcertoParserT__0           = 1
	ConcertoParserT__1           = 2
	ConcertoParserT__2           = 3
	ConcertoParserIDENTIFIER     = 4
	ConcertoParserSTRING_LIT     = 5
	ConcertoParserWS             = 6
	ConcertoParserNEWLINE        = 7
	ConcertoParserLITTLE_U_VALUE = 8
	ConcertoParserBIG_U_VALUE    = 9
	ConcertoParserBINARY_OP      = 10
	ConcertoParserINT_LIT        = 11
	ConcertoParserFLOAT_LIT      = 12
	ConcertoParserIMAGINARY_LIT  = 13
	ConcertoParserRUNE_LIT       = 14
)

// ConcertoParser rules.
const (
	ConcertoParserRULE_prog                   = 0
	ConcertoParserRULE_statement              = 1
	ConcertoParserRULE_declaration            = 2
	ConcertoParserRULE_simpleDeclaration      = 3
	ConcertoParserRULE_declarationConstructor = 4
	ConcertoParserRULE_arg                    = 5
	ConcertoParserRULE_expression             = 6
)

// IProgContext is an interface to support dynamic dispatch.
type IProgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsProgContext differentiates from other interfaces.
	IsProgContext()
}

type ProgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgContext() *ProgContext {
	var p = new(ProgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_prog
	return p
}

func (*ProgContext) IsProgContext() {}

func NewProgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgContext {
	var p = new(ProgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_prog

	return p
}

func (s *ProgContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ProgContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterProg(s)
	}
}

func (s *ProgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitProg(s)
	}
}

func (s *ProgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitProg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Prog() (localctx IProgContext) {
	localctx = NewProgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ConcertoParserRULE_prog)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == ConcertoParserIDENTIFIER {
		{
			p.SetState(14)
			p.Statement()
		}

		p.SetState(17)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Declaration() IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ConcertoParserRULE_statement)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
			p.Declaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
			p.Expression()
		}

	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) SimpleDeclaration() ISimpleDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISimpleDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISimpleDeclarationContext)
}

func (s *DeclarationContext) DeclarationConstructor() IDeclarationConstructorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationConstructorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclarationConstructorContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ConcertoParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(25)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.SimpleDeclaration()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.DeclarationConstructor()
		}

	}

	return localctx
}

// ISimpleDeclarationContext is an interface to support dynamic dispatch.
type ISimpleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSimpleDeclarationContext differentiates from other interfaces.
	IsSimpleDeclarationContext()
}

type SimpleDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySimpleDeclarationContext() *SimpleDeclarationContext {
	var p = new(SimpleDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_simpleDeclaration
	return p
}

func (*SimpleDeclarationContext) IsSimpleDeclarationContext() {}

func NewSimpleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SimpleDeclarationContext {
	var p = new(SimpleDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_simpleDeclaration

	return p
}

func (s *SimpleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *SimpleDeclarationContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *SimpleDeclarationContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *SimpleDeclarationContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, 0)
}

func (s *SimpleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SimpleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SimpleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterSimpleDeclaration(s)
	}
}

func (s *SimpleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitSimpleDeclaration(s)
	}
}

func (s *SimpleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitSimpleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) SimpleDeclaration() (localctx ISimpleDeclarationContext) {
	localctx = NewSimpleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ConcertoParserRULE_simpleDeclaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(27)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(28)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(29)
		p.Match(ConcertoParserNEWLINE)
	}

	return localctx
}

// IDeclarationConstructorContext is an interface to support dynamic dispatch.
type IDeclarationConstructorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDeclarationConstructorContext differentiates from other interfaces.
	IsDeclarationConstructorContext()
}

type DeclarationConstructorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclarationConstructorContext() *DeclarationConstructorContext {
	var p = new(DeclarationConstructorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_declarationConstructor
	return p
}

func (*DeclarationConstructorContext) IsDeclarationConstructorContext() {}

func NewDeclarationConstructorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationConstructorContext {
	var p = new(DeclarationConstructorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_declarationConstructor

	return p
}

func (s *DeclarationConstructorContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationConstructorContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *DeclarationConstructorContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *DeclarationConstructorContext) AllArg() []IArgContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgContext)(nil)).Elem())
	var tst = make([]IArgContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgContext)
		}
	}

	return tst
}

func (s *DeclarationConstructorContext) Arg(i int) IArgContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgContext)
}

func (s *DeclarationConstructorContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, 0)
}

func (s *DeclarationConstructorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationConstructorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationConstructorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterDeclarationConstructor(s)
	}
}

func (s *DeclarationConstructorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitDeclarationConstructor(s)
	}
}

func (s *DeclarationConstructorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitDeclarationConstructor(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) DeclarationConstructor() (localctx IDeclarationConstructorContext) {
	localctx = NewDeclarationConstructorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ConcertoParserRULE_declarationConstructor)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(31)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(32)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(33)
		p.Match(ConcertoParserT__0)
	}
	{
		p.SetState(34)
		p.Arg()
	}
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == ConcertoParserT__1 {
		{
			p.SetState(35)
			p.Match(ConcertoParserT__1)
		}
		{
			p.SetState(36)
			p.Arg()
		}

		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(42)
		p.Match(ConcertoParserT__2)
	}
	{
		p.SetState(43)
		p.Match(ConcertoParserNEWLINE)
	}

	return localctx
}

// IArgContext is an interface to support dynamic dispatch.
type IArgContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgContext differentiates from other interfaces.
	IsArgContext()
}

type ArgContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgContext() *ArgContext {
	var p = new(ArgContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_arg
	return p
}

func (*ArgContext) IsArgContext() {}

func NewArgContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgContext {
	var p = new(ArgContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_arg

	return p
}

func (s *ArgContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgContext) IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, 0)
}

func (s *ArgContext) STRING_LIT() antlr.TerminalNode {
	return s.GetToken(ConcertoParserSTRING_LIT, 0)
}

func (s *ArgContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterArg(s)
	}
}

func (s *ArgContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitArg(s)
	}
}

func (s *ArgContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitArg(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Arg() (localctx IArgContext) {
	localctx = NewArgContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ConcertoParserRULE_arg)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(45)
		_la = p.GetTokenStream().LA(1)

		if !(_la == ConcertoParserIDENTIFIER || _la == ConcertoParserSTRING_LIT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ConcertoParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ConcertoParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AllIDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(ConcertoParserIDENTIFIER)
}

func (s *ExpressionContext) IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(ConcertoParserIDENTIFIER, i)
}

func (s *ExpressionContext) NEWLINE() antlr.TerminalNode {
	return s.GetToken(ConcertoParserNEWLINE, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ConcertoListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ConcertoVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ConcertoParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ConcertoParserRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(48)
		p.Match(ConcertoParserIDENTIFIER)
	}
	{
		p.SetState(49)
		p.Match(ConcertoParserNEWLINE)
	}

	return localctx
}
