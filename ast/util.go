package ast

import (
	"errors"
	"fmt"

	"github.com/sirkon/ldetool/token"
)

// TokenError printer
func TokenError(t *token.Token) func(format string, args ...interface{}) error {
	return func(format string, args ...interface{}) error {
		prefix := fmt.Sprintf("%d:column %d: ", t.Line, t.Column)
		return errors.New(prefix + fmt.Sprintf(format, args...))
	}

}
