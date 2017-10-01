// Generated from LDE.g4 by ANTLR 4.7.

package parser // LDE

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseLDEListener is a complete listener for a parse tree produced by LDEParser.
type BaseLDEListener struct{}

var _ LDEListener = &BaseLDEListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseLDEListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseLDEListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseLDEListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseLDEListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRules is called when production rules is entered.
func (s *BaseLDEListener) EnterRules(ctx *RulesContext) {}

// ExitRules is called when production rules is exited.
func (s *BaseLDEListener) ExitRules(ctx *RulesContext) {}

// EnterAtomicRule is called when production atomicRule is entered.
func (s *BaseLDEListener) EnterAtomicRule(ctx *AtomicRuleContext) {}

// ExitAtomicRule is called when production atomicRule is exited.
func (s *BaseLDEListener) ExitAtomicRule(ctx *AtomicRuleContext) {}

// EnterBaseAction is called when production baseAction is entered.
func (s *BaseLDEListener) EnterBaseAction(ctx *BaseActionContext) {}

// ExitBaseAction is called when production baseAction is exited.
func (s *BaseLDEListener) ExitBaseAction(ctx *BaseActionContext) {}

// EnterAtomicAction is called when production atomicAction is entered.
func (s *BaseLDEListener) EnterAtomicAction(ctx *AtomicActionContext) {}

// ExitAtomicAction is called when production atomicAction is exited.
func (s *BaseLDEListener) ExitAtomicAction(ctx *AtomicActionContext) {}

// EnterPassStringPrefix is called when production passStringPrefix is entered.
func (s *BaseLDEListener) EnterPassStringPrefix(ctx *PassStringPrefixContext) {}

// ExitPassStringPrefix is called when production passStringPrefix is exited.
func (s *BaseLDEListener) ExitPassStringPrefix(ctx *PassStringPrefixContext) {}

// EnterPassCharPrefix is called when production passCharPrefix is entered.
func (s *BaseLDEListener) EnterPassCharPrefix(ctx *PassCharPrefixContext) {}

// ExitPassCharPrefix is called when production passCharPrefix is exited.
func (s *BaseLDEListener) ExitPassCharPrefix(ctx *PassCharPrefixContext) {}

// EnterMayPassStringPrefix is called when production mayPassStringPrefix is entered.
func (s *BaseLDEListener) EnterMayPassStringPrefix(ctx *MayPassStringPrefixContext) {}

// ExitMayPassStringPrefix is called when production mayPassStringPrefix is exited.
func (s *BaseLDEListener) ExitMayPassStringPrefix(ctx *MayPassStringPrefixContext) {}

// EnterMayPassCharPrefix is called when production mayPassCharPrefix is entered.
func (s *BaseLDEListener) EnterMayPassCharPrefix(ctx *MayPassCharPrefixContext) {}

// ExitMayPassCharPrefix is called when production mayPassCharPrefix is exited.
func (s *BaseLDEListener) ExitMayPassCharPrefix(ctx *MayPassCharPrefixContext) {}

// EnterPassChars is called when production passChars is entered.
func (s *BaseLDEListener) EnterPassChars(ctx *PassCharsContext) {}

// ExitPassChars is called when production passChars is exited.
func (s *BaseLDEListener) ExitPassChars(ctx *PassCharsContext) {}

// EnterPassUntil is called when production passUntil is entered.
func (s *BaseLDEListener) EnterPassUntil(ctx *PassUntilContext) {}

// ExitPassUntil is called when production passUntil is exited.
func (s *BaseLDEListener) ExitPassUntil(ctx *PassUntilContext) {}

// EnterMayPassUntil is called when production mayPassUntil is entered.
func (s *BaseLDEListener) EnterMayPassUntil(ctx *MayPassUntilContext) {}

// ExitMayPassUntil is called when production mayPassUntil is exited.
func (s *BaseLDEListener) ExitMayPassUntil(ctx *MayPassUntilContext) {}

// EnterTakeUntil is called when production takeUntil is entered.
func (s *BaseLDEListener) EnterTakeUntil(ctx *TakeUntilContext) {}

// ExitTakeUntil is called when production takeUntil is exited.
func (s *BaseLDEListener) ExitTakeUntil(ctx *TakeUntilContext) {}

// EnterTakeUntilOrRest is called when production takeUntilOrRest is entered.
func (s *BaseLDEListener) EnterTakeUntilOrRest(ctx *TakeUntilOrRestContext) {}

// ExitTakeUntilOrRest is called when production takeUntilOrRest is exited.
func (s *BaseLDEListener) ExitTakeUntilOrRest(ctx *TakeUntilOrRestContext) {}

// EnterTakeUntilRest is called when production takeUntilRest is entered.
func (s *BaseLDEListener) EnterTakeUntilRest(ctx *TakeUntilRestContext) {}

// ExitTakeUntilRest is called when production takeUntilRest is exited.
func (s *BaseLDEListener) ExitTakeUntilRest(ctx *TakeUntilRestContext) {}

// EnterOptionalNamedArea is called when production optionalNamedArea is entered.
func (s *BaseLDEListener) EnterOptionalNamedArea(ctx *OptionalNamedAreaContext) {}

// ExitOptionalNamedArea is called when production optionalNamedArea is exited.
func (s *BaseLDEListener) ExitOptionalNamedArea(ctx *OptionalNamedAreaContext) {}

// EnterOptionalArea is called when production optionalArea is entered.
func (s *BaseLDEListener) EnterOptionalArea(ctx *OptionalAreaContext) {}

// ExitOptionalArea is called when production optionalArea is exited.
func (s *BaseLDEListener) ExitOptionalArea(ctx *OptionalAreaContext) {}

// EnterAtEnd is called when production atEnd is entered.
func (s *BaseLDEListener) EnterAtEnd(ctx *AtEndContext) {}

// ExitAtEnd is called when production atEnd is exited.
func (s *BaseLDEListener) ExitAtEnd(ctx *AtEndContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseLDEListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseLDEListener) ExitTarget(ctx *TargetContext) {}

// EnterTargetLit is called when production targetLit is entered.
func (s *BaseLDEListener) EnterTargetLit(ctx *TargetLitContext) {}

// ExitTargetLit is called when production targetLit is exited.
func (s *BaseLDEListener) ExitTargetLit(ctx *TargetLitContext) {}

// EnterBound is called when production bound is entered.
func (s *BaseLDEListener) EnterBound(ctx *BoundContext) {}

// ExitBound is called when production bound is exited.
func (s *BaseLDEListener) ExitBound(ctx *BoundContext) {}

// EnterLimit is called when production limit is entered.
func (s *BaseLDEListener) EnterLimit(ctx *LimitContext) {}

// ExitLimit is called when production limit is exited.
func (s *BaseLDEListener) ExitLimit(ctx *LimitContext) {}

// EnterJump is called when production jump is entered.
func (s *BaseLDEListener) EnterJump(ctx *JumpContext) {}

// ExitJump is called when production jump is exited.
func (s *BaseLDEListener) ExitJump(ctx *JumpContext) {}

// EnterExact is called when production exact is entered.
func (s *BaseLDEListener) EnterExact(ctx *ExactContext) {}

// ExitExact is called when production exact is exited.
func (s *BaseLDEListener) ExitExact(ctx *ExactContext) {}

// EnterFieldType is called when production fieldType is entered.
func (s *BaseLDEListener) EnterFieldType(ctx *FieldTypeContext) {}

// ExitFieldType is called when production fieldType is exited.
func (s *BaseLDEListener) ExitFieldType(ctx *FieldTypeContext) {}
