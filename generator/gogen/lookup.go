package gogen

import (
	"fmt"

	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

// LookupString ...
func (g *Generator) LookupString(anchor string, lower, upper int, close, ignore bool) {
	g.regVar("pos", "int")
	g.regVar(g.curRestVar(), "[]byte")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)

	var rest srcobj.Source
	switch {
	case lower > 0 && upper > 0:
		rest = srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower), srcobj.Literal(upper))

	case lower == 0 && upper > 0:
		rest = srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Literal(upper))

	case lower > 0 && upper == 0:
		rest = srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))

	default:
		rest = srcobj.Raw(g.curRestVar())
	}

	body := g.indent()

	var lookup srcobj.Source
	if upper > 0 {
		body.Append(
			srcobj.If{
				Expr: srcobj.OperatorLT(
					srcobj.NewCall("len", g.rest()),
					srcobj.Literal(upper),
				),
				Then: g.sliceTooLarge(upper),
			},
		)
	} else if lower > 0 {
		body.Append(
			srcobj.If{
				Expr: srcobj.OperatorLT(
					srcobj.NewCall("len", g.rest()),
					srcobj.Literal(lower),
				),
				Then: g.jumpTooLarge(lower),
			},
		)
	}

	body.Append(srcobj.Comment(fmt.Sprintf("Looking for %s and then pass it", anchor)))
	if close {
		lookup = srcobj.LookupStringShort{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(constName),
		}
	} else {
		g.regImport("", "bytes")
		var detector srcobj.Source = srcobj.LookupStringLong{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(constName),
		}
		lookup = srcobj.NewBody(srcobj.Trim(detector), srcobj.Raw("\n"))
	}
	body.Append(lookup)

	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`",
			srcobj.Raw(constName),
			srcobj.Stringify(rest),
		)
	}

	var offset srcobj.Source
	if lower == 0 {
		offset = srcobj.OperatorAdd(
			srcobj.Raw("pos"),
			srcobj.NewCall("len", srcobj.Raw(constName)),
		)
	} else {
		l := fmt.Sprintf("%d", lower)
		offset = srcobj.OperatorAdd(
			srcobj.Raw("pos"),
			srcobj.OperatorAdd(
				srcobj.NewCall("len", srcobj.Raw(constName)),
				srcobj.Raw(l),
			),
		)
	}

	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: srcobj.LineAssign{
			Receiver: g.curRestVar(),
			Expr: srcobj.SliceFrom(
				srcobj.Raw(g.curRestVar()),
				offset,
			),
		},
		Else: failure,
	})
}

// LookupFixedString ...
func (g *Generator) LookupFixedString(anchor string, offset int, ignore bool) {
	g.checkStringPrefix(anchor, offset, ignore)
}

// LookupCharEx ...
func (g *Generator) LookupChar(char string, lower, upper int, close, ignore bool) {
	g.regVar("pos", "int")
	g.regVar(g.curRestVar(), "[]byte")

	var rest srcobj.Source
	switch {
	case lower > 0 && upper > 0:
		rest = srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower), srcobj.Literal(upper))

	case lower == 0 && upper > 0:
		rest = srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Literal(upper))

	case lower > 0 && upper == 0:
		rest = srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))

	default:
		rest = srcobj.Raw(g.curRestVar())
	}

	body := g.indent()
	body.Append(srcobj.Comment(fmt.Sprintf("Looking for %s and then pass it", char)))
	if upper > 0 {
		body.Append(
			srcobj.If{
				Expr: srcobj.OperatorLT(
					srcobj.NewCall("len", g.rest()),
					srcobj.Literal(upper),
				),
				Then: g.sliceTooLarge(upper),
			},
		)
	} else if lower > 0 {
		body.Append(
			srcobj.If{
				Expr: srcobj.OperatorLT(
					srcobj.NewCall("len", g.rest()),
					srcobj.Literal(lower),
				),
				Then: g.jumpTooLarge(lower),
			},
		)
	}
	if close {
		body.Append(srcobj.LookupByteShort{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(char),
		})
	} else {
		g.regImport("", "bytes")
		body.Append(srcobj.LookupByteLong{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(char),
		})
	}
	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`",
			srcobj.Raw(char),
			srcobj.Stringify(rest),
		)
	}

	var offset srcobj.Source
	if lower <= 0 {
		offset = srcobj.OperatorAdd(
			srcobj.Raw("pos"),
			srcobj.Raw("1"),
		)
	} else {
		l := fmt.Sprintf("%d", lower)
		offset = srcobj.OperatorAdd(
			srcobj.Raw("pos"),
			srcobj.OperatorAdd(
				srcobj.Raw("1"),
				srcobj.Raw(l),
			),
		)
	}

	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: srcobj.LineAssign{
			Receiver: g.curRestVar(),
			Expr: srcobj.SliceFrom(
				srcobj.Raw(g.curRestVar()),
				offset,
			),
		},
		Else: failure,
	})
}

// LookupFixedChar ...
func (g *Generator) LookupFixedChar(anchor string, offset int, ignore bool) {
	g.checkCharPrefix(anchor, offset, ignore)
}
