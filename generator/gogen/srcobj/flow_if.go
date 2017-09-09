package srcobj

import (
	"io"
)

// If structure
type If struct {
	Expr Source
	Then Source
	Else Source
}

func (i If) Dump(w io.Writer) error {
	if _, err := io.WriteString(w, "if "); err != nil {
		return err
	}
	if err := i.Expr.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, " {\n"); err != nil {
		return err
	}
	if err := i.Then.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "}"); err != nil {
		return err
	}
	if i.Else != nil {
		if _, err := io.WriteString(w, " else {\n"); err != nil {
			return err
		}
		if err := i.Else.Dump(w); err != nil {
			return err
		}
		if _, err := io.WriteString(w, "}"); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}
	return nil
}
