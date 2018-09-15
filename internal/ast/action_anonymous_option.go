package ast

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &AnonymousOption{}

// AnonymousOption ...
type AnonymousOption struct {
	access

	Name    string
	Actions []Action

	StartToken antlr.Token
}

// Accept ...
func (a *AnonymousOption) Accept(d ActionDispatcher) error {
	return d.DispatchAnonymousOption(a)
}

func (a *AnonymousOption) String() string {
	return "Anonymous option"
}

// Anonymous ...
func Anonymous(opt antlr.Token) *AnonymousOption {
	res := &AnonymousOption{
		Actions: []Action{},

		StartToken: opt,
	}
	return res
}

// Append ...
func (a *AnonymousOption) Append(i Action) {
	a.Actions = append(a.Actions, i)
}
