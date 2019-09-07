package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &TakeUntilOrRest{}

// TakeUntilOrRest ...
type TakeUntilOrRest struct {
	access

	Field Field
	Limit *Target
}

func (t *TakeUntilOrRest) Accept(d ActionDispatcher) error {
	return d.DispatchTakeUntilOrRest(t)
}

func (t *TakeUntilOrRest) String() string {
	return fmt.Sprintf(
		"Take until %s (or all the rest if not found) as \033[32m%s(%s)\033[0m",
		t.Limit.Value, t.Field.Name, t.Field.Type,
	)
}

// TakeUntilTargetOrRest ...
func TakeUntilTargetOrRest(comment []string, field antlr.Token, fieldType antlr.Token) *TakeUntilOrRest {
	f := NewField(comment, field, fieldType)
	return &TakeUntilOrRest{
		Field: f,
		Limit: NewTarget(),
	}
}
