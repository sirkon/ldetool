package gogen

import (
	"fmt"
	"strings"

	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

/* take_before_string
if pos = bytes.Index(p.rest, {{ .ConstName }}); pos >= 0 {
    {{ call .Decoder "p.rest[:pos]" .Dest }}
} else {
    return false, {{ if .Serious }}fmt.Errorf("Can't find a string limiting a value of {{ .Dest }}{{ else }}nil{{end}}
}
*/

// getterGen generates optional getter
func (g *Generator) getterGen(name, fieldType string) {
	if len(g.parserName) == 0 {
		panic(fmt.Errorf("Rule set up required"))
	}
	if len(g.namespaces) == 0 {
		return
	}
	goType := g.goType(fieldType)
	accesses := []string{}
	for i := 1; i <= len(g.namespaces); i++ {
		accesses = append(accesses, strings.Join(g.namespaces[:i], "."))
	}

	g.tc.MustExecute("getter", g.opgetters, GParams{
		Accesses:   accesses,
		Access:     strings.Join(g.namespaces, "."),
		LongName:   g.goish.Public(strings.Join(append(g.namespaces, name), "_")),
		ParserName: g.parserName,
		Name:       name,
		Type:       goType,
	})
}

func numerator(num int) string {
	switch num {
	case 0:
		return "1st"
	case 1:
		return "2nd"
	case 2:
		return "3rd"
	default:
		return fmt.Sprintf("%dth", num+1)
	}
}

// TakeBeforeStringEx ...
func (g *Generator) TakeBeforeString(name, fieldType, anchor string, lower, upper int, expand bool) {
	g.regVar("pos", "int")
	g.regVar(g.curRestVar(), "[]byte")
	g.regVar("tmp", "[]byte")

	item := g.fields[g.fullName(name)]
	g.getterGen(name, fieldType)

	constName := g.constNameFromContent(anchor)

	var rest srcobj.Source
	if upper > 0 {
		u := fmt.Sprintf("%d", upper)
		if lower > 0 {
			l := fmt.Sprintf("%d", lower)
			rest = srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Raw(l), srcobj.Raw(u))
		} else {
			rest = srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Raw(u))
		}
	} else {
		rest = srcobj.Raw(g.curRestVar())
	}

	body := srcobj.NewBody(srcobj.Raw("\n"))
	ccc := " "
	if expand {
		ccc = " (or all the rest if not found) "
	}
	if lower > 0 && lower == upper {
		body.Append(srcobj.Comment(
			fmt.Sprintf(
				"Take until %s character if it starts %s substring%sas %s(%s)",
				numerator(lower), anchor, ccc, name, fieldType)))
		g.regImport("", "bytes")
		cond := srcobj.OperatorAnd(
			srcobj.OperatorGE(
				srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
				srcobj.OperatorAdd(
					srcobj.NewCall("len", srcobj.Raw(constName)),
					srcobj.Literal(lower),
				),
			),
			srcobj.NewCall(
				"bytes.HasPrefix",
				srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower)),
				srcobj.Raw(constName)),
		)
		body.Append(srcobj.If{
			Expr: cond,
			Then: srcobj.Assign("pos", srcobj.Literal(lower)),
			Else: srcobj.Assign("pos", srcobj.Literal(-1)),
		})
	} else if lower >= 0 {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%sas %s(%s)", anchor, ccc, name, fieldType)))
		g.regImport("", "bytes")
		var detector srcobj.Source = srcobj.LookupStringLong{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(constName),
		}
		if lower > 0 {
			detector = srcobj.OperatorAdd(detector, srcobj.Literal(lower))
		}
		body.Append(srcobj.Trim(detector))
		body.Append(srcobj.Raw("\n"))
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%sas %s(%s)", anchor, ccc, name, fieldType)))
		body.Append(srcobj.LookupStringShort{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(constName),
		})
	}
	validPath := "p." + strings.Join(g.namespaces, ".") + ".Valid"
	var alternative srcobj.Source
	if !expand {
		if len(g.namespaces) > 0 {
			g.abandon()
			alternative = srcobj.Goto(g.goish.Private(strings.Join(g.namespaces, "_") + "_label"))
		} else if g.serious {
			g.regImport("", "fmt")
			alternative = srcobj.ReturnError(
				"Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m` to bound data for field "+name,
				srcobj.Raw(constName),
				rest,
			)
		} else {
			alternative = srcobj.ReturnFail
		}
		if len(g.namespaces) > 0 {
			alternative = srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: validPath,
					Expr:     srcobj.Raw("false"),
				},
				alternative,
			)
		}
	} else {
		alternative = srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "tmp",
				Expr:     srcobj.Raw(g.curRestVar()),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
				),
			},
		)
	}

	var offset srcobj.Source
	offset = srcobj.OperatorAdd(
		srcobj.Raw("pos"),
		srcobj.NewCall("len", srcobj.Raw(constName)),
	)

	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "tmp",
				Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Raw("pos")),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					offset,
				),
			},
		),
		Else: alternative,
	})

	body.Append(g.dgen.Source("p."+item.name, srcobj.Raw("tmp"), fieldType))
	if len(g.namespaces) > 0 {
		body.Append(
			srcobj.LineAssign{
				Receiver: validPath,
				Expr:     srcobj.Raw("true"),
			},
		)
	}
	if err := body.Dump(g.curBody); err != nil {
		panic(err)
	}
}

