package srcobj

import "io"

// LookupByteLong look for the distant character
type LookupByteLong struct {
	Var    string
	Src    Source
	Needle Source
}

// Dump ...
func (l LookupByteLong) Dump(w io.Writer) error {
	ass := LineAssign{
		Receiver: l.Var,
		Expr: Call{
			Name:   "bytes.IndexByte",
			Params: []Source{l.Src, l.Needle},
		},
	}
	return ass.Dump(w)
}

// LookupStringLong look for the distant string
type LookupStringLong struct {
	Var    string
	Src    Source
	Needle Source
}

// Dump ...
func (l LookupStringLong) Dump(w io.Writer) error {
	ass := LineAssign{
		Receiver: l.Var,
		Expr: Call{
			Name:   "bytes.Index",
			Params: []Source{l.Src, l.Needle},
		},
	}
	return ass.Dump(w)
}
