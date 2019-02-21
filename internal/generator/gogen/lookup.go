package gogen

import (
	"fmt"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

func (g *Generator) regRightPkg() error {
	var pkgName string
	if g.useString {
		pkgName = "strings"
	} else {
		pkgName = "bytes"
	}
	return g.regImport("", pkgName)
}

func (g *Generator) regRightVar(name string) error {
	var varType string
	if g.useString {
		varType = "string"
	} else {
		varType = "[]byte"
	}
	return g.regVar(name, varType)
}

// LookupString ...
func (g *Generator) LookupString(anchor string, lower, upper int, close, ignore bool) error {
	if err := g.regVar("pos", "int"); err != nil {
		return err
	}
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}
	if err := g.regRightPkg(); err != nil {
		return err
	}

	constName := g.constNameFromContent(anchor)

	var rest srcobj.Source
	switch {
	case lower > 0 && upper > 0:
		if err := g.regRightVar("tmpRest"); err != nil {
			return err
		}
		rest = srcobj.Literal("tmpRest")

	case lower == 0 && upper > 0:
		if err := g.regRightVar("tmpRest"); err != nil {
			return err
		}
		rest = srcobj.Literal("tmpRest")

	case lower > 0 && upper == 0:
		rest = srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))

	default:
		rest = srcobj.Raw(g.curRestVar())
	}

	body := g.indent()

	var lookup srcobj.Source
	if upper > 0 {
		if lower > 0 {
			body.Append(
				srcobj.If{
					Expr: srcobj.OperatorGE(
						srcobj.NewCall("len", g.rest()),
						srcobj.Literal(upper),
					),
					Then: srcobj.Assign("tmpRest", srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower), srcobj.Literal(upper))),
					Else: srcobj.Assign("tmpRest", srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))),
				},
			)
		} else {
			body.Append(
				srcobj.If{
					Expr: srcobj.OperatorGE(
						srcobj.NewCall("len", g.rest()),
						srcobj.Literal(upper),
					),
					Then: srcobj.Assign("tmpRest", srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Literal(upper))),
					Else: srcobj.Assign("tmpRest", g.rest()),
				},
			)
		}
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
		lookup = srcobj.LookupStringShort(g.useString, "pos", rest, srcobj.Raw(constName))
	} else {
		g.regRightPkg()
		var detector srcobj.Source = srcobj.LookupStringLong(g.useString, "pos", rest, srcobj.Raw(constName))
		lookup = srcobj.NewBody(srcobj.Trim(detector), srcobj.Raw("\n"))
	}
	body.Append(lookup)

	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`",
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
	return nil
}

// LookupFixedString ...
func (g *Generator) LookupFixedString(anchor string, offset int, ignore bool) error {
	return g.checkStringPrefix(anchor, offset, ignore)
}

// LookupCharEx ...
func (g *Generator) LookupChar(char string, lower, upper int, close, ignore bool) error {
	if err := g.regVar("pos", "int"); err != nil {
		return err
	}
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}

	var rest srcobj.Source
	switch {
	case lower > 0 && upper > 0 && lower != upper:
		if err := g.regRightVar("tmpRest"); err != nil {
			return err
		}
		rest = srcobj.Literal("tmpRest")

	case lower == 0 && upper > 0:
		if err := g.regRightVar("tmpRest"); err != nil {
			return err
		}
		rest = srcobj.Literal("tmpRest")

	case lower > 0 && upper == 0:
		rest = srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))

	default:
		rest = srcobj.Raw(g.curRestVar())
	}

	body := g.indent()
	body.Append(srcobj.Comment(fmt.Sprintf("Looking for %s and then pass it", char)))
	if upper > 0 {
		if lower > 0 {
			body.Append(
				srcobj.If{
					Expr: srcobj.OperatorGE(
						srcobj.NewCall("len", g.rest()),
						srcobj.Literal(upper),
					),
					Then: srcobj.Assign("tmpRest", srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower), srcobj.Literal(upper))),
					Else: srcobj.Assign("tmpRest", srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower))),
				},
			)
		} else {
			body.Append(
				srcobj.If{
					Expr: srcobj.OperatorGE(
						srcobj.NewCall("len", g.rest()),
						srcobj.Literal(upper),
					),
					Then: srcobj.Assign("tmpRest", srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Literal(upper))),
					Else: srcobj.Assign("tmpRest", g.rest()),
				},
			)
		}
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
		g.regRightPkg()
		body.Append(srcobj.LookupByteLong(g.useString, "pos", rest, srcobj.Raw(char)))
	}
	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"cannot find \033[1m%s\033[0m in `\033[1m%s\033[0m`",
			srcobj.DrawChar(char),
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
	return nil
}

// LookupFixedChar ...
func (g *Generator) LookupFixedChar(anchor string, offset int, ignore bool) error {
	return g.checkCharPrefix(anchor, offset, ignore)
}
