package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &Take{}

// Take ...
type Take struct {
	access
	Field Field
	Limit *Target
}

func (t *Take) Accept(d ActionDispatcher) error {
	return d.DispatchTake(t)
}

func (t *Take) String() string {
	return fmt.Sprintf(
		"Take until \033[1m%s\033[0m as \033[32m%s(%s)\033[0m",
		t.Limit.Value, t.Field.Name, t.Field.Type,
	)
}

// TakeUntilTarget ...
func TakeUntilTarget(comment []string, field, fieldType antlr.Token) *Take {
	f := NewField(comment, field, fieldType)
	return &Take{
		Field: f,
		Limit: NewTarget(),
	}
}
