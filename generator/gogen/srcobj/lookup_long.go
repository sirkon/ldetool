package srcobj

import "io"

// lookupByteLong look for the distant character
type lookupByteLong struct {
	useString bool
	Var       string
	Src       Source
	Needle    Source
}

// LookupByteLong creates new lookup
func LookupByteLong(useString bool, v string, src Source, needle Source) lookupByteLong {
	return lookupByteLong{
		useString: useString,
		Var:       v,
		Src:       src,
		Needle:    needle,
	}
}

// Dump ...
func (l lookupByteLong) Dump(w io.Writer) error {
	ass := LineAssign{
		Receiver: l.Var,
		Expr: Call{
			Name:   RightPkg(l.useString) + ".IndexByte",
			Params: []Source{l.Src, l.Needle},
		},
	}
	return ass.Dump(w)
}

// lookupStringLong look for the distant string
type lookupStringLong struct {
	useString bool
	Var       string
	Src       Source
	Needle    Source
}

// LookupStringLong creates new lookup
func LookupStringLong(useString bool, v string, src Source, needle Source) lookupStringLong {
	return lookupStringLong{
		useString: useString,
		Var:       v,
		Src:       src,
		Needle:    needle,
	}
}

// Dump ...
func (l lookupStringLong) Dump(w io.Writer) error {
	ass := LineAssign{
		Receiver: l.Var,
		Expr: Call{
			Name:   RightPkg(l.useString) + ".Index",
			Params: []Source{l.Src, l.Needle},
		},
	}
	return ass.Dump(w)
}
