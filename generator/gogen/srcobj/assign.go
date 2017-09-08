package srcobj

import (
	"fmt"
	"io"
)

// Assign represents assignment operation
type Assign struct {
	Receiver string
	Expr     Source
}

// Dump ...
func (a Assign) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s = ", a.Receiver); err != nil {
		return err
	}
	if err := a.Expr.Dump(w); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}
