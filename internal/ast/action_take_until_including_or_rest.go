package ast

import (
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &TakeUntilIncludingOrRest{}

// TakeUntilOrRest ...
type TakeUntilIncludingOrRest struct {
	access

	Field Field
	Limit *Target
}

// Accept ...
func (t *TakeUntilIncludingOrRest) Accept(d ActionDispatcher) error {
	return d.DispatchTakeUntilIncludingOrRest(t)
}

func (t *TakeUntilIncludingOrRest) String() string {
	return fmt.Sprintf(
		"Take until %s (or all the rest if not found) as \033[32m%s(%s)\033[0m",
		t.Limit.Value, t.Field.Name, t.Field.Type,
	)
}

// TakeUntilTargetIncludingOrRest ...
func TakeUntilTargetIncludingOrRest(comment []string, field antlr.Token, fieldType antlr.Token) *TakeUntilIncludingOrRest {
	f := NewField(comment, field, fieldType)
	return &TakeUntilIncludingOrRest{
		Field: f,
		Limit: NewTarget(),
	}
}
