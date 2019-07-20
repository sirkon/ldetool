package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Field ...
type Field struct {
	Name string
	Type string
	Meta FieldMeta

	NameToken antlr.Token
	TypeToken antlr.Token
}

// FieldMeta additional type data, introduced specially for decimal types
type FieldMeta struct {
	Precision int
	Scale     int
}

// NewField constructor
func NewField(name antlr.Token, typeToken antlr.Token) Field {
	de := decimalExtractor{}

	var meta FieldMeta

	typeName := typeToken.GetText()
	switch typeName {
	case "dec32", "dec64", "dec128":
		panic(&ErrorListener{
			Line: typeToken.GetLine(),
			Col:  typeToken.GetColumn(),
			Msg:  fmt.Sprintf("decimal types should be used as decX.Y with precision X and scale Y"),
		})
	}
	ok, err := de.Extract(typeName)
	if err != nil {
		// type name starts with dec, thus it must be decX.Y with these constraints X≥Y, 1≤X≤38
		panic(&ErrorListener{
			Line: typeToken.GetLine(),
			Col:  typeToken.GetColumn(),
			Msg:  fmt.Sprintf("error decimal type, must be decX.Y with integer X and Y"),
		})
	}
	if ok {
		if de.Scale > de.Precision {
			panic(&ErrorListener{
				Line: typeToken.GetLine(),
				Col:  typeToken.GetColumn(),
				Msg:  fmt.Sprintf("error decimal type, must be decX.Y, where integer Y ≤ X"),
			})
		}
		if de.Precision == 0 || de.Precision > 38 {
			panic(&ErrorListener{
				Line: typeToken.GetLine(),
				Col:  typeToken.GetColumn(),
				Msg:  fmt.Sprintf("error decimal type, must be decX.Y, where 1 ≤ X ≤ 38"),
			})
		}
		meta.Precision = int(de.Precision)
		meta.Scale = int(de.Scale)
		switch {
		case de.Precision < 10:
			typeName = "dec32"
		case de.Precision < 19:
			typeName = "dec64"
		case de.Precision < 39:
			typeName = "dec128"
		default:
			panic("must not be here")
		}
	}

	return Field{
		Name: string(name.GetText()),
		Type: typeName,

		Meta:      meta,
		NameToken: name,
		TypeToken: typeToken,
	}
}
