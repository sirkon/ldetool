package srcobj

import (
	"fmt"
	"io"
)

// LineAssign represents assignment operation
type LineAssign struct {
	Receiver string
	Expr     Source
}

// Dump ...
func (a LineAssign) Dump(w io.Writer) error {
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

type hardToAccessAssign struct {
	receiver string
	expr     Source
}

func (a hardToAccessAssign) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s = ", a.receiver); err != nil {
		return err
	}
	if err := a.expr.Dump(w); err != nil {
		return err
	}
	return nil
}

func Assign(receiver string, src Source) Source {
	return hardToAccessAssign{
		receiver: receiver,
		expr:     src,
	}
}

// Define represents variable definition
type Define struct {
	Receiver Source
	Expr     Source
}

// Dump ...
func (d Define) Dump(w io.Writer) error {
	if err := d.Receiver.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, ":="); err != nil {
		return err
	}
	if err := d.Expr.Dump(w); err != nil {
		return err
	}
	return nil
}
