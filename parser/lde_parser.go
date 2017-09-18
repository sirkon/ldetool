// Generated from LDE.g4 by ANTLR 4.7.

package parser // LDE

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 21, 187,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 7, 2, 54, 10, 2, 12, 2, 14,
	2, 57, 11, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 79,
	10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 5, 5, 94, 10, 5, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	166, 10, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24,
	3, 24, 2, 3, 2, 25, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 2, 3, 3, 2, 17, 18, 2, 184, 2, 48,
	3, 2, 2, 2, 4, 58, 3, 2, 2, 2, 6, 78, 3, 2, 2, 2, 8, 93, 3, 2, 2, 2, 10,
	95, 3, 2, 2, 2, 12, 98, 3, 2, 2, 2, 14, 101, 3, 2, 2, 2, 16, 105, 3, 2,
	2, 2, 18, 109, 3, 2, 2, 2, 20, 115, 3, 2, 2, 2, 22, 118, 3, 2, 2, 2, 24,
	122, 3, 2, 2, 2, 26, 128, 3, 2, 2, 2, 28, 135, 3, 2, 2, 2, 30, 140, 3,
	2, 2, 2, 32, 146, 3, 2, 2, 2, 34, 151, 3, 2, 2, 2, 36, 165, 3, 2, 2, 2,
	38, 167, 3, 2, 2, 2, 40, 169, 3, 2, 2, 2, 42, 175, 3, 2, 2, 2, 44, 180,
	3, 2, 2, 2, 46, 184, 3, 2, 2, 2, 48, 49, 8, 2, 1, 2, 49, 50, 5, 4, 3, 2,
	50, 55, 3, 2, 2, 2, 51, 52, 12, 4, 2, 2, 52, 54, 5, 4, 3, 2, 53, 51, 3,
	2, 2, 2, 54, 57, 3, 2, 2, 2, 55, 53, 3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56,
	3, 3, 2, 2, 2, 57, 55, 3, 2, 2, 2, 58, 59, 7, 15, 2, 2, 59, 60, 7, 3, 2,
	2, 60, 61, 5, 6, 4, 2, 61, 62, 7, 4, 2, 2, 62, 5, 3, 2, 2, 2, 63, 64, 7,
	21, 2, 2, 64, 79, 5, 6, 4, 2, 65, 66, 7, 5, 2, 2, 66, 67, 5, 6, 4, 2, 67,
	68, 7, 6, 2, 2, 68, 69, 5, 6, 4, 2, 69, 79, 3, 2, 2, 2, 70, 71, 7, 5, 2,
	2, 71, 72, 5, 6, 4, 2, 72, 73, 7, 6, 2, 2, 73, 79, 3, 2, 2, 2, 74, 75,
	5, 8, 5, 2, 75, 76, 5, 6, 4, 2, 76, 79, 3, 2, 2, 2, 77, 79, 5, 8, 5, 2,
	78, 63, 3, 2, 2, 2, 78, 65, 3, 2, 2, 2, 78, 70, 3, 2, 2, 2, 78, 74, 3,
	2, 2, 2, 78, 77, 3, 2, 2, 2, 79, 7, 3, 2, 2, 2, 80, 94, 5, 10, 6, 2, 81,
	94, 5, 12, 7, 2, 82, 94, 5, 14, 8, 2, 83, 94, 5, 16, 9, 2, 84, 94, 5, 18,
	10, 2, 85, 94, 5, 20, 11, 2, 86, 94, 5, 22, 12, 2, 87, 94, 5, 24, 13, 2,
	88, 94, 5, 26, 14, 2, 89, 94, 5, 28, 15, 2, 90, 94, 5, 30, 16, 2, 91, 94,
	5, 32, 17, 2, 92, 94, 5, 34, 18, 2, 93, 80, 3, 2, 2, 2, 93, 81, 3, 2, 2,
	2, 93, 82, 3, 2, 2, 2, 93, 83, 3, 2, 2, 2, 93, 84, 3, 2, 2, 2, 93, 85,
	3, 2, 2, 2, 93, 86, 3, 2, 2, 2, 93, 87, 3, 2, 2, 2, 93, 88, 3, 2, 2, 2,
	93, 89, 3, 2, 2, 2, 93, 90, 3, 2, 2, 2, 93, 91, 3, 2, 2, 2, 93, 92, 3,
	2, 2, 2, 94, 9, 3, 2, 2, 2, 95, 96, 7, 7, 2, 2, 96, 97, 7, 17, 2, 2, 97,
	11, 3, 2, 2, 2, 98, 99, 7, 7, 2, 2, 99, 100, 7, 18, 2, 2, 100, 13, 3, 2,
	2, 2, 101, 102, 7, 8, 2, 2, 102, 103, 7, 7, 2, 2, 103, 104, 7, 17, 2, 2,
	104, 15, 3, 2, 2, 2, 105, 106, 7, 8, 2, 2, 106, 107, 7, 7, 2, 2, 107, 108,
	7, 18, 2, 2, 108, 17, 3, 2, 2, 2, 109, 110, 7, 9, 2, 2, 110, 111, 7, 10,
	2, 2, 111, 112, 7, 16, 2, 2, 112, 113, 7, 11, 2, 2, 113, 114, 7, 12, 2,
	2, 114, 19, 3, 2, 2, 2, 115, 116, 7, 9, 2, 2, 116, 117, 5, 36, 19, 2, 117,
	21, 3, 2, 2, 2, 118, 119, 7, 8, 2, 2, 119, 120, 7, 9, 2, 2, 120, 121, 5,
	36, 19, 2, 121, 23, 3, 2, 2, 2, 122, 123, 7, 15, 2, 2, 123, 124, 7, 5,
	2, 2, 124, 125, 5, 46, 24, 2, 125, 126, 7, 6, 2, 2, 126, 127, 5, 36, 19,
	2, 127, 25, 3, 2, 2, 2, 128, 129, 7, 15, 2, 2, 129, 130, 7, 5, 2, 2, 130,
	131, 5, 46, 24, 2, 131, 132, 7, 6, 2, 2, 132, 133, 7, 8, 2, 2, 133, 134,
	5, 36, 19, 2, 134, 27, 3, 2, 2, 2, 135, 136, 7, 15, 2, 2, 136, 137, 7,
	5, 2, 2, 137, 138, 5, 46, 24, 2, 138, 139, 7, 6, 2, 2, 139, 29, 3, 2, 2,
	2, 140, 141, 7, 8, 2, 2, 141, 142, 7, 15, 2, 2, 142, 143, 7, 5, 2, 2, 143,
	144, 5, 6, 4, 2, 144, 145, 7, 6, 2, 2, 145, 31, 3, 2, 2, 2, 146, 147, 7,
	8, 2, 2, 147, 148, 7, 5, 2, 2, 148, 149, 5, 6, 4, 2, 149, 150, 7, 6, 2,
	2, 150, 33, 3, 2, 2, 2, 151, 152, 7, 13, 2, 2, 152, 35, 3, 2, 2, 2, 153,
	154, 5, 38, 20, 2, 154, 155, 5, 40, 21, 2, 155, 166, 3, 2, 2, 2, 156, 157,
	5, 38, 20, 2, 157, 158, 5, 42, 22, 2, 158, 166, 3, 2, 2, 2, 159, 160, 5,
	38, 20, 2, 160, 161, 5, 44, 23, 2, 161, 166, 3, 2, 2, 2, 162, 166, 5, 38,
	20, 2, 163, 164, 7, 14, 2, 2, 164, 166, 5, 36, 19, 2, 165, 153, 3, 2, 2,
	2, 165, 156, 3, 2, 2, 2, 165, 159, 3, 2, 2, 2, 165, 162, 3, 2, 2, 2, 165,
	163, 3, 2, 2, 2, 166, 37, 3, 2, 2, 2, 167, 168, 9, 2, 2, 2, 168, 39, 3,
	2, 2, 2, 169, 170, 7, 10, 2, 2, 170, 171, 7, 16, 2, 2, 171, 172, 7, 11,
	2, 2, 172, 173, 7, 16, 2, 2, 173, 174, 7, 12, 2, 2, 174, 41, 3, 2, 2, 2,
	175, 176, 7, 10, 2, 2, 176, 177, 7, 11, 2, 2, 177, 178, 7, 16, 2, 2, 178,
	179, 7, 12, 2, 2, 179, 43, 3, 2, 2, 2, 180, 181, 7, 10, 2, 2, 181, 182,
	7, 16, 2, 2, 182, 183, 7, 12, 2, 2, 183, 45, 3, 2, 2, 2, 184, 185, 7, 15,
	2, 2, 185, 47, 3, 2, 2, 2, 6, 55, 78, 93, 165,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'='", "';'", "'('", "')'", "'^'", "'?'", "'_'", "'['", "':'", "']'",
	"'$'", "'~'", "", "", "", "", "", "", "'!'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "Identifier", "IntLit",
	"StringLit", "CharLit", "WS", "LineComment", "Stress",
}

