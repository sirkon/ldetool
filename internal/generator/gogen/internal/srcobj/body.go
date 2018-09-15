package srcobj

import "io"

// Body describes LDE generated listing of program elements
type Body struct {
	l []Source
}

func NewBody(pieces ...Source) *Body {
	return &Body{
		l: pieces,
	}
}

// Append appends new program element
func (l *Body) Append(item Source) {
	l.l = append(l.l, item)
}

// Dump implementation
func (l *Body) Dump(w io.Writer) error {
	for _, item := range l.l {
		if err := item.Dump(w); err != nil {
			return err
		}
	}
	return nil
}
