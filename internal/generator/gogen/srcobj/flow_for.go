package srcobj

import (
	"fmt"
	"io"
)

// For range representation
type For struct {
	I         string
	Value     string
	Container Source
	Body      Source
}

// Dump ...
func (f For) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "for %s, %s := range ", f.I, f.Value); err != nil {
		return err
	}
	if err := f.Container.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, " {\n"); err != nil {
		return err
	}
	if err := f.Body.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "}\n"); err != nil {
		return err
	}
	return nil
}
