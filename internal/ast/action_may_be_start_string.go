package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &MayBeStartString{}

// MayBeStartString ...
type MayBeStartString struct {
	access
	Value string
}

func (m *MayBeStartString) Accept(d ActionDispatcher) error {
	return d.DispatchMayBeStartString(m)
}

func (m *MayBeStartString) String() string {
	return fmt.Sprintf("Pass \033[1m%s\033[0m if starts with", m.Value)

}

// MayBeStartsWithString ...
func MayBeStartsWithString(value antlr.Token) *MayBeStartString {
	return &MayBeStartString{Value: value.GetText()}
}
