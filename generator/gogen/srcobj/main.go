package srcobj

import (
	"bytes"
	"io"
)

// Source is an abstraction over
type Source interface {
	Dump(w io.Writer) error
}

func String(s Source) string {
	w := &bytes.Buffer{}
	if err := s.Dump(w); err != nil {
		panic(err)
	}
	return w.String()
}