var ruleNames = []string{
	"rules", "atomicRule", "baseAction", "atomicAction", "passStringPrefix",
	"passCharPrefix", "mayPassStringPrefix", "mayPassCharPrefix", "passChars",
	"passUntil", "mayPassUntil", "takeUntil", "takeUntilOrRest", "takeUntilRest",
	"optionalNamedArea", "optionalArea", "atEnd", "target", "targetLit", "bound",
	"limit", "exact", "fieldType",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LDEParser struct {
	*antlr.BaseParser
}

func NewLDEParser(input antlr.TokenStream) *LDEParser {
	this := new(LDEParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LDE.g4"

	return this
}

// LDEParser tokens.
const (
	LDEParserEOF         = antlr.TokenEOF
	LDEParserT__0        = 1
	LDEParserT__1        = 2
	LDEParserT__2        = 3
	LDEParserT__3        = 4
	LDEParserT__4        = 5
	LDEParserT__5        = 6
	LDEParserT__6        = 7
	LDEParserT__7        = 8
	LDEParserT__8        = 9
	LDEParserT__9        = 10
	LDEParserT__10       = 11
	LDEParserT__11       = 12
	LDEParserIdentifier  = 13
	LDEParserIntLit      = 14
	LDEParserStringLit   = 15
	LDEParserCharLit     = 16
	LDEParserWS          = 17
	LDEParserLineComment = 18
	LDEParserStress      = 19
)

// LDEParser rules.
const (
	LDEParserRULE_rules               = 0
	LDEParserRULE_atomicRule          = 1
	LDEParserRULE_baseAction          = 2
	LDEParserRULE_atomicAction        = 3
	LDEParserRULE_passStringPrefix    = 4
	LDEParserRULE_passCharPrefix      = 5
	LDEParserRULE_mayPassStringPrefix = 6
	LDEParserRULE_mayPassCharPrefix   = 7
	LDEParserRULE_passChars           = 8
	LDEParserRULE_passUntil           = 9
	LDEParserRULE_mayPassUntil        = 10
	LDEParserRULE_takeUntil           = 11
	LDEParserRULE_takeUntilOrRest     = 12
	LDEParserRULE_takeUntilRest       = 13
	LDEParserRULE_optionalNamedArea   = 14
	LDEParserRULE_optionalArea        = 15
	LDEParserRULE_atEnd               = 16
	LDEParserRULE_target              = 17
	LDEParserRULE_targetLit           = 18
	LDEParserRULE_bound               = 19
	LDEParserRULE_limit               = 20
	LDEParserRULE_exact               = 21
	LDEParserRULE_fieldType           = 22
)

// IRulesContext is an interface to support dynamic dispatch.
type IRulesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRulesContext differentiates from other interfaces.
	IsRulesContext()
}

type RulesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRulesContext() *RulesContext {
	var p = new(RulesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_rules
	return p
}

func (*RulesContext) IsRulesContext() {}

func NewRulesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RulesContext {
	var p = new(RulesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_rules

	return p
}

func (s *RulesContext) GetParser() antlr.Parser { return s.parser }

func (s *RulesContext) AtomicRule() IAtomicRuleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomicRuleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomicRuleContext)
}

func (s *RulesContext) Rules() IRulesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRulesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRulesContext)
}

