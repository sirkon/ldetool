/*
file has weird name with z just to be after all other action_*
*/

package ast

// ActionDispatcher is to be used by various action to generate their arbitrary code
type ActionDispatcher interface {
	DispatchAnonymousOption(a *AnonymousOption) error
	DispatchAtEnd(a AtEnd) error
	DispatchRestLengthCheck(a RestLengthCheck) error
	DispatchErrorMismatch(a ErrorOnMismatch) error
	DispatchMayBeStartChar(a *MayBeStartChar) error
	DispatchMayBeStartString(a *MayBeStartString) error
	DispatchOptional(a *Optional) error
	DispatchPassHeadingCharacters(a PassHeadingCharacters) error
	DispatchPassFirst(a PassFixed) error
	DispatchPassAfter(a *PassAfter) error
	DispatchPassAfterOrIgnore(a *PassAfterOrIgnore) error
	DispatchPassBefore(a *PassBefore) error
	DispatchPassBeforeOrIgnore(a *PassBeforeOrIgnore) error
	DispatchStartChar(a *StartChar) error
	DispatchStartCharWithoutPass(a *StartCharWithoutPass) error
	DispatchStartString(a *StartString) error
	DispatchStartStringWithoutPass(a *StartStringWithoutPass) error
	DispatchTake(a *Take) error
	DispatchTakeIncluding(a *TakeIncluding) error
	DispatchTakeRest(a *TakeRest) error
	DispatchTakeUntilOrRest(a *TakeUntilOrRest) error
	DispatchTakeUntilIncludingOrRest(a *TakeUntilIncludingOrRest) error
	DispatchRule(a *Rule) error
}
