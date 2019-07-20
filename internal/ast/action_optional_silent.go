package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &OptionalSilent{}

// OptionalSilent ...
type OptionalSilent struct {
	access

	Name    string
	Actions []Action

	NameToken antlr.Token
}

func (o *OptionalSilent) Accept(d ActionDispatcher) error {
	return d.DispatchOptionalSilent(o)
}

func (o *OptionalSilent) String() string {
	return fmt.Sprintf("Silent named option \033[1m%s\033[0m", o.Name)
}

// OptionSilent ...
func OptionSilent(opt antlr.Token) *OptionalSilent {
	res := &OptionalSilent{
		Name:    opt.GetText(),
		Actions: []Action{},

		NameToken: opt,
	}
	return res
}

// Append ...
func (o *OptionalSilent) Append(i Action) {
	o.Actions = append(o.Actions, i)
}