// TakeBeforeChar ...
func (g *Generator) TakeBeforeChar(name, fieldType, char string, lower, upper int, expand bool) {
	g.regVar("pos", "int")
	g.regVar(g.curRestVar(), "[]byte")
	g.regVar("tmp", "[]byte")

	item := g.fields[g.fullName(name)]
	g.getterGen(name, fieldType)

	var rest srcobj.Source
	if upper > 0 {
		u := fmt.Sprintf("%d", upper)
		if lower > 0 {
			l := fmt.Sprintf("%d", lower)
			rest = srcobj.Slice(srcobj.Raw(g.curRestVar()), srcobj.Raw(l), srcobj.Raw(u))
		} else {
			rest = srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Raw(u))
		}
	} else {
		rest = srcobj.Raw(g.curRestVar())
	}

	body := srcobj.NewBody(srcobj.Raw("\n"))
	ccc := " "
	if expand {
		ccc = " (or all the rest if not found) "
	}
	if lower > 0 && lower == upper {
		body.Append(srcobj.Comment(
			fmt.Sprintf(
				"Take until %s character if it is%s%sas %s(%s)",
				numerator(lower), char, ccc, name, fieldType)))
		cond := srcobj.OperatorAnd(
			srcobj.OperatorGE(
				srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
				srcobj.OperatorAdd(
					srcobj.Literal(lower),
					srcobj.Literal(1),
				),
			),
			srcobj.OperatorEq(
				srcobj.Index{
					Src:   srcobj.Raw(g.curRestVar()),
					Index: srcobj.Literal(lower),
				},
				srcobj.Raw(char),
			),
		)
		body.Append(srcobj.If{
			Expr: cond,
			Then: srcobj.Assign("pos", srcobj.Literal(lower)),
			Else: srcobj.Assign("pos", srcobj.Literal(-1)),
		})
	} else if lower >= 0 {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%sas %s(%s)", char, ccc, name, fieldType)))
		g.regImport("", "bytes")
		var detector srcobj.Source = srcobj.LookupByteLong{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(char),
		}
		if lower > 0 {
			detector = srcobj.OperatorAdd(detector, srcobj.Literal(lower))
		}
		body.Append(srcobj.Trim(detector))
		body.Append(srcobj.Raw("\n"))
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%sas %s(%s)", char, ccc, name, fieldType)))
		body.Append(srcobj.LookupByteShort{
			Var:    "pos",
			Src:    rest,
			Needle: srcobj.Raw(char),
		})
	}
	validPath := "p." + strings.Join(g.namespaces, ".") + ".Valid"
	var alternative srcobj.Source
	if !expand {
		if len(g.namespaces) > 0 {
			g.abandon()
			alternative = srcobj.Goto(g.goish.Private(strings.Join(g.namespaces, "_") + "_label"))
		} else if g.serious {
			g.regImport("", "fmt")
			alternative = srcobj.ReturnError(
				"Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m` to bound data for field "+name,
				srcobj.Raw(char),
				rest,
			)
		} else {
			alternative = srcobj.ReturnFail
		}
		if len(g.namespaces) > 0 {
			alternative = srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: validPath,
					Expr:     srcobj.Raw("false"),
				},
				alternative,
			)
		}
	} else {
		alternative = srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "tmp",
				Expr:     srcobj.Raw(g.curRestVar()),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
				),
			},
		)
	}

	var offset srcobj.Source
	offset = srcobj.OperatorAdd(
		srcobj.Raw("pos"),
		srcobj.Literal(1),
	)

	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "tmp",
				Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), srcobj.Raw("pos")),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					offset,
				),
			},
		),
		Else: alternative,
	})

	body.Append(g.dgen.Source("p."+item.name, srcobj.Raw("tmp"), fieldType))

	if err := body.Dump(g.curBody); err != nil {
		panic(err)
	}
}

// TakeRest ...
func (g *Generator) TakeRest(name, fieldType string) {
	g.regVar(g.curRestVar(), "[]byte")
	g.regImport("", "bytes")

	item := g.fields[g.fullName(name)]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_rest", g.curBody, TParams{
		Rest:       g.curRestVar(),
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}
