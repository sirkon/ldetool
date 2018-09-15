package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &TakeRest{}

// TakeRest ...
type TakeRest struct {
	access
	Field Field
}

func (tr *TakeRest) Accept(d ActionDispatcher) error {
	return d.DispatchTakeRest(tr)
}

func (tr *TakeRest) String() string {
	return fmt.Sprintf("Take the rest as \033[32m%s(%s)\033[0m", tr.Field.Name, tr.Field.Type)
}

// TakeTheRest ...
func TakeTheRest(field, fieldType antlr.Token) *TakeRest {
	f := NewField(field, fieldType)
	res := &TakeRest{
		Field: f,
	}
	return res
}
