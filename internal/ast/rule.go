package ast

import (
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &Rule{}

// Rule is action as well
type Rule struct {
	access

	Name    string
	Comment []string
	Actions []Action

	NameToken antlr.Token
}

func (r *Rule) Accept(d ActionDispatcher) error {
	return d.DispatchRule(r)
}

func (r *Rule) String() string {
	res := &strings.Builder{}
	res.WriteString("rule ")
	res.WriteString(r.Name)
	for _, action := range r.Actions {
		res.WriteString("\n")
		res.WriteString(action.String())
	}
	return res.String()
}

// NewRule constructor
func NewRule(comment []string, name antlr.Token) *Rule {
	n := name.(antlr.Token)
	res := &Rule{
		Name:      string(n.GetText()),
		Comment:   comment,
		NameToken: n,
	}
	return res
}

func (r *Rule) Append(ai Action) {
	r.Actions = append(r.Actions, ai)
}
