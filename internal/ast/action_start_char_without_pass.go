package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &StartCharWithoutPass{}

// StartCharWithoutPass ...
type StartCharWithoutPass struct {
	access
	Value string
}

func (sc *StartCharWithoutPass) Accept(d ActionDispatcher) error {
	return d.DispatchStartCharWithoutPass(sc)
}

func (sc *StartCharWithoutPass) String() string {
	return fmt.Sprintf("Check and pass character \033[1m%s\033[0m", sc.Value)
}

// StartsWithCharWithoutPass ...
func StartsWithCharWithoutPass(target antlr.Token) *StartCharWithoutPass {
	return &StartCharWithoutPass{Value: target.GetText()}
}
