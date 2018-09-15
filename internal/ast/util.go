package ast

import (
	"errors"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// TokenError printer
func TokenError(t antlr.Token) func(format string, args ...interface{}) error {
	return func(format string, args ...interface{}) error {
		prefix := fmt.Sprintf("%d:column %d: ", t.GetLine(), t.GetColumn())
		return errors.New(prefix + fmt.Sprintf(format, args...))
	}

}

func posLit(i int) string {
	switch i {
	case 1:
		return "1st"
	case 2:
		return "2nd"
	case 3:
		return "3rd"
	default:
		return fmt.Sprintf("%dth", i)
	}
}
