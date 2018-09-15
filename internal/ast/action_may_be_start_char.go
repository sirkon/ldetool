package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &MayBeStartChar{}

// MayBeStartChar ...
type MayBeStartChar struct {
	access
	Value string
}

func (m *MayBeStartChar) Accept(d ActionDispatcher) error {
	return d.DispatchMayBeStartChar(m)
}

func (m *MayBeStartChar) String() string {
	return fmt.Sprintf("Pass character \033[1m%s\033[0m if starts with", m.Value)
}

// MayBeStartsWithChar ...
func MayBeStartsWithChar(value antlr.Token) *MayBeStartChar {
	return &MayBeStartChar{Value: value.GetText()}
}
