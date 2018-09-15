package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Field ...
type Field struct {
	Name string
	Type string

	NameToken antlr.Token
	TypeToken antlr.Token
}

// NewField constructor
func NewField(name antlr.Token, typeToken antlr.Token) Field {
	return Field{
		Name: string(name.GetText()),
		Type: string(typeToken.GetText()),

		NameToken: name,
		TypeToken: typeToken,
	}
}
