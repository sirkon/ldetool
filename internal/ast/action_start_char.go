package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &StartChar{}

// StartChar ...
type StartChar struct {
	access
	Value string
}

func (sc *StartChar) Accept(d ActionDispatcher) error {
	return d.DispatchStartChar(sc)
}

func (sc *StartChar) String() string {
	return fmt.Sprintf("Check and pass character \033[1m%s\033[0m", sc.Value)
}

// StartsWithChar ...
func StartsWithChar(target antlr.Token) *StartChar {
	return &StartChar{Value: target.GetText()}
}
