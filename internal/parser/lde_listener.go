// Code generated from LDE.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // LDE

import "github.com/antlr/antlr4/runtime/Go/antlr"

// LDEListener is a complete listener for a parse tree produced by LDEParser.
type LDEListener interface {
	antlr.ParseTreeListener

	// EnterRules is called when entering the rules production.
	EnterRules(c *RulesContext)

	// EnterTypeDeclaration is called when entering the typeDeclaration production.
	EnterTypeDeclaration(c *TypeDeclarationContext)

	// EnterAtomicRule is called when entering the atomicRule production.
	EnterAtomicRule(c *AtomicRuleContext)

	// EnterBaseAction is called when entering the baseAction production.
	EnterBaseAction(c *BaseActionContext)

	// EnterAtomicAction is called when entering the atomicAction production.
	EnterAtomicAction(c *AtomicActionContext)

	// EnterPassHeadingCharacters is called when entering the passHeadingCharacters production.
	EnterPassHeadingCharacters(c *PassHeadingCharactersContext)

	// EnterPassTargetPrefix is called when entering the passTargetPrefix production.
	EnterPassTargetPrefix(c *PassTargetPrefixContext)

	// EnterCheckTargetPrefix is called when entering the checkTargetPrefix production.
	EnterCheckTargetPrefix(c *CheckTargetPrefixContext)

	// EnterMayBePassTargetPrefix is called when entering the mayBePassTargetPrefix production.
	EnterMayBePassTargetPrefix(c *MayBePassTargetPrefixContext)

	// EnterPassChars is called when entering the passChars production.
	EnterPassChars(c *PassCharsContext)

	// EnterGoUntil is called when entering the goUntil production.
	EnterGoUntil(c *GoUntilContext)

	// EnterMayGoUntil is called when entering the mayGoUntil production.
	EnterMayGoUntil(c *MayGoUntilContext)

	// EnterPassUntil is called when entering the passUntil production.
	EnterPassUntil(c *PassUntilContext)

	// EnterMayPassUntil is called when entering the mayPassUntil production.
	EnterMayPassUntil(c *MayPassUntilContext)

	// EnterTakeUntil is called when entering the takeUntil production.
	EnterTakeUntil(c *TakeUntilContext)

	// EnterTakeUntilIncluding is called when entering the takeUntilIncluding production.
	EnterTakeUntilIncluding(c *TakeUntilIncludingContext)

	// EnterTakeUntilOrRest is called when entering the takeUntilOrRest production.
	EnterTakeUntilOrRest(c *TakeUntilOrRestContext)

	// EnterTakeUntilIncludingOrRest is called when entering the takeUntilIncludingOrRest production.
	EnterTakeUntilIncludingOrRest(c *TakeUntilIncludingOrRestContext)

	// EnterTakeUntilRest is called when entering the takeUntilRest production.
	EnterTakeUntilRest(c *TakeUntilRestContext)

	// EnterOptionalNamedArea is called when entering the optionalNamedArea production.
	EnterOptionalNamedArea(c *OptionalNamedAreaContext)

	// EnterOptionalNamedSilentArea is called when entering the optionalNamedSilentArea production.
	EnterOptionalNamedSilentArea(c *OptionalNamedSilentAreaContext)

	// EnterOptionalArea is called when entering the optionalArea production.
	EnterOptionalArea(c *OptionalAreaContext)

	// EnterRestCheck is called when entering the restCheck production.
	EnterRestCheck(c *RestCheckContext)

	// EnterAtEnd is called when entering the atEnd production.
	EnterAtEnd(c *AtEndContext)

	// EnterTarget is called when entering the target production.
	EnterTarget(c *TargetContext)

	// EnterTargetLit is called when entering the targetLit production.
	EnterTargetLit(c *TargetLitContext)

	// EnterBound is called when entering the bound production.
	EnterBound(c *BoundContext)

	// EnterLimit is called when entering the limit production.
	EnterLimit(c *LimitContext)

	// EnterJump is called when entering the jump production.
	EnterJump(c *JumpContext)

	// EnterExact is called when entering the exact production.
	EnterExact(c *ExactContext)

	// EnterFieldType is called when entering the fieldType production.
	EnterFieldType(c *FieldTypeContext)

	// ExitRules is called when exiting the rules production.
	ExitRules(c *RulesContext)

	// ExitTypeDeclaration is called when exiting the typeDeclaration production.
	ExitTypeDeclaration(c *TypeDeclarationContext)

	// ExitAtomicRule is called when exiting the atomicRule production.
	ExitAtomicRule(c *AtomicRuleContext)

	// ExitBaseAction is called when exiting the baseAction production.
	ExitBaseAction(c *BaseActionContext)

	// ExitAtomicAction is called when exiting the atomicAction production.
	ExitAtomicAction(c *AtomicActionContext)

	// ExitPassHeadingCharacters is called when exiting the passHeadingCharacters production.
	ExitPassHeadingCharacters(c *PassHeadingCharactersContext)

	// ExitPassTargetPrefix is called when exiting the passTargetPrefix production.
	ExitPassTargetPrefix(c *PassTargetPrefixContext)

	// ExitCheckTargetPrefix is called when exiting the checkTargetPrefix production.
	ExitCheckTargetPrefix(c *CheckTargetPrefixContext)

	// ExitMayBePassTargetPrefix is called when exiting the mayBePassTargetPrefix production.
	ExitMayBePassTargetPrefix(c *MayBePassTargetPrefixContext)

	// ExitPassChars is called when exiting the passChars production.
	ExitPassChars(c *PassCharsContext)

	// ExitGoUntil is called when exiting the goUntil production.
	ExitGoUntil(c *GoUntilContext)

	// ExitMayGoUntil is called when exiting the mayGoUntil production.
	ExitMayGoUntil(c *MayGoUntilContext)

	// ExitPassUntil is called when exiting the passUntil production.
	ExitPassUntil(c *PassUntilContext)

	// ExitMayPassUntil is called when exiting the mayPassUntil production.
	ExitMayPassUntil(c *MayPassUntilContext)

	// ExitTakeUntil is called when exiting the takeUntil production.
	ExitTakeUntil(c *TakeUntilContext)

	// ExitTakeUntilIncluding is called when exiting the takeUntilIncluding production.
	ExitTakeUntilIncluding(c *TakeUntilIncludingContext)

	// ExitTakeUntilOrRest is called when exiting the takeUntilOrRest production.
	ExitTakeUntilOrRest(c *TakeUntilOrRestContext)

	// ExitTakeUntilIncludingOrRest is called when exiting the takeUntilIncludingOrRest production.
	ExitTakeUntilIncludingOrRest(c *TakeUntilIncludingOrRestContext)

	// ExitTakeUntilRest is called when exiting the takeUntilRest production.
	ExitTakeUntilRest(c *TakeUntilRestContext)

	// ExitOptionalNamedArea is called when exiting the optionalNamedArea production.
	ExitOptionalNamedArea(c *OptionalNamedAreaContext)

	// ExitOptionalNamedSilentArea is called when exiting the optionalNamedSilentArea production.
	ExitOptionalNamedSilentArea(c *OptionalNamedSilentAreaContext)

	// ExitOptionalArea is called when exiting the optionalArea production.
	ExitOptionalArea(c *OptionalAreaContext)

	// ExitRestCheck is called when exiting the restCheck production.
	ExitRestCheck(c *RestCheckContext)

	// ExitAtEnd is called when exiting the atEnd production.
	ExitAtEnd(c *AtEndContext)

	// ExitTarget is called when exiting the target production.
	ExitTarget(c *TargetContext)

	// ExitTargetLit is called when exiting the targetLit production.
	ExitTargetLit(c *TargetLitContext)

	// ExitBound is called when exiting the bound production.
	ExitBound(c *BoundContext)

	// ExitLimit is called when exiting the limit production.
	ExitLimit(c *LimitContext)

	// ExitJump is called when exiting the jump production.
	ExitJump(c *JumpContext)

	// ExitExact is called when exiting the exact production.
	ExitExact(c *ExactContext)

	// ExitFieldType is called when exiting the fieldType production.
	ExitFieldType(c *FieldTypeContext)
}
