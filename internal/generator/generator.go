package generator

import (
	"io"

	"github.com/sirkon/ldetool/internal/ast"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type (
	// Generator describes methods needed of data lookup and extraction
	Generator interface {
		// Data handlers
		AddField(comment []string, name string, t antlr.Token, fieldType string) error
		RegGravity(name string) error

		// Pass
		PassN(n int) error
		PassHeadCharacters(char string) error

		//
		AtEnd() error

		// Head
		HeadString(anchor string, ignore bool, pass bool) error
		HeadChar(char string, ignore bool, pass bool) error

		// Lookups
		LookupString(anchor string, lower, upper int, close, ignore, pass bool) error
		LookupFixedString(anchor string, offset int, ignore, pass bool) error
		LookupChar(anchor string, lower, upper int, close, ignore, pass bool) error
		LookupFixedChar(anchor string, offset int, ignore, pass bool) error

		// Takes
		// Take before anchor (string or character)
		TakeBeforeString(name, fieldType, anchor string, meta ast.FieldMeta, lower, upper int, close, expand, include bool) error
		TakeBeforeChar(name, fieldType, char string, meta ast.FieldMeta, lower, upper int, close, expand, include bool) error

		TakeBeforeStringOnExactPosition(name, fieldType, anchor string, meta ast.FieldMeta, off int, close, expand, include bool) error
		TakeBeforeCharOnExactPosition(name, fieldType, anchor string, meta ast.FieldMeta, off int, close, expand, include bool) error
		// Take the rest
		TakeRest(name, fieldType string, meta ast.FieldMeta) error

		// RestLengthCheck how many symbols left in the rest
		RestLengthCheck(operator string, length int) error

		// Optionals
		OpenOptionalScope(comment []string, name string, t antlr.Token) error
		CloseOptionalScope() error
		OpenSilentOptionalScope(comment []string, name string, t antlr.Token) error
		CloseSilentOptionalScope() error

		// Stress set mismatch treatment as critical error
		Stress()

		// Relax set mismatch error as not critical
		Relax()

		// UseRule ...
		UseRule(comment []string, t antlr.Token, name string) error

		// Push is used to signal all the data for current parser was generated
		Push() error

		// Generate code at the place
		Generate(pkgName string, dest io.Writer) error
		ErrorToken(token antlr.Token, format string, params ...interface{}) error

		// PlatformType to generate code for
		PlatformType(t PlatformType)

		// RegImport go-specific thing for import registration
		RegImport(importAs, path string) error
	}
)
