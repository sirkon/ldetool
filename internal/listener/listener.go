// Generated from LDE.g4 by ANTLR 4.7.

package listener // LDE

import (
	"strconv"

	"fmt"

	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/internal/ast"
	"github.com/sirkon/ldetool/internal/parser"
)

type appender interface {
	Append(i ast.Action)
}

var reservedWords = map[string]struct{}{
	"Valid":   {},
	"Extract": {},
}

func checkReserved(token antlr.Token) {
	if _, ok := reservedWords[token.GetText()]; !ok {
		return
	}
	panic(
		fmt.Sprintf("%d:%d: \033[1m%s\033[0m is reserved identifier",
			token.GetLine(),
			token.GetColumn()+1,
			token.GetText(),
		),
	)
}

// Listener is a complete listener for a parse tree produced by LDEParser.
type Listener struct {
	rules   []*ast.Rule
	actions []appender
	target  *ast.Target

	prefixJump int

	optional       bool
	lookup         bool
	stateIsPrefix  bool
	expectEnd      bool
	mustNotBeExact bool
}

func New() *Listener {
	return &Listener{}
}

func (l *Listener) Rules() []*ast.Rule {
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
			"%d:%d: previous action consumed the rest of the string, the remaining ops will do nothing",
			token.GetLine(),
			token.GetColumn()+1,
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
func (l *Listener) ExitAtomicRule(ctx *parser.AtomicRuleContext) {
	l.expectEnd = false
}

// EnterBaseAction is called when production baseAction is entered.
func (l *Listener) EnterBaseAction(ctx *parser.BaseActionContext) {
	if ctx.Stress() != nil {
		res := ast.ErrorOnMismatch{}
		l.seq().Append(res)
	}
}

// ExitBaseAction is called when production baseAction is exited.
func (l *Listener) ExitBaseAction(ctx *parser.BaseActionContext) {}

// EnterAtomicAction is called when production atomicAction is entered.
func (l *Listener) EnterAtomicAction(ctx *parser.AtomicActionContext) {}

// ExitAtomicAction is called when production atomicAction is exited.
func (l *Listener) ExitAtomicAction(ctx *parser.AtomicActionContext) {}

// EnterPassTargetPrefix is called when production passTargetPrefix is entered.
func (l *Listener) EnterPassTargetPrefix(ctx *parser.PassTargetPrefixContext) {
	l.stateIsPrefix = true
	l.prefixJump = 0
	if ctx.IntLit() != nil {
		l.prefixJump, _ = strconv.Atoi(ctx.IntLit().GetText())
	}
}

// ExitPassTargetPrefix is called when production passTargetPrefix is exited.
func (l *Listener) ExitPassTargetPrefix(ctx *parser.PassTargetPrefixContext) {
	l.stateIsPrefix = false
	l.optional = false
}

// EnterMayBePassTargetPrefix is called when production mayBePassTargetPrefix is entered.
func (l *Listener) EnterMayBePassTargetPrefix(ctx *parser.MayBePassTargetPrefixContext) {
	l.stateIsPrefix = true
	l.optional = true
	l.prefixJump = 0
	if ctx.IntLit() != nil {
		l.prefixJump, _ = strconv.Atoi(ctx.IntLit().GetText())
	}
}

// ExitMayBePassTargetPrefix is called when production mayBePassTargetPrefix is exited.
func (l *Listener) ExitMayBePassTargetPrefix(ctx *parser.MayBePassTargetPrefixContext) {
	l.stateIsPrefix = false
	l.optional = false
}

// EnterPassChars is called when production passChars is entered.
func (l *Listener) EnterPassChars(ctx *parser.PassCharsContext) {
	var err error
	a, err := ast.PassFirst(ctx.IntLit().GetSymbol())
	if err != nil {
		panic(fmt.Sprintf("%d:%d: positive integer expected, got %s", ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()+1, ctx.IntLit().GetSymbol()))
	}
	l.seq().Append(a)
}

// ExitPassChars is called when production passChars is exited.
func (l *Listener) ExitPassChars(ctx *parser.PassCharsContext) {}

// EnterPassUntil is called when production passUntil is entered.
func (l *Listener) EnterPassUntil(ctx *parser.PassUntilContext) {
	a := ast.PassUntilTarget()
	l.seq().Append(a)
	l.target = a.Limit
	l.lookup = true
}

// ExitPassUntil is called when production passUntil is exited.
func (l *Listener) ExitPassUntil(ctx *parser.PassUntilContext) {
	l.lookup = false
	if l.mustNotBeExact {
		panic(fmt.Sprintf(
			"%d:%d: use prefix operator (\033[1m^\033[0m) instead of \033[1m_\033[0m",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn()+1,
		))
	}
}

// EnterMayPassUntil is called when production mayPassUntil is entered.
func (l *Listener) EnterMayPassUntil(ctx *parser.MayPassUntilContext) {
	a := ast.PassUntilTargetOrIgnore()
	l.seq().Append(a)
	l.target = a.Limit
	l.lookup = true
}

// ExitMayPassUntil is called when production mayPassUntil is exited.
func (l *Listener) ExitMayPassUntil(ctx *parser.MayPassUntilContext) {
	l.lookup = false
	if l.mustNotBeExact {
		panic(fmt.Sprintf(
			"%d:%d: use prefix operator (\033[1m^\033[0m) instead of \033[1m_\033[0m",
			ctx.GetStart().GetLine(),
			ctx.GetStart().GetColumn()+1,
		))
	}
}

// EnterTakeUntil is called when production takeUntil is entered.
func (l *Listener) EnterTakeUntil(ctx *parser.TakeUntilContext) {
	checkReserved(ctx.Identifier().GetSymbol())
	a := ast.TakeUntilTarget(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
	l.seq().Append(a)
	l.target = a.Limit
}

// ExitTakeUntil is called when production takeUntil is exited.
func (l *Listener) ExitTakeUntil(ctx *parser.TakeUntilContext) {}

// EnterTakeUntilOrRest is called when production takeUntilOrRest is entered.
func (l *Listener) EnterTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {
	checkReserved(ctx.Identifier().GetSymbol())
	a := ast.TakeUntilTargetOrRest(
		ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart(),
	)
	l.seq().Append(a)
	l.target = a.Limit
}

// ExitTakeUntilOrRest is called when production takeUntilOrRest is exited.
func (l *Listener) ExitTakeUntilOrRest(ctx *parser.TakeUntilOrRestContext) {}

// EnterTakeUntilRest is called when production takeUntilRest is entered.
func (l *Listener) EnterTakeUntilRest(ctx *parser.TakeUntilRestContext) {
	checkReserved(ctx.Identifier().GetSymbol())
	a := ast.TakeTheRest(ctx.Identifier().GetSymbol(), ctx.FieldType().GetStart())
	l.seq().Append(a)
}

// ExitTakeUntilRest is called when production takeUntilRest is exited.
func (l *Listener) ExitTakeUntilRest(ctx *parser.TakeUntilRestContext) {
	l.expectEnd = true
}

// EnterOptionalNamedArea is called when production optionalNamedArea is entered.
func (l *Listener) EnterOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	checkReserved(ctx.Identifier().GetSymbol())

	a := ast.Option(ctx.Identifier().GetSymbol())
	l.seq().Append(a)
	l.actions = append(l.actions, a)
}

// ExitOptionalNamedArea is called when production optionalNamedArea is exited.
func (l *Listener) ExitOptionalNamedArea(ctx *parser.OptionalNamedAreaContext) {
	l.actions = l.actions[:len(l.actions)-1]
}

// EnterOptionalArea is called when production optionalArea is entered.
func (l *Listener) EnterOptionalArea(ctx *parser.OptionalAreaContext) {
	a := ast.Anonymous(ctx.GetStart())
	l.seq().Append(a)
	l.actions = append(l.actions, a)
}

// ExitOptionalArea is called when production optionalArea is exited.
func (l *Listener) ExitOptionalArea(ctx *parser.OptionalAreaContext) {
	l.actions = l.actions[:len(l.actions)-1]
}

// EnterAtEnd is called when production atEnd is entered.
func (l *Listener) EnterAtEnd(ctx *parser.AtEndContext) {
	a := ast.AtEnd{}
	l.seq().Append(a)
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
	var a ast.Action
	if l.stateIsPrefix {
		if l.prefixJump == 0 {
			if ctx.StringLit() != nil {
				if l.optional {
					a = ast.MayBeStartsWithString(ctx.StringLit().GetSymbol())
				} else {
					a = ast.StartsWithString(ctx.StringLit().GetSymbol())
				}
			} else if ctx.CharLit() != nil {
				if l.optional {
					a = ast.MayBeStartsWithChar(ctx.CharLit().GetSymbol())
				} else {
					a = ast.StartsWithChar(ctx.CharLit().GetSymbol())
				}
			}
			if a != nil {
				l.seq().Append(a)
			}
			return
		} else {
			if l.optional {
				ll := ast.PassUntilTargetOrIgnore()
				a = ll
				l.target = ll.Limit
			} else {
				ll := ast.PassUntilTarget()
				a = ll
				l.target = ll.Limit
			}
			l.target.SetBound(l.prefixJump, l.prefixJump)
		}
	}
	if a != nil {
		l.seq().Append(a)
	}
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
	//if l.target.Close {
	//	panic(fmt.Sprintf(
	//		"%d:%d: short lookup does not make a sense on bounded areas",
	//		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()))
	//}
	lower, _ := strconv.Atoi(ctx.IntLit(0).GetText())
	upper, _ := strconv.Atoi(ctx.IntLit(1).GetText())
	if lower == 0 {
		token := ctx.IntLit(0).GetSymbol()
		panic(fmt.Sprintf("%d:%d offset value must be greater than 0", token.GetLine(), token.GetColumn()+1))
	}
	if upper < lower {
		token := ctx.IntLit(1).GetSymbol()
		panic(fmt.Sprintf("%d:%d: upper bound must be greater than lower",
			token.GetLine(), token.GetColumn()+1))
	}
	l.target.SetBound(lower, upper)
}

// ExitBound is called when production bound is exited.
func (l *Listener) ExitBound(ctx *parser.BoundContext) {}

// EnterLimit is called when production limit is entered.
func (l *Listener) EnterLimit(ctx *parser.LimitContext) {
	//if l.target.Close {
	//	panic(fmt.Sprintf(
	//		"%d:%d: short lookup does not make a sense on limited areas",
	//		ctx.GetStart().GetLine(), ctx.GetStart().GetColumn()))
	//}
	upper, _ := strconv.Atoi(ctx.IntLit().GetText())
	if upper == 0 {
		token := ctx.IntLit().GetSymbol()
		panic(fmt.Sprintf("%d:%d upper bound must be greater than 0", token.GetLine(), token.GetColumn()+1))
	}
	l.target.SetLimit(upper)
}

// ExitLimit is called when production limit is exited.
func (l *Listener) ExitLimit(ctx *parser.LimitContext) {}

// EnterJump is called when production jump is entered.
func (l *Listener) EnterJump(ctx *parser.JumpContext) {
	lower, _ := strconv.Atoi(ctx.IntLit().GetText())
	if lower == 0 {
		token := ctx.IntLit().GetSymbol()
		panic(fmt.Sprintf("%d:%d offset value must be greater than 0", token.GetLine(), token.GetColumn()+1))
	}
	l.target.SetJump(lower)
}

// ExitJump is called when production jump is exited.
func (l *Listener) ExitJump(ctx *parser.JumpContext) {}

// EnterExact is called when production exact is entered.
func (l *Listener) EnterExact(ctx *parser.ExactContext) {
	if l.lookup {
		l.mustNotBeExact = true
		return
	}
	index, _ := strconv.Atoi(ctx.IntLit().GetText())
	l.target.SetBound(index, index)
}

// ExitExact is called when production exact is exited.
func (l *Listener) ExitExact(ctx *parser.ExactContext) {}

var acceptablesTypeList string
var acceptableTypesMap = func() map[string]struct{} {
	acceptableTypes := []string{
		"int", "uint",
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
		"%d:%d: unsupported type `\033[1m%s\033[0m`, must be one of %s",
		ctx.Identifier().GetSymbol().GetLine(),
		ctx.Identifier().GetSymbol().GetColumn()+1,
		typeName,
		acceptablesTypeList,
	))
}

// ExitFieldType is called when production fieldType is exited.
func (l *Listener) ExitFieldType(ctx *parser.FieldTypeContext) {}
