package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &TakeIncluding{}

// TakeIncluding ...
type TakeIncluding struct {
	access
	Field Field
	Limit *Target
}

func (t *TakeIncluding) String() string {
	return fmt.Sprintf(
		"Take until \033[1m%s\033[0m including it as \033[32m%s(%s)\033[0m",
		t.Limit.Value, t.Field.Name, t.Field.Type,
	)
}

// Accept ...
func (t *TakeIncluding) Accept(d ActionDispatcher) error {
	return d.DispatchTakeIncluding(t)
}

// TakeUntilTargetIncluding ...
func TakeUntilTargetIncluding(field, fieldType antlr.Token) *TakeIncluding {
	f := NewField(field, fieldType)
	return &TakeIncluding{
		Field: f,
		Limit: NewTarget(),
	}
}
