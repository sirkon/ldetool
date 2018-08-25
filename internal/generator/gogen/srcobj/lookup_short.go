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
	b.Append(LineAssign{
		Receiver: l.Var,
		Expr:     Raw("-1"),
	})
	breaking := &Body{}
	breaking.Append(LineAssign{Receiver: l.Var, Expr: Raw("i")})
	breaking.Append(Break)
	b.Append(For{
		I:         "i",
		Value:     "char",
		Container: l.Src,
		Body: If{
			Expr: OperatorEq(Raw("char"), l.Needle),
			Then: breaking,
		},
	})
	return b.Dump(w)
}

// lookupStringShort look for the distant string
type lookupStringShort struct {
	useString bool
	Var       string
	Src       Source
	Needle    Source
}

// LookupStringShort creates private lookupStringShort for external consumption
func LookupStringShort(useString bool, v string, src, needle Source) lookupStringShort {
	return lookupStringShort{
		useString: useString,
		Var:       v,
		Src:       src,
		Needle:    needle,
	}
}

// Dump ...
func (l lookupStringShort) Dump(w io.Writer) error {
	r := LookupStringLong(l.useString, l.Var, l.Src, l.Needle)
	return r.Dump(w)
}
