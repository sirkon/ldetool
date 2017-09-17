// Generated from LDE.g4 by ANTLR 4.7.

package listener // LDE

import (
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/ast"
	"github.com/sirkon/ldetool/parser"
)

type appender interface {
	Append(i *ast.ActionItem)
}

// Listener is a complete listener for a parse tree produced by LDEParser.
type Listener struct {
	rules   []*ast.RuleItem
	ai      *ast.ActionItem
	actions []appender
	target  *ast.Target
}

func New() *Listener {
	return &Listener{}
}

func (l *Listener) Rules() []*ast.RuleItem {
	return l.rules
}

// VisitTerminal is called when a terminal node is visited.
func (s *Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRules is called when production rules is entered.
func (s *Listener) EnterRules(ctx *parser.RulesContext) {
}

func (s *Listener) seq() appender {
	res := s.actions[len(s.actions)-1]
	return res
}

// ExitRules is called when production rules is exited.
func (s *Listener) ExitRules(ctx *parser.RulesContext) {}

// EnterAtomicRule is called when production atomicRule is entered.
func (s *Listener) EnterAtomicRule(ctx *parser.AtomicRuleContext) {
	rule := ast.NewRule(ctx.Identifier().GetSymbol())
	s.rules = append(s.rules, rule)
	s.actions = append(s.actions, rule)
}

// ExitAtomicRule is called when production atomicRule is exited.
func (s *Listener) ExitAtomicRule(ctx *parser.AtomicRuleContext) {}

// EnterBaseAction is called when production baseAction is entered.
func (s *Listener) EnterBaseAction(ctx *parser.BaseActionContext) {
	if ctx.Stress() != nil {
		res := &ast.ActionItem{
			ErrorOnMismatch: true,
		}
		s.seq().Append(res)
		s.ai = nil
	}
}

// ExitBaseAction is called when production baseAction is exited.
func (s *Listener) ExitBaseAction(ctx *parser.BaseActionContext) {}

// EnterAtomicAction is called when production atomicAction is entered.
func (s *Listener) EnterAtomicAction(ctx *parser.AtomicActionContext) {
	i := &ast.ActionItem{}
	s.seq().Append(i)
	s.ai = i
}

// ExitAtomicAction is called when production atomicAction is exited.
func (s *Listener) ExitAtomicAction(ctx *parser.AtomicActionContext) {}

// EnterPassStringPrefix is called when production passStringPrefix is entered.
func (s *Listener) EnterPassStringPrefix(ctx *parser.PassStringPrefixContext) {
	//s.ai.StartWithString, _ = ast.StartsWithString(token.)
	s.ai.StartWithString, _ = ast.StartsWithString(ctx.StringLit().GetSymbol())
}

// ExitPassStringPrefix is called when production passStringPrefix is exited.
func (s *Listener) ExitPassStringPrefix(ctx *parser.PassStringPrefixContext) {}

// EnterPassCharPrefix is called when production passCharPrefix is entered.
func (s *Listener) EnterPassCharPrefix(ctx *parser.PassCharPrefixContext) {
	s.ai.StartWithChar, _ = ast.StartsWithChar(ctx.CharLit().GetSymbol())
}

// ExitPassCharPrefix is called when production passCharPrefix is exited.
func (s *Listener) ExitPassCharPrefix(ctx *parser.PassCharPrefixContext) {}

// EnterMayPassStringPrefix is called when production mayPassStringPrefix is entered.
func (s *Listener) EnterMayPassStringPrefix(ctx *parser.MayPassStringPrefixContext) {
	s.ai.MayBeStartWithString, _ = ast.MayBeStartsWithString(ctx.StringLit().GetSymbol())
}

// ExitMayPassStringPrefix is called when production mayPassStringPrefix is exited.
func (s *Listener) ExitMayPassStringPrefix(ctx *parser.MayPassStringPrefixContext) {}

// EnterMayPassCharPrefix is called when production mayPassCharPrefix is entered.
func (s *Listener) EnterMayPassCharPrefix(ctx *parser.MayPassCharPrefixContext) {
	s.ai.MayBeStartWithChar, _ = ast.MayBeStartsWithChar(ctx.CharLit().GetSymbol())
}

// ExitMayPassCharPrefix is called when production mayPassCharPrefix is exited.
func (s *Listener) ExitMayPassCharPrefix(ctx *parser.MayPassCharPrefixContext) {}

// EnterPassChars is called when production passChars is entered.
func (s *Listener) EnterPassChars(ctx *parser.PassCharsContext) {
	s.ai.PassFirst, _ = ast.PassFirst(ctx.IntLit().GetSymbol())
}

// ExitPassChars is called when production passChars is exited.
func (s *Listener) ExitPassChars(ctx *parser.PassCharsContext) {}

// EnterPassUntil is called when production passUntil is entered.
func (s *Listener) EnterPassUntil(ctx *parser.PassUntilContext) {
	s.ai.Pass, _ = ast.PassUntilTarget()
	s.target = s.ai.Pass.Limit
}

// ExitPassUntil is called when production passUntil is exited.
func (s *Listener) ExitPassUntil(ctx *parser.PassUntilContext) {}

// EnterMayPassUntil is called when production mayPassUntil is entered.
func (s *Listener) EnterMayPassUntil(ctx *parser.MayPassUntilContext) {
	s.ai.PassOrIgnore, _ = ast.PassUntilTargetOrIgnore()
	s.target = s.ai.PassOrIgnore.Limit
}

// ExitMayPassUntil is called when production mayPassUntil is exited.
func (s *Listener) ExitMayPassUntil(ctx *parser.MayPassUntilContext) {}

// EnterTakeUntil is called when production takeUntil is entered.
func (s *Listener) EnterTakeUntil(ctx *parser.TakeUntilContext) {
	s.ai.Take, _ = ast.TakeUntilTarget(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
	s.target = s.ai.Take.Limit
}

// ExitTakeUntil is called when production takeUntil is exited.
func (s *Listener) ExitTakeUntil(ctx *parser.TakeUntilContext) {}

// EnterTakeUntilOrRest is called when production takeUntilOrRest is entered.
func (s *Listener) EnterTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {
	s.ai.TakeUntilOrRest, _ = ast.TakeUntilTargetOrRest(
		ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart(),
	)
	s.target = s.ai.TakeUntilOrRest.Limit
}

// ExitTakeUntilOrRest is called when production takeUntilOrRest is exited.
func (s *Listener) ExitTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {}

// EnterTakeUntilRest is called when production takeUntilRest is entered.
func (s *Listener) EnterTakeUntilRest(ctx *parser.TakeUntilRestContext) {
	s.ai.TakeRest, _ = ast.TakeTheRest(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
}

// ExitTakeUntilRest is called when production takeUntilRest is exited.
func (s *Listener) ExitTakeUntilRest(ctx *parser.TakeUntilRestContext) {}

// EnterOptionalNamedArea is called when production optionalNamedArea is entered.
func (s *Listener) EnterOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	s.ai.Option, _ = ast.Option(ctx.Identifier().GetSymbol())

	s.actions = append(s.actions, s.ai.Option)
}

// ExitOptionalNamedArea is called when production optionalNamedArea is exited.
func (s *Listener) ExitOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	s.actions = s.actions[:len(s.actions)-1]
}

// EnterOptionalArea is called when production optionalArea is entered.
func (s *Listener) EnterOptionalArea(ctx *parser.OptionalAreaContext) {
	panic("Not supported")
}

// ExitOptionalArea is called when production optionalArea is exited.
func (s *Listener) ExitOptionalArea(ctx *parser.OptionalAreaContext) {}

// EnterAtEnd is called when production atEnd is entered.
func (s *Listener) EnterAtEnd(ctx *parser.AtEndContext) {
	s.ai.End, _ = ast.AtTheEnd()
}

// ExitAtEnd is called when production atEnd is exited.
func (s *Listener) ExitAtEnd(ctx *parser.AtEndContext) {}

// EnterTarget is called when production target is entered.
func (s *Listener) EnterTarget(ctx *parser.TargetContext) {
	if ctx.Target() != nil {
		s.target.SetClose()
	}
}

// ExitTarget is called when production target is exited.
func (s *Listener) ExitTarget(ctx *parser.TargetContext) {}

// EnterTargetLit is called when production targetLit is entered.
func (s *Listener) EnterTargetLit(ctx *parser.TargetLitContext) {
	if ctx.StringLit() != nil {
		s.target.SetString(ctx.StringLit().GetText())
	} else if ctx.CharLit() != nil {
		s.target.SetChar(ctx.CharLit().GetText())
	} else {
		panic("Integerity error")
	}
}

// ExitTargetLit is called when production targetLit is exited.
func (s *Listener) ExitTargetLit(ctx *parser.TargetLitContext) {}

// EnterBound is called when production bound is entered.
func (s *Listener) EnterBound(ctx *parser.BoundContext) {
	lower, _ := strconv.Atoi(ctx.IntLit(0).GetText())
	upper, _ := strconv.Atoi(ctx.IntLit(1).GetText())
	s.target.SetBound(lower, upper)
}

// ExitBound is called when production bound is exited.
func (s *Listener) ExitBound(ctx *parser.BoundContext) {}

// EnterLimit is called when production limit is entered.
func (s *Listener) EnterLimit(ctx *parser.LimitContext) {
	lower, _ := strconv.Atoi(ctx.IntLit().GetText())
	s.target.SetLimit(lower)
}

// ExitLimit is called when production limit is exited.
func (s *Listener) ExitLimit(ctx *parser.LimitContext) {}

// EnterExact is called when production exact is entered.
func (s *Listener) EnterExact(ctx *parser.ExactContext) {
	index, _ := strconv.Atoi(ctx.IntLit().GetText())
	s.target.SetBound(index, index)
}

// ExitExact is called when production exact is exited.
func (s *Listener) ExitExact(ctx *parser.ExactContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *Listener) EnterFieldType(ctx *parser.FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *Listener) ExitFieldType(ctx *parser.FieldTypeContext) {}
