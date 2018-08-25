package srcobj

import (
	"fmt"
	"io"
)

// DrawChar draws char 'a' as "'a'" in Go code
type DrawChar string

// Dump to implement Source
func (dc DrawChar) Dump(w io.Writer) error {
	_, err := fmt.Fprintf(w, `"%s"`, string(dc))
	return err
}
