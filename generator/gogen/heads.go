package gogen

import (
	"strings"

	"fmt"

	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

func (g *Generator) label() string {
	return g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
}

func (g *Generator) checkStringPrefix(anchor string, offset int, ignore bool) {
	g.regVar(g.curRestVar(), "[]byte")
	g.regImport("", "bytes")
	constName := g.constNameFromContent(anchor)

	var rest srcobj.Source = srcobj.Raw(g.curRestVar())
	if offset != 0 {
		rest = srcobj.SliceFrom(rest, srcobj.Raw(fmt.Sprintf("%d", offset)))
	}

	var shift srcobj.Source = srcobj.NewCall("len", srcobj.Raw(constName))
	if offset != 0 {
		shift = srcobj.OperatorAdd(shift, srcobj.Raw(fmt.Sprintf("%d", offset)))
	}

	var failure srcobj.Source
	if !ignore {
		if len(g.namespaces) > 0 {
			failure = srcobj.Goto(g.label())
		} else if g.serious {
			g.regImport("", "fmt")
			failure = srcobj.ReturnError(
				"`\033[1m%s\033[0m)` is expected to start with `\033[1m%s\033[0m`", rest, srcobj.Raw(anchor))
		} else {
			failure = srcobj.ReturnFail
		}
	}

	body := srcobj.NewBody(srcobj.Raw("\n"))
	if offset > 0 {
		body.Append(
			srcobj.Comment(fmt.Sprintf("Checks if rest[%d:] starts with `%s` and pass it", offset, anchor)))
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Checks if the rest starts with `%s` and pass it", anchor)))
	}

	var cond srcobj.Source = srcobj.NewCall("bytes.HasPrefix", rest, srcobj.Raw(constName))
	if offset > 0 {
		cond = srcobj.OperatorAnd(
			srcobj.OperatorGE(
				srcobj.NewCall("len", rest),
				srcobj.OperatorAdd(
					srcobj.Literal(offset),
					srcobj.NewCall("len", srcobj.Raw(constName)),
				),
			),
			cond,
		)
	}

	body.Append(srcobj.If{
		Expr: cond,
		Then: srcobj.Assign{
			Receiver: g.curRestVar(),
			Expr:     srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), shift),
		},
		Else: failure,
	})

	if err := body.Dump(g.curBody); err != nil {
		return
	}
}

// HeadString checks if the rest starts with the given string and passes it
func (g *Generator) HeadString(anchor string, ignore bool) {
	g.checkStringPrefix(anchor, 0, ignore)
}

func (g *Generator) checkCharPrefix(char string, offset int, ignore bool) {
	g.regVar(g.curRestVar(), "[]byte")

	var rest srcobj.Source = srcobj.Raw(g.curRestVar())

	var shift srcobj.Source = srcobj.Literal(1)
	if offset != 0 {
		shift = srcobj.OperatorAdd(srcobj.Literal(offset), shift)
	}

	var failure srcobj.Source
	if !ignore {
		if len(g.namespaces) > 0 {
			failure = srcobj.Goto(g.label())
		} else if g.serious {
			g.regImport("", "fmt")
			failure = srcobj.ReturnError(
				"`\033[1m%s\033[0m)` is expected to start with \033[1m%s\033[0m", rest, srcobj.Raw(char))
		} else {
			failure = srcobj.ReturnFail
		}
	}

	body := srcobj.NewBody(srcobj.Raw("\n"))
	if offset > 0 {
		body.Append(
			srcobj.Comment(fmt.Sprintf("Checks if rest[%d:] starts with %s and pass it", offset, char)))
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Checks if the rest starts with %s and pass it", char)))
	}

	var cond srcobj.Source
	if offset > 0 {
		cond = srcobj.OperatorGE(
			srcobj.NewCall("len", rest),
			srcobj.OperatorAdd(
				srcobj.Literal(offset),
				srcobj.Literal(1),
			),
		)
	} else {
		cond = srcobj.OperatorGE(
			srcobj.NewCall("len", rest),
			srcobj.Literal(1),
		)
	}
	cond = srcobj.OperatorAnd(
		cond,
		srcobj.OperatorEq(
			srcobj.Index{
				Src:   rest,
				Index: srcobj.Literal(offset),
			},
			srcobj.Raw(char),
		),
	)

	body.Append(srcobj.If{
		Expr: cond,
		Then: srcobj.Assign{
			Receiver: g.curRestVar(),
			Expr:     srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), shift),
		},
		Else: failure,
	})

	if err := body.Dump(g.curBody); err != nil {
		return
	}
}

// HeadChar checks if rest starts with the given char
func (g *Generator) HeadChar(char string, ignore bool) {
	g.checkCharPrefix(char, 0, false)
}