func (s *RulesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RulesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RulesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterRules(s)
	}
}

func (s *RulesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitRules(s)
	}
}

func (p *LDEParser) Rules() (localctx IRulesContext) {
	return p.rules(0)
}

func (p *LDEParser) rules(_p int) (localctx IRulesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewRulesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IRulesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 0
	p.EnterRecursionRule(localctx, 0, LDEParserRULE_rules, _p)

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
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

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(47)
		p.AtomicRule()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(53)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewRulesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, LDEParserRULE_rules)
			p.SetState(49)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(50)
				p.AtomicRule()
			}

		}
		p.SetState(55)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext())
	}

	return localctx
}

// IAtomicRuleContext is an interface to support dynamic dispatch.
type IAtomicRuleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicRuleContext differentiates from other interfaces.
	IsAtomicRuleContext()
}

type AtomicRuleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicRuleContext() *AtomicRuleContext {
	var p = new(AtomicRuleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atomicRule
	return p
}

func (*AtomicRuleContext) IsAtomicRuleContext() {}

func NewAtomicRuleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicRuleContext {
	var p = new(AtomicRuleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atomicRule

	return p
}

func (s *AtomicRuleContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicRuleContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *AtomicRuleContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *AtomicRuleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicRuleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicRuleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtomicRule(s)
	}
}

func (s *AtomicRuleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtomicRule(s)
	}
}

func (p *LDEParser) AtomicRule() (localctx IAtomicRuleContext) {
	localctx = NewAtomicRuleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LDEParserRULE_atomicRule)

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
		p.SetState(56)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(57)
		p.Match(LDEParserT__0)
	}
	{
		p.SetState(58)
		p.BaseAction()
	}
	{
		p.SetState(59)
		p.Match(LDEParserT__1)
	}

	return localctx
}

// IBaseActionContext is an interface to support dynamic dispatch.
type IBaseActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBaseActionContext differentiates from other interfaces.
	IsBaseActionContext()
}

type BaseActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseActionContext() *BaseActionContext {
	var p = new(BaseActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_baseAction
	return p
}

func (*BaseActionContext) IsBaseActionContext() {}

func NewBaseActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseActionContext {
	var p = new(BaseActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_baseAction

	return p
}

func (s *BaseActionContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseActionContext) Stress() antlr.TerminalNode {
	return s.GetToken(LDEParserStress, 0)
}

func (s *BaseActionContext) AllBaseAction() []IBaseActionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBaseActionContext)(nil)).Elem())
	var tst = make([]IBaseActionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBaseActionContext)
		}
	}

	return tst
}

func (s *BaseActionContext) BaseAction(i int) IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *BaseActionContext) AtomicAction() IAtomicActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtomicActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtomicActionContext)
}

func (s *BaseActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterBaseAction(s)
	}
}

func (s *BaseActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitBaseAction(s)
	}
}

