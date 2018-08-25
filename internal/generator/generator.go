package generator

import (
	"io"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Generator describes methods needed of data lookup and extraction
type Generator interface {
	// Data handlers
	AddField(name string, fieldType string, t antlr.Token)
	RegGravity(name string)

	// Pass
	PassN(n int)

	//
	AtEnd()

	// Head
	HeadString(anchor string, ignore bool)
	HeadChar(char string, ignore bool)

	// Lookups
	LookupString(anchor string, lower, upper int, close, ignore bool)
	LookupFixedString(anchor string, offset int, ignore bool)
	LookupChar(anchor string, lower, upper int, close, ignore bool)
	LookupFixedChar(anchor string, offset int, ignore bool)

	// Takes
	// Take before anchor (string or character)
	TakeBeforeString(name, fieldType, anchor string, lower, upper int, close, expand bool)
	TakeBeforeChar(name, fieldType, char string, lower, upper int, close, expand bool)

	// Take the rest
	TakeRest(name, fieldType string)

	// Optionals
	OpenOptionalScope(name string, t antlr.Token)
	CloseOptionalScope()

	// Stress set mismatch treatment as critical error
	Stress()

	// Relax set mismatch error as not critical
	Relax()

	// UseRule ...
	UseRule(name string, t antlr.Token)

	// Push is used to signal all the data for current parser was generated
	Push()

	// Generate code at the place
	Generate(pkgName string, dest io.Writer)
	ErrorToken(token antlr.Token, format string, params ...interface{})

	// PlatformType to generate code for
	PlatformType(t PlatformType)
}
