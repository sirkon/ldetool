package generator

import (
	"io"

	"github.com/DenisCheremisov/ldegen/token"
)

// Generator describes methods needed of data lookup and extraction
type Generator interface {
	// Data handlers
	AddField(name string, fieldType string, t *token.Token) (accessName string)

	// Pass
	PassN(n int)

	// Head
	HeadString(anchor string)
	MayBeHeadString(anchor string)
	HeadChar(char string)
	MayBeHeadChar(char string)

	// Lookups
	LookupString(anchor string)
	LookupLimitedString(anchor string, upper int)
	LookupBoundedString(anchor string, lower, upper int)
	LookupChar(anchor string)
	LookupLimitedChar(anchor string, upper int)
	LookupBoundedChar(anchor string, lower, upper int)

	// Takes
	// Take before anchor (string or character)
	TakeBeforeString(name, fieldType, anchor string)
	TakeBeforeLimitedString(name, fieldType, anchor string, upper int)
	TakeBeforeBoundedString(name, fieldType, anchor string, lower, upper int)
	TakeBeforeChar(name, fieldType, char string)
	TakeBeforeLimitedChar(name, fieldType, char string, upper int)
	TakeBeforeBoundedChar(name, fieldType, char string, lower, upper int)

	// Take all
	TakeRest(name, fieldType string)

	// Take before anchor or to the rest
	TakeBeforeStringOrRest(name, fieldType, anchor string)
	TakeBeforeLimitedStringOrRest(name, fieldType, anchor string, upper int)
	TakeBeforeBoundedStringOrRest(name, fieldType, anchor string, lower, upper int)
	TakeBeforeCharOrRest(name, fieldType, char string)
	TakeBeforeLimitedCharOrRest(name, fieldType, char string, upper int)
	TakeBeforeBoundedCharOrRest(name, fieldType, char string, lower, upper int)

	// Optionals
	OpenOptionalScope(name string)
	ExitOptionalScope() // We always know what scope we are in
	CloseOptionalScope()

	// Stress trigger mismatch treatment as serious error
	Stress()

	// Generate code
	Generate(pkgName, parserName string, dest io.Writer)
}