func (p *LDEParser) BaseAction() (localctx IBaseActionContext) {
	localctx = NewBaseActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LDEParserRULE_baseAction)

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

	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(LDEParserStress)
		}
		{
			p.SetState(62)
			p.BaseAction()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(LDEParserT__2)
		}
		{
			p.SetState(64)
			p.BaseAction()
		}
		{
			p.SetState(65)
			p.Match(LDEParserT__3)
		}
		{
			p.SetState(66)
			p.BaseAction()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(68)
			p.Match(LDEParserT__2)
		}
		{
			p.SetState(69)
			p.BaseAction()
		}
		{
			p.SetState(70)
			p.Match(LDEParserT__3)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(72)
			p.AtomicAction()
		}
		{
			p.SetState(73)
			p.BaseAction()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(75)
			p.AtomicAction()
		}

	}

	return localctx
}

// IAtomicActionContext is an interface to support dynamic dispatch.
type IAtomicActionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtomicActionContext differentiates from other interfaces.
	IsAtomicActionContext()
}

type AtomicActionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtomicActionContext() *AtomicActionContext {
	var p = new(AtomicActionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atomicAction
	return p
}

func (*AtomicActionContext) IsAtomicActionContext() {}

func NewAtomicActionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtomicActionContext {
	var p = new(AtomicActionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atomicAction

	return p
}

func (s *AtomicActionContext) GetParser() antlr.Parser { return s.parser }

func (s *AtomicActionContext) PassStringPrefix() IPassStringPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassStringPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassStringPrefixContext)
}

func (s *AtomicActionContext) PassCharPrefix() IPassCharPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassCharPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassCharPrefixContext)
}

func (s *AtomicActionContext) MayPassStringPrefix() IMayPassStringPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayPassStringPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayPassStringPrefixContext)
}

func (s *AtomicActionContext) MayPassCharPrefix() IMayPassCharPrefixContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayPassCharPrefixContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayPassCharPrefixContext)
}

func (s *AtomicActionContext) PassChars() IPassCharsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassCharsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassCharsContext)
}

func (s *AtomicActionContext) PassUntil() IPassUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPassUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPassUntilContext)
}

func (s *AtomicActionContext) MayPassUntil() IMayPassUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMayPassUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMayPassUntilContext)
}

func (s *AtomicActionContext) TakeUntil() ITakeUntilContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilContext)
}

func (s *AtomicActionContext) TakeUntilOrRest() ITakeUntilOrRestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilOrRestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilOrRestContext)
}

func (s *AtomicActionContext) TakeUntilRest() ITakeUntilRestContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITakeUntilRestContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITakeUntilRestContext)
}

func (s *AtomicActionContext) OptionalNamedArea() IOptionalNamedAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalNamedAreaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalNamedAreaContext)
}

func (s *AtomicActionContext) OptionalArea() IOptionalAreaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOptionalAreaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOptionalAreaContext)
}

func (s *AtomicActionContext) AtEnd() IAtEndContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAtEndContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAtEndContext)
}

func (s *AtomicActionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtomicActionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtomicActionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtomicAction(s)
	}
}

func (s *AtomicActionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtomicAction(s)
	}
}

func (p *LDEParser) AtomicAction() (localctx IAtomicActionContext) {
	localctx = NewAtomicActionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LDEParserRULE_atomicAction)

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

	p.SetState(91)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(78)
			p.PassStringPrefix()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(79)
			p.PassCharPrefix()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(80)
			p.MayPassStringPrefix()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(81)
			p.MayPassCharPrefix()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(82)
			p.PassChars()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(83)
			p.PassUntil()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(84)
			p.MayPassUntil()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(85)
			p.TakeUntil()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(86)
			p.TakeUntilOrRest()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(87)
			p.TakeUntilRest()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(88)
			p.OptionalNamedArea()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(89)
			p.OptionalArea()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(90)
			p.AtEnd()
		}

	}

	return localctx
}

// IPassStringPrefixContext is an interface to support dynamic dispatch.
type IPassStringPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassStringPrefixContext differentiates from other interfaces.
	IsPassStringPrefixContext()
}

type PassStringPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassStringPrefixContext() *PassStringPrefixContext {
	var p = new(PassStringPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passStringPrefix
	return p
}

func (*PassStringPrefixContext) IsPassStringPrefixContext() {}

func NewPassStringPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassStringPrefixContext {
	var p = new(PassStringPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passStringPrefix

	return p
}

func (s *PassStringPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *PassStringPrefixContext) StringLit() antlr.TerminalNode {
	return s.GetToken(LDEParserStringLit, 0)
}

func (s *PassStringPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassStringPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassStringPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassStringPrefix(s)
	}
}

func (s *PassStringPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassStringPrefix(s)
	}
}

func (p *LDEParser) PassStringPrefix() (localctx IPassStringPrefixContext) {
	localctx = NewPassStringPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LDEParserRULE_passStringPrefix)

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
		p.SetState(93)
		p.Match(LDEParserT__4)
	}
	{
		p.SetState(94)
		p.Match(LDEParserStringLit)
	}

	return localctx
}

// IPassCharPrefixContext is an interface to support dynamic dispatch.
type IPassCharPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassCharPrefixContext differentiates from other interfaces.
	IsPassCharPrefixContext()
}

type PassCharPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassCharPrefixContext() *PassCharPrefixContext {
	var p = new(PassCharPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passCharPrefix
	return p
}

func (*PassCharPrefixContext) IsPassCharPrefixContext() {}

func NewPassCharPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassCharPrefixContext {
	var p = new(PassCharPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passCharPrefix

	return p
}

func (s *PassCharPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *PassCharPrefixContext) CharLit() antlr.TerminalNode {
	return s.GetToken(LDEParserCharLit, 0)
}

func (s *PassCharPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassCharPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassCharPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassCharPrefix(s)
	}
}

func (s *PassCharPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassCharPrefix(s)
	}
}

func (p *LDEParser) PassCharPrefix() (localctx IPassCharPrefixContext) {
	localctx = NewPassCharPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LDEParserRULE_passCharPrefix)

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
		p.SetState(96)
		p.Match(LDEParserT__4)
	}
	{
		p.SetState(97)
		p.Match(LDEParserCharLit)
	}

	return localctx
}

// IMayPassStringPrefixContext is an interface to support dynamic dispatch.
type IMayPassStringPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayPassStringPrefixContext differentiates from other interfaces.
	IsMayPassStringPrefixContext()
}

type MayPassStringPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayPassStringPrefixContext() *MayPassStringPrefixContext {
	var p = new(MayPassStringPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayPassStringPrefix
	return p
}

func (*MayPassStringPrefixContext) IsMayPassStringPrefixContext() {}

func NewMayPassStringPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayPassStringPrefixContext {
	var p = new(MayPassStringPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayPassStringPrefix

	return p
}

func (s *MayPassStringPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *MayPassStringPrefixContext) StringLit() antlr.TerminalNode {
	return s.GetToken(LDEParserStringLit, 0)
}

func (s *MayPassStringPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayPassStringPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayPassStringPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayPassStringPrefix(s)
	}
}

func (s *MayPassStringPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayPassStringPrefix(s)
	}
}

func (p *LDEParser) MayPassStringPrefix() (localctx IMayPassStringPrefixContext) {
	localctx = NewMayPassStringPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LDEParserRULE_mayPassStringPrefix)

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
		p.SetState(99)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(100)
		p.Match(LDEParserT__4)
	}
	{
		p.SetState(101)
		p.Match(LDEParserStringLit)
	}

	return localctx
}

// IMayPassCharPrefixContext is an interface to support dynamic dispatch.
type IMayPassCharPrefixContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayPassCharPrefixContext differentiates from other interfaces.
	IsMayPassCharPrefixContext()
}

type MayPassCharPrefixContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayPassCharPrefixContext() *MayPassCharPrefixContext {
	var p = new(MayPassCharPrefixContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayPassCharPrefix
	return p
}

func (*MayPassCharPrefixContext) IsMayPassCharPrefixContext() {}

func NewMayPassCharPrefixContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayPassCharPrefixContext {
	var p = new(MayPassCharPrefixContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayPassCharPrefix

	return p
}

func (s *MayPassCharPrefixContext) GetParser() antlr.Parser { return s.parser }

func (s *MayPassCharPrefixContext) CharLit() antlr.TerminalNode {
	return s.GetToken(LDEParserCharLit, 0)
}

func (s *MayPassCharPrefixContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayPassCharPrefixContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayPassCharPrefixContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayPassCharPrefix(s)
	}
}

func (s *MayPassCharPrefixContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayPassCharPrefix(s)
	}
}

func (p *LDEParser) MayPassCharPrefix() (localctx IMayPassCharPrefixContext) {
	localctx = NewMayPassCharPrefixContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LDEParserRULE_mayPassCharPrefix)

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
		p.SetState(103)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(104)
		p.Match(LDEParserT__4)
	}
	{
		p.SetState(105)
		p.Match(LDEParserCharLit)
	}

	return localctx
}

// IPassCharsContext is an interface to support dynamic dispatch.
type IPassCharsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassCharsContext differentiates from other interfaces.
	IsPassCharsContext()
}

type PassCharsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassCharsContext() *PassCharsContext {
	var p = new(PassCharsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passChars
	return p
}

func (*PassCharsContext) IsPassCharsContext() {}

func NewPassCharsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassCharsContext {
	var p = new(PassCharsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passChars

	return p
}

func (s *PassCharsContext) GetParser() antlr.Parser { return s.parser }

func (s *PassCharsContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *PassCharsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassCharsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassCharsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassChars(s)
	}
}

func (s *PassCharsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassChars(s)
	}
}

func (p *LDEParser) PassChars() (localctx IPassCharsContext) {
	localctx = NewPassCharsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LDEParserRULE_passChars)

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
		p.SetState(107)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(108)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(109)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(110)
		p.Match(LDEParserT__8)
	}
	{
		p.SetState(111)
		p.Match(LDEParserT__9)
	}

	return localctx
}

// IPassUntilContext is an interface to support dynamic dispatch.
type IPassUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPassUntilContext differentiates from other interfaces.
	IsPassUntilContext()
}

type PassUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPassUntilContext() *PassUntilContext {
	var p = new(PassUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_passUntil
	return p
}

func (*PassUntilContext) IsPassUntilContext() {}

func NewPassUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PassUntilContext {
	var p = new(PassUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_passUntil

	return p
}

func (s *PassUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *PassUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *PassUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PassUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PassUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterPassUntil(s)
	}
}

func (s *PassUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitPassUntil(s)
	}
}

