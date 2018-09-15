package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &Optional{}

// Optional ...
type Optional struct {
	access

	Name    string
	Actions []Action

	NameToken antlr.Token
}

func (o *Optional) Accept(d ActionDispatcher) error {
	return d.DispatchOptional(o)
}

func (o *Optional) String() string {
	return fmt.Sprintf("Option \033[1m%s\033[0m", o.Name)
}

// Option ...
func Option(opt antlr.Token) *Optional {
	res := &Optional{
		Name:    opt.GetText(),
		Actions: []Action{},

		NameToken: opt,
	}
	return res
}

// Append ...
func (o *Optional) Append(i Action) {
	o.Actions = append(o.Actions, i)
}
