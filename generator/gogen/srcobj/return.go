package srcobj

import (
	"io"
)

const (
	ReturnOK   = Raw("return true, nil")
	ReturnFail = Raw("return false, nil")
)

// hardToAccessReturnError ...
type hardToAccessReturnError struct {
	err Source
}

// Dump ...
func (r hardToAccessReturnError) Dump(w io.Writer) error {
	if _, err := io.WriteString(w, "return false, "); err != nil {
		return err
	}
	if err := r.err.Dump(w); err != nil {
		return err
	}
	if _, err := io.WriteString(w, "\n"); err != nil {
		return err
	}
	return nil
}

// ReturnError ...
func ReturnError(format string, params ...Source) Source {
	return hardToAccessReturnError{err: Error(format, params...)}
}