func (p *LDEParser) PassUntil() (localctx IPassUntilContext) {
	localctx = NewPassUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LDEParserRULE_passUntil)

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
		p.SetState(113)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(114)
		p.Target()
	}

	return localctx
}

// IMayPassUntilContext is an interface to support dynamic dispatch.
type IMayPassUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMayPassUntilContext differentiates from other interfaces.
	IsMayPassUntilContext()
}

type MayPassUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMayPassUntilContext() *MayPassUntilContext {
	var p = new(MayPassUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_mayPassUntil
	return p
}

func (*MayPassUntilContext) IsMayPassUntilContext() {}

func NewMayPassUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MayPassUntilContext {
	var p = new(MayPassUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_mayPassUntil

	return p
}

func (s *MayPassUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *MayPassUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *MayPassUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MayPassUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MayPassUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterMayPassUntil(s)
	}
}

func (s *MayPassUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitMayPassUntil(s)
	}
}

func (p *LDEParser) MayPassUntil() (localctx IMayPassUntilContext) {
	localctx = NewMayPassUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LDEParserRULE_mayPassUntil)

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
		p.SetState(116)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(117)
		p.Match(LDEParserT__6)
	}
	{
		p.SetState(118)
		p.Target()
	}

	return localctx
}

// ITakeUntilContext is an interface to support dynamic dispatch.
type ITakeUntilContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilContext differentiates from other interfaces.
	IsTakeUntilContext()
}

type TakeUntilContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilContext() *TakeUntilContext {
	var p = new(TakeUntilContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntil
	return p
}

func (*TakeUntilContext) IsTakeUntilContext() {}

func NewTakeUntilContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilContext {
	var p = new(TakeUntilContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntil

	return p
}

func (s *TakeUntilContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntil(s)
	}
}

func (s *TakeUntilContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntil(s)
	}
}

func (p *LDEParser) TakeUntil() (localctx ITakeUntilContext) {
	localctx = NewTakeUntilContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LDEParserRULE_takeUntil)

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
		p.SetState(120)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(121)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(122)
		p.FieldType()
	}
	{
		p.SetState(123)
		p.Match(LDEParserT__3)
	}
	{
		p.SetState(124)
		p.Target()
	}

	return localctx
}

// ITakeUntilOrRestContext is an interface to support dynamic dispatch.
type ITakeUntilOrRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilOrRestContext differentiates from other interfaces.
	IsTakeUntilOrRestContext()
}

type TakeUntilOrRestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilOrRestContext() *TakeUntilOrRestContext {
	var p = new(TakeUntilOrRestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilOrRest
	return p
}

func (*TakeUntilOrRestContext) IsTakeUntilOrRestContext() {}

func NewTakeUntilOrRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilOrRestContext {
	var p = new(TakeUntilOrRestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilOrRest

	return p
}

func (s *TakeUntilOrRestContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilOrRestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilOrRestContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilOrRestContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TakeUntilOrRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilOrRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilOrRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilOrRest(s)
	}
}

func (s *TakeUntilOrRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilOrRest(s)
	}
}

func (p *LDEParser) TakeUntilOrRest() (localctx ITakeUntilOrRestContext) {
	localctx = NewTakeUntilOrRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LDEParserRULE_takeUntilOrRest)

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
		p.SetState(126)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(127)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(128)
		p.FieldType()
	}
	{
		p.SetState(129)
		p.Match(LDEParserT__3)
	}
	{
		p.SetState(130)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(131)
		p.Target()
	}

	return localctx
}

// ITakeUntilRestContext is an interface to support dynamic dispatch.
type ITakeUntilRestContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTakeUntilRestContext differentiates from other interfaces.
	IsTakeUntilRestContext()
}

type TakeUntilRestContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTakeUntilRestContext() *TakeUntilRestContext {
	var p = new(TakeUntilRestContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_takeUntilRest
	return p
}

func (*TakeUntilRestContext) IsTakeUntilRestContext() {}

func NewTakeUntilRestContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TakeUntilRestContext {
	var p = new(TakeUntilRestContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_takeUntilRest

	return p
}

func (s *TakeUntilRestContext) GetParser() antlr.Parser { return s.parser }

func (s *TakeUntilRestContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *TakeUntilRestContext) FieldType() IFieldTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldTypeContext)
}

func (s *TakeUntilRestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TakeUntilRestContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TakeUntilRestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTakeUntilRest(s)
	}
}

func (s *TakeUntilRestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTakeUntilRest(s)
	}
}

func (p *LDEParser) TakeUntilRest() (localctx ITakeUntilRestContext) {
	localctx = NewTakeUntilRestContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LDEParserRULE_takeUntilRest)

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
		p.SetState(133)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(134)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(135)
		p.FieldType()
	}
	{
		p.SetState(136)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IOptionalNamedAreaContext is an interface to support dynamic dispatch.
type IOptionalNamedAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalNamedAreaContext differentiates from other interfaces.
	IsOptionalNamedAreaContext()
}

type OptionalNamedAreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalNamedAreaContext() *OptionalNamedAreaContext {
	var p = new(OptionalNamedAreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_optionalNamedArea
	return p
}

func (*OptionalNamedAreaContext) IsOptionalNamedAreaContext() {}

func NewOptionalNamedAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalNamedAreaContext {
	var p = new(OptionalNamedAreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_optionalNamedArea

	return p
}

func (s *OptionalNamedAreaContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalNamedAreaContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *OptionalNamedAreaContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *OptionalNamedAreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalNamedAreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalNamedAreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterOptionalNamedArea(s)
	}
}

func (s *OptionalNamedAreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitOptionalNamedArea(s)
	}
}

func (p *LDEParser) OptionalNamedArea() (localctx IOptionalNamedAreaContext) {
	localctx = NewOptionalNamedAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LDEParserRULE_optionalNamedArea)

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
		p.SetState(138)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(139)
		p.Match(LDEParserIdentifier)
	}
	{
		p.SetState(140)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(141)
		p.BaseAction()
	}
	{
		p.SetState(142)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IOptionalAreaContext is an interface to support dynamic dispatch.
type IOptionalAreaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOptionalAreaContext differentiates from other interfaces.
	IsOptionalAreaContext()
}

type OptionalAreaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalAreaContext() *OptionalAreaContext {
	var p = new(OptionalAreaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_optionalArea
	return p
}

func (*OptionalAreaContext) IsOptionalAreaContext() {}

func NewOptionalAreaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalAreaContext {
	var p = new(OptionalAreaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_optionalArea

	return p
}

func (s *OptionalAreaContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalAreaContext) BaseAction() IBaseActionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBaseActionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBaseActionContext)
}

func (s *OptionalAreaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalAreaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalAreaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterOptionalArea(s)
	}
}

func (s *OptionalAreaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitOptionalArea(s)
	}
}

func (p *LDEParser) OptionalArea() (localctx IOptionalAreaContext) {
	localctx = NewOptionalAreaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LDEParserRULE_optionalArea)

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
		p.SetState(144)
		p.Match(LDEParserT__5)
	}
	{
		p.SetState(145)
		p.Match(LDEParserT__2)
	}
	{
		p.SetState(146)
		p.BaseAction()
	}
	{
		p.SetState(147)
		p.Match(LDEParserT__3)
	}

	return localctx
}

// IAtEndContext is an interface to support dynamic dispatch.
type IAtEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAtEndContext differentiates from other interfaces.
	IsAtEndContext()
}

type AtEndContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAtEndContext() *AtEndContext {
	var p = new(AtEndContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_atEnd
	return p
}

func (*AtEndContext) IsAtEndContext() {}

func NewAtEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AtEndContext {
	var p = new(AtEndContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_atEnd

	return p
}

func (s *AtEndContext) GetParser() antlr.Parser { return s.parser }
func (s *AtEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AtEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AtEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterAtEnd(s)
	}
}

func (s *AtEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitAtEnd(s)
	}
}

func (p *LDEParser) AtEnd() (localctx IAtEndContext) {
	localctx = NewAtEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LDEParserRULE_atEnd)

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
		p.SetState(149)
		p.Match(LDEParserT__10)
	}

	return localctx
}

// ITargetContext is an interface to support dynamic dispatch.
type ITargetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetContext differentiates from other interfaces.
	IsTargetContext()
}

type TargetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetContext() *TargetContext {
	var p = new(TargetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_target
	return p
}

func (*TargetContext) IsTargetContext() {}

func NewTargetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetContext {
	var p = new(TargetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_target

	return p
}

func (s *TargetContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetContext) TargetLit() ITargetLitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetLitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetLitContext)
}

func (s *TargetContext) Bound() IBoundContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBoundContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBoundContext)
}

func (s *TargetContext) Limit() ILimitContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitContext)
}

func (s *TargetContext) Exact() IExactContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExactContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExactContext)
}

func (s *TargetContext) Target() ITargetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITargetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITargetContext)
}

func (s *TargetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTarget(s)
	}
}

func (s *TargetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTarget(s)
	}
}

func (p *LDEParser) Target() (localctx ITargetContext) {
	localctx = NewTargetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LDEParserRULE_target)

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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(151)
			p.TargetLit()
		}
		{
			p.SetState(152)
			p.Bound()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.TargetLit()
		}
		{
			p.SetState(155)
			p.Limit()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(157)
			p.TargetLit()
		}
		{
			p.SetState(158)
			p.Exact()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(160)
			p.TargetLit()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(161)
			p.Match(LDEParserT__11)
		}
		{
			p.SetState(162)
			p.Target()
		}

	}

	return localctx
}

// ITargetLitContext is an interface to support dynamic dispatch.
type ITargetLitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTargetLitContext differentiates from other interfaces.
	IsTargetLitContext()
}

type TargetLitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTargetLitContext() *TargetLitContext {
	var p = new(TargetLitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_targetLit
	return p
}

func (*TargetLitContext) IsTargetLitContext() {}

func NewTargetLitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TargetLitContext {
	var p = new(TargetLitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_targetLit

	return p
}

func (s *TargetLitContext) GetParser() antlr.Parser { return s.parser }

func (s *TargetLitContext) CharLit() antlr.TerminalNode {
	return s.GetToken(LDEParserCharLit, 0)
}

func (s *TargetLitContext) StringLit() antlr.TerminalNode {
	return s.GetToken(LDEParserStringLit, 0)
}

func (s *TargetLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TargetLitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TargetLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterTargetLit(s)
	}
}

func (s *TargetLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitTargetLit(s)
	}
}

func (p *LDEParser) TargetLit() (localctx ITargetLitContext) {
	localctx = NewTargetLitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LDEParserRULE_targetLit)
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
	p.SetState(165)
	_la = p.GetTokenStream().LA(1)

	if !(_la == LDEParserStringLit || _la == LDEParserCharLit) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IBoundContext is an interface to support dynamic dispatch.
type IBoundContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBoundContext differentiates from other interfaces.
	IsBoundContext()
}

type BoundContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBoundContext() *BoundContext {
	var p = new(BoundContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_bound
	return p
}

func (*BoundContext) IsBoundContext() {}

func NewBoundContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BoundContext {
	var p = new(BoundContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_bound

	return p
}

func (s *BoundContext) GetParser() antlr.Parser { return s.parser }

func (s *BoundContext) AllIntLit() []antlr.TerminalNode {
	return s.GetTokens(LDEParserIntLit)
}

func (s *BoundContext) IntLit(i int) antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, i)
}

func (s *BoundContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoundContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BoundContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterBound(s)
	}
}

func (s *BoundContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitBound(s)
	}
}

func (p *LDEParser) Bound() (localctx IBoundContext) {
	localctx = NewBoundContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LDEParserRULE_bound)

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
		p.SetState(167)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(168)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(169)
		p.Match(LDEParserT__8)
	}
	{
		p.SetState(170)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(171)
		p.Match(LDEParserT__9)
	}

	return localctx
}

// ILimitContext is an interface to support dynamic dispatch.
type ILimitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitContext differentiates from other interfaces.
	IsLimitContext()
}

type LimitContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitContext() *LimitContext {
	var p = new(LimitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_limit
	return p
}

func (*LimitContext) IsLimitContext() {}

func NewLimitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitContext {
	var p = new(LimitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_limit

	return p
}

func (s *LimitContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *LimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterLimit(s)
	}
}

func (s *LimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitLimit(s)
	}
}

func (p *LDEParser) Limit() (localctx ILimitContext) {
	localctx = NewLimitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LDEParserRULE_limit)

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
		p.SetState(173)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(174)
		p.Match(LDEParserT__8)
	}
	{
		p.SetState(175)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(176)
		p.Match(LDEParserT__9)
	}

	return localctx
}

// IExactContext is an interface to support dynamic dispatch.
type IExactContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExactContext differentiates from other interfaces.
	IsExactContext()
}

type ExactContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExactContext() *ExactContext {
	var p = new(ExactContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_exact
	return p
}

func (*ExactContext) IsExactContext() {}

func NewExactContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExactContext {
	var p = new(ExactContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_exact

	return p
}

func (s *ExactContext) GetParser() antlr.Parser { return s.parser }

func (s *ExactContext) IntLit() antlr.TerminalNode {
	return s.GetToken(LDEParserIntLit, 0)
}

func (s *ExactContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExactContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExactContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterExact(s)
	}
}

func (s *ExactContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitExact(s)
	}
}

func (p *LDEParser) Exact() (localctx IExactContext) {
	localctx = NewExactContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LDEParserRULE_exact)

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
		p.SetState(178)
		p.Match(LDEParserT__7)
	}
	{
		p.SetState(179)
		p.Match(LDEParserIntLit)
	}
	{
		p.SetState(180)
		p.Match(LDEParserT__9)
	}

	return localctx
}

// IFieldTypeContext is an interface to support dynamic dispatch.
type IFieldTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldTypeContext differentiates from other interfaces.
	IsFieldTypeContext()
}

type FieldTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldTypeContext() *FieldTypeContext {
	var p = new(FieldTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LDEParserRULE_fieldType
	return p
}

func (*FieldTypeContext) IsFieldTypeContext() {}

func NewFieldTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldTypeContext {
	var p = new(FieldTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LDEParserRULE_fieldType

	return p
}

func (s *FieldTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldTypeContext) Identifier() antlr.TerminalNode {
	return s.GetToken(LDEParserIdentifier, 0)
}

func (s *FieldTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.EnterFieldType(s)
	}
}

func (s *FieldTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LDEListener); ok {
		listenerT.ExitFieldType(s)
	}
}

func (p *LDEParser) FieldType() (localctx IFieldTypeContext) {
	localctx = NewFieldTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LDEParserRULE_fieldType)

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
		p.SetState(182)
		p.Match(LDEParserIdentifier)
	}

	return localctx
}

func (p *LDEParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 0:
		var t *RulesContext = nil
		if localctx != nil {
			t = localctx.(*RulesContext)
		}
		return p.Rules_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LDEParser) Rules_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
