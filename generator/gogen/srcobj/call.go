package srcobj

import (
	"fmt"
	"io"
)

// Call function call representation
type Call struct {
	Name   string
	Params []Source
}

func NewCall(name string, params ...Source) Call {
	return Call{
		Name:   name,
		Params: params,
	}
}

// Dump ...
func (c Call) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s(", c.Name); err != nil {
		return err
	}
	for _, param := range c.Params {
		if err := param.Dump(w); err != nil {
			return err
		}
		if _, err := fmt.Fprintf(w, ", "); err != nil {
			return err
		}
	}
	if _, err := io.WriteString(w, ")"); err != nil {
		return err
	}
	return nil
}
