package srcobj

import (
	"io"
)

// LookupByteShort represents short byte lookups
type LookupByteShort struct {
	Var    string
	Src    Source
	Needle Source
}

// Dump ...
func (l LookupByteShort) Dump(w io.Writer) error {
	b := &Body{}
	b.Append(Assign{
		Receiver: l.Var,
		Expr:     Raw("-1"),
	})
	breaking := &Body{}
	breaking.Append(Assign{Receiver: l.Var, Expr: Raw("i")})
	b.Append(For{
		I:         "i",
		Value:     "char",
		Container: l.Src,
		Body: If{
			Expr: OperatorEq(Raw(l.Var), Raw("char")),
			Then: breaking,
		},
	})
	return b.Dump(w)
}

// LookupStringShort look for the distant string
type LookupStringShort struct {
	Var    string
	Src    Source
	Needle Source
}

// Dump ...
func (l LookupStringShort) Dump(w io.Writer) error {
	r := LookupStringLong{
		Var:    l.Var,
		Src:    l.Src,
		Needle: l.Needle,
	}
	return r.Dump(w)
}
