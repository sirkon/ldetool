package ast

import (
	"fmt"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = PassFixed(0)

// PassFixed ...
type PassFixed int

func (pf PassFixed) Accept(d ActionDispatcher) error {
	return d.DispatchPassFirst(pf)
}

func (pf PassFixed) String() string {
	return fmt.Sprintf("Passing first %d letters", pf)
}

// implements Action for reuse
func (PassFixed) neverCallMe() {}

// PassFirst ...
func PassFirst(field antlr.Token) (res PassFixed, err error) {
	f := field.(antlr.Token)
	value, err := strconv.ParseInt(string(f.GetText()), 10, 64)
	result := PassFixed(value)
	return result, nil
}
