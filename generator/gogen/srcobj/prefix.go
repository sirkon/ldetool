package srcobj

import "io"

// PrefixByte for prefix byte check code generation
type PrefixByte struct {
	Var    string
	Src    Source
	Needle Source
}

// Dump ...
func (p PrefixByte) Dump(w io.Writer) error {
	src := LineAssign{
		Receiver: "ok",
		Expr: OperatorAnd(
			OperatorGT(Call{Name: "len", Params: []Source{p.Src}}, Raw("0")),
			OperatorEq(Index{Src: p.Src, Index: Raw("0")}, p.Needle),
		),
	}
	return src.Dump(w)
}

func RightPkg(useString bool) string {
	if useString {
		return "strings"
	}
	return "bytes"
}

func RightType(useString bool) string {
	if useString {
		return "string"
	}
	return "[]byte"
}

// prefixString for prefix string check code generation
type prefixString struct {
	useString bool
	Src       Source
	Needle    Source
}

// PrefixString creates private prefixString for external consumption
func PrefixString(useString bool, src, needle Source) prefixString {
	return prefixString{
		useString: useString,
		Src:       src,
		Needle:    needle,
	}
}

// Dump ...
func (p prefixString) Dump(w io.Writer) error {
	src := LineAssign{
		Receiver: "ok",
		Expr:     Call{Name: RightPkg(p.useString) + ".HasPrefix", Params: []Source{p.Src, p.Needle}},
	}
	return src.Dump(w)
}

// prefixStringShort when it is known prefix is short
type prefixStringShort struct {
	Src    Source
	Needle Source
}

// Dump ...
func (p prefixStringShort) Dump(w io.Writer) error {
	body := &Body{}
	body.Append(LineAssign{Receiver: "ok", Expr: Raw("true")})
	body.Append(If{
		Expr: OperatorGE(NewCall("len", p.Src), NewCall("len", p.Needle)),
		Then: For{
			I:         "i",
			Value:     "char",
			Container: p.Needle,
			Body: If{
				Expr: OperatorNEq(Raw("char"), Index{Src: p.Src, Index: Raw("i")}),
				Then: NewBody(
					LineAssign{Receiver: "ok", Expr: Raw("false")},
					Break,
				),
			},
		},
		Else: LineAssign{Receiver: "ok", Expr: Raw("false")},
	})
	return body.Dump(w)
}
