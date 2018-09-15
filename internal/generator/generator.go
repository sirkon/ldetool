package generator

import (
	"io"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Generator describes methods needed of data lookup and extraction
type Generator interface {
	// Data handlers
	AddField(name string, fieldType string, t antlr.Token) error
	RegGravity(name string) error

	// Pass
	PassN(n int) error

	//
	AtEnd() error

	// Head
	HeadString(anchor string, ignore bool) error
	HeadChar(char string, ignore bool) error

	// Lookups
	LookupString(anchor string, lower, upper int, close, ignore bool) error
	LookupFixedString(anchor string, offset int, ignore bool) error
	LookupChar(anchor string, lower, upper int, close, ignore bool) error
	LookupFixedChar(anchor string, offset int, ignore bool) error

	// Takes
	// Take before anchor (string or character)
	TakeBeforeString(name, fieldType, anchor string, lower, upper int, close, expand bool) error
	TakeBeforeChar(name, fieldType, char string, lower, upper int, close, expand bool) error

	// Take the rest
	TakeRest(name, fieldType string) error

	// Optionals
	OpenOptionalScope(name string, t antlr.Token) error
	CloseOptionalScope() error

	// Stress set mismatch treatment as critical error
	Stress() error

	// Relax set mismatch error as not critical
	Relax() error

	// UseRule ...
	UseRule(name string, t antlr.Token) error

	// Push is used to signal all the data for current parser was generated
	Push() error

	// Generate code at the place
	Generate(pkgName string, dest io.Writer) error
	ErrorToken(token antlr.Token, format string, params ...interface{}) error

	// PlatformType to generate code for
	PlatformType(t PlatformType) error
}
