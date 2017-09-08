package srcobj

import "io"

// Index generates indexing access
type Index struct {
	Src   Source
	Index Source
}

// Dump ...
func (i Index) Dump(w io.Writer) error {
	if err := i.Src.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "["); err != nil {
		return err
	}
	if err := i.Index.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "]"); err != nil {
		return err
	}
	return nil
}
