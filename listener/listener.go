// Generated from LDE.g4 by ANTLR 4.7.

package listener // LDE

import (
	"strconv"

	"fmt"

	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/ast"
	"github.com/sirkon/ldetool/parser"
)

type appender interface {
	Append(i *ast.ActionItem)
}

// Listener is a complete listener for a parse tree produced by LDEParser.
type Listener struct {
	rules     []*ast.RuleItem
	ai        *ast.ActionItem
	actions   []appender
	target    *ast.Target
	expectEnd bool
}

func New() *Listener {
	return &Listener{}
}

func (l *Listener) Rules() []*ast.RuleItem {
	return l.rules
}

// VisitTerminal is called when a terminal node is visited.
func (l *Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (l *Listener) VisitErrorNode(node antlr.ErrorNode) {
}

// EnterEveryRule is called when any rule is entered.
func (l *Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	if l.expectEnd {
		token := ctx.GetStart()
		panic(fmt.Sprintf(
			"%d:%d: previous action consumed the rest of the string, the rest will do nothing",
			token.GetLine(),
			token.GetColumn(),
		))
	}
}

// ExitEveryRule is called when any rule is exited.
func (l *Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRules is called when production rules is entered.
func (l *Listener) EnterRules(ctx *parser.RulesContext) {
}

func (l *Listener) seq() appender {
	res := l.actions[len(l.actions)-1]
	return res
}

// ExitRules is called when production rules is exited.
func (l *Listener) ExitRules(ctx *parser.RulesContext) {}

// EnterAtomicRule is called when production atomicRule is entered.
func (l *Listener) EnterAtomicRule(ctx *parser.AtomicRuleContext) {
	rule := ast.NewRule(ctx.Identifier().GetSymbol())
	l.rules = append(l.rules, rule)
	l.actions = append(l.actions, rule)
}

// ExitAtomicRule is called when production atomicRule is exited.
func (l *Listener) ExitAtomicRule(ctx *parser.AtomicRuleContext) {}

// EnterBaseAction is called when production baseAction is entered.
func (l *Listener) EnterBaseAction(ctx *parser.BaseActionContext) {
	if ctx.Stress() != nil {
		res := &ast.ActionItem{
			ErrorOnMismatch: true,
		}
		l.seq().Append(res)
		l.ai = nil
	}
}

// ExitBaseAction is called when production baseAction is exited.
func (l *Listener) ExitBaseAction(ctx *parser.BaseActionContext) {}

// EnterAtomicAction is called when production atomicAction is entered.
func (l *Listener) EnterAtomicAction(ctx *parser.AtomicActionContext) {
	i := &ast.ActionItem{}
	l.seq().Append(i)
	l.ai = i
}

// ExitAtomicAction is called when production atomicAction is exited.
func (l *Listener) ExitAtomicAction(ctx *parser.AtomicActionContext) {}

// EnterPassStringPrefix is called when production passStringPrefix is entered.
func (l *Listener) EnterPassStringPrefix(ctx *parser.PassStringPrefixContext) {
	//l.ai.StartWithString, _ = ast.StartsWithString(token.)
	l.ai.StartWithString, _ = ast.StartsWithString(ctx.StringLit().GetSymbol())
}

// ExitPassStringPrefix is called when production passStringPrefix is exited.
func (l *Listener) ExitPassStringPrefix(ctx *parser.PassStringPrefixContext) {}

// EnterPassCharPrefix is called when production passCharPrefix is entered.
func (l *Listener) EnterPassCharPrefix(ctx *parser.PassCharPrefixContext) {
	l.ai.StartWithChar, _ = ast.StartsWithChar(ctx.CharLit().GetSymbol())
}

// ExitPassCharPrefix is called when production passCharPrefix is exited.
func (l *Listener) ExitPassCharPrefix(ctx *parser.PassCharPrefixContext) {}

// EnterMayPassStringPrefix is called when production mayPassStringPrefix is entered.
func (l *Listener) EnterMayPassStringPrefix(ctx *parser.MayPassStringPrefixContext) {
	l.ai.MayBeStartWithString, _ = ast.MayBeStartsWithString(ctx.StringLit().GetSymbol())
}

// ExitMayPassStringPrefix is called when production mayPassStringPrefix is exited.
func (l *Listener) ExitMayPassStringPrefix(ctx *parser.MayPassStringPrefixContext) {}

// EnterMayPassCharPrefix is called when production mayPassCharPrefix is entered.
func (l *Listener) EnterMayPassCharPrefix(ctx *parser.MayPassCharPrefixContext) {
	l.ai.MayBeStartWithChar, _ = ast.MayBeStartsWithChar(ctx.CharLit().GetSymbol())
}

// ExitMayPassCharPrefix is called when production mayPassCharPrefix is exited.
func (l *Listener) ExitMayPassCharPrefix(ctx *parser.MayPassCharPrefixContext) {}

// EnterPassChars is called when production passChars is entered.
func (l *Listener) EnterPassChars(ctx *parser.PassCharsContext) {
	l.ai.PassFirst, _ = ast.PassFirst(ctx.IntLit().GetSymbol())
}

// ExitPassChars is called when production passChars is exited.
func (l *Listener) ExitPassChars(ctx *parser.PassCharsContext) {}

// EnterPassUntil is called when production passUntil is entered.
func (l *Listener) EnterPassUntil(ctx *parser.PassUntilContext) {
	l.ai.Pass, _ = ast.PassUntilTarget()
	l.target = l.ai.Pass.Limit
}

// ExitPassUntil is called when production passUntil is exited.
func (l *Listener) ExitPassUntil(ctx *parser.PassUntilContext) {}

// EnterMayPassUntil is called when production mayPassUntil is entered.
func (l *Listener) EnterMayPassUntil(ctx *parser.MayPassUntilContext) {
	l.ai.PassOrIgnore, _ = ast.PassUntilTargetOrIgnore()
	l.target = l.ai.PassOrIgnore.Limit
}

// ExitMayPassUntil is called when production mayPassUntil is exited.
func (l *Listener) ExitMayPassUntil(ctx *parser.MayPassUntilContext) {}

// EnterTakeUntil is called when production takeUntil is entered.
func (l *Listener) EnterTakeUntil(ctx *parser.TakeUntilContext) {
	l.ai.Take, _ = ast.TakeUntilTarget(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
	l.target = l.ai.Take.Limit
}

// ExitTakeUntil is called when production takeUntil is exited.
func (l *Listener) ExitTakeUntil(ctx *parser.TakeUntilContext) {}

// EnterTakeUntilOrRest is called when production takeUntilOrRest is entered.
func (l *Listener) EnterTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {
	l.ai.TakeUntilOrRest, _ = ast.TakeUntilTargetOrRest(
		ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart(),
	)
	l.target = l.ai.TakeUntilOrRest.Limit
}

// ExitTakeUntilOrRest is called when production takeUntilOrRest is exited.
func (l *Listener) ExitTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {}

// EnterTakeUntilRest is called when production takeUntilRest is entered.
func (l *Listener) EnterTakeUntilRest(ctx *parser.TakeUntilRestContext) {
	l.ai.TakeRest, _ = ast.TakeTheRest(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
}

// ExitTakeUntilRest is called when production takeUntilRest is exited.
func (l *Listener) ExitTakeUntilRest(ctx *parser.TakeUntilRestContext) {
	l.expectEnd = true
}

// EnterOptionalNamedArea is called when production optionalNamedArea is entered.
func (l *Listener) EnterOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	l.ai.Option, _ = ast.Option(ctx.Identifier().GetSymbol())

	l.actions = append(l.actions, l.ai.Option)
}

// ExitOptionalNamedArea is called when production optionalNamedArea is exited.
func (l *Listener) ExitOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	l.actions = l.actions[:len(l.actions)-1]
}

// EnterOptionalArea is called when production optionalArea is entered.
func (l *Listener) EnterOptionalArea(ctx *parser.OptionalAreaContext) {
	l.ai.Anonymous, _ = ast.Anonymous(ctx.GetStart())
	l.actions = append(l.actions, l.ai.Anonymous)
}

// ExitOptionalArea is called when production optionalArea is exited.
func (l *Listener) ExitOptionalArea(ctx *parser.OptionalAreaContext) {
	l.actions = l.actions[:len(l.actions)-1]
}

// EnterAtEnd is called when production atEnd is entered.
func (l *Listener) EnterAtEnd(ctx *parser.AtEndContext) {
	l.ai.End, _ = ast.AtTheEnd()
}

// ExitAtEnd is called when production atEnd is exited.
func (l *Listener) ExitAtEnd(ctx *parser.AtEndContext) {}

// EnterTarget is called when production target is entered.
func (l *Listener) EnterTarget(ctx *parser.TargetContext) {
	if ctx.Target() != nil {
		l.target.SetClose()
	}
}

// ExitTarget is called when production target is exited.
func (l *Listener) ExitTarget(ctx *parser.TargetContext) {}

// EnterTargetLit is called when production targetLit is entered.
func (l *Listener) EnterTargetLit(ctx *parser.TargetLitContext) {
	if ctx.StringLit() != nil {
		l.target.SetString(ctx.StringLit().GetText())
	} else if ctx.CharLit() != nil {
		l.target.SetChar(ctx.CharLit().GetText())
	} else {
		panic("Integerity error")
	}
}

// ExitTargetLit is called when production targetLit is exited.
func (l *Listener) ExitTargetLit(ctx *parser.TargetLitContext) {}

// EnterBound is called when production bound is entered.
func (l *Listener) EnterBound(ctx *parser.BoundContext) {
	lower, _ := strconv.Atoi(ctx.IntLit(0).GetText())
	upper, _ := strconv.Atoi(ctx.IntLit(1).GetText())
	l.target.SetBound(lower, upper)
}

// ExitBound is called when production bound is exited.
func (l *Listener) ExitBound(ctx *parser.BoundContext) {}

// EnterLimit is called when production limit is entered.
func (l *Listener) EnterLimit(ctx *parser.LimitContext) {
	lower, _ := strconv.Atoi(ctx.IntLit().GetText())
	l.target.SetLimit(lower)
}

// ExitLimit is called when production limit is exited.
func (l *Listener) ExitLimit(ctx *parser.LimitContext) {}

// EnterExact is called when production exact is entered.
func (l *Listener) EnterExact(ctx *parser.ExactContext) {
	index, _ := strconv.Atoi(ctx.IntLit().GetText())
	l.target.SetBound(index, index)
}

// ExitExact is called when production exact is exited.
func (l *Listener) ExitExact(ctx *parser.ExactContext) {}

var acceptablesTypeList string
var acceptableTypesMap = func() map[string]struct{} {
	acceptableTypes := []string{
		"int8", "int16", "int32", "int64",
		"uint8", "uint16", "uint32", "uint64",
		"float32", "float64",
		"string",
	}
	res := map[string]struct{}{}
	for i, typeName := range acceptableTypes {
		res[typeName] = struct{}{}
		acceptableTypes[i] = fmt.Sprintf("\033[1m%s\033[0m", typeName)
	}
	acceptablesTypeList = strings.Join(acceptableTypes, ", ")
	return res
}()

// EnterFieldType is called when production fieldType is entered.
func (l *Listener) EnterFieldType(ctx *parser.FieldTypeContext) {
	typeName := ctx.Identifier().GetText()
	if _, ok := acceptableTypesMap[typeName]; ok {
		return
	}
	panic(fmt.Sprintf(
		"%d:%d: unsupported type `\033[1m%l\033[0m`, must be one of %l",
		ctx.Identifier().GetSymbol().GetLine(),
		ctx.Identifier().GetSymbol().GetColumn()+len(typeName)/2,
		typeName,
		acceptablesTypeList,
	))
}

// ExitFieldType is called when production fieldType is exited.
func (l *Listener) ExitFieldType(ctx *parser.FieldTypeContext) {}
