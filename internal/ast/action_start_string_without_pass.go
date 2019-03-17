package ast

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var _ Action = &StartStringWithoutPass{}

// StartStringWithoutPass ...
type StartStringWithoutPass struct {
	access
	Value string
}

func (ss *StartStringWithoutPass) Accept(d ActionDispatcher) error {
	return d.DispatchStartStringWithoutPass(ss)
}

func (ss *StartStringWithoutPass) String() string {
	return fmt.Sprintf("Check and pass \033[1m%s\033[0m", ss.Value)

}

// StartsWithStringWithoutPass constructor
func StartsWithStringWithoutPass(target antlr.Token) *StartStringWithoutPass {
	return &StartStringWithoutPass{Value: target.GetText()}
}
