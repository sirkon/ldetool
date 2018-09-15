package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &StartString{}

// StartString ...
type StartString struct {
	access
	Value string
}

func (ss *StartString) Accept(d ActionDispatcher) error {
	return d.DispatchStartString(ss)
}

func (ss *StartString) String() string {
	return fmt.Sprintf("Check and pass \033[1m%s\033[0m", ss.Value)

}

// StartsWithString constructor
func StartsWithString(target antlr.Token) *StartString {
	return &StartString{Value: target.GetText()}
}
