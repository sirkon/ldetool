package gogen

import (
	"fmt"
	"github.com/sirkon/ldetool/internal/ast"
	"strings"

	"encoding/json"

	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

/* take_before_string
if pos = bytes.Index(p.Rest, {{ .ConstName }}); pos >= 0 {
    {{ call .Decoder "p.Rest[:pos]" .Dest }}
} else {
    return false, {{ if .Serious }}fmt.Errorf("Can't find a string limiting a value of {{ .Dest }}{{ else }}nil{{end}}
}
*/

// getterGen generates optional getter
func (g *Generator) getterGen(name, fieldType string) error {
	if len(g.ruleName) == 0 {
		return fmt.Errorf("Rule set up required")
	}
	if len(g.namespaces) == 0 {
		return nil
	}

	arg, err := srcobj.Go2ResultType(g.useString, fieldType)
	if err != nil {
		return err
	}
	method := srcobj.NewAccessor(
		g.ruleName,
		g.goish.Public("get_"+strings.Join(append(g.namespaces, name), "_")),
		arg,
	)
	g.optgetters.Append(method)
	body := method.Body()
	origBody := body
	for i := 1; i <= len(g.namespaces); i++ {
		valid := "p." + strings.Join(g.namespaces[:i], ".") + ".Valid"
		newBody := srcobj.NewBody()
		body.Append(srcobj.If{
			Expr: srcobj.Raw(valid),
			Then: newBody,
		})
		body = newBody
	}
	body.Append(srcobj.LineAssign{
		Receiver: "res",
		Expr:     srcobj.Raw("p." + strings.Join(append(g.namespaces, name), ".")),
	})
	origBody.Append(srcobj.Raw("return"))
	return nil
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

func (g *Generator) sliceTooLarge(upper int) srcobj.Source {
	return g.failure("cannot slice up to %d as only %d characters left in the rest (`\033[1m%s\033[0m`)",
		srcobj.Literal(upper),
		srcobj.NewCall("len", g.rest()),
		srcobj.Stringify(g.rest()))
}

func (g *Generator) jumpTooLarge(lower int) srcobj.Source {
	return g.failure(
		"cannot slice from %d as only %d characters left in the rest (`\033[1m%s\033[0m`)",
		srcobj.Literal(lower),
		srcobj.NewCall("len", g.rest()),
		srcobj.Stringify(g.rest()),
	)
}

// TakeBeforeStringEx ...
func (g *Generator) TakeBeforeString(name, fieldType, anchor string, meta ast.FieldMeta, lower, upper int, close, expand, include bool) error {
	if err := g.regVar("pos", "int"); err != nil {
		return err
	}
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}

	item := g.fields[g.fullName(name)]
	if err := g.getterGen(name, fieldType); err != nil {
		return err
	}
	g.curField = item
	g.curFieldType = fieldType

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
	ccc := " "
	if expand {
		ccc = " (or all the rest if not found) "
	}
	if lower > 0 && lower == upper {
		body.Append(srcobj.Comment(
			fmt.Sprintf(
				"Take until %s %s if it starts %s substring%sas %s(%s)",
				numerator(lower), valueIfTrue(include, "including it"), anchor, ccc, name, fieldType)))

		var unquoted string
		if err := json.Unmarshal([]byte(anchor), &unquoted); err != nil {
			return fmt.Errorf("cannot unqouote \033[1m%s\033[0m: %s", anchor, err)
		}
		var cond srcobj.Source
		if len(unquoted) <= 8 && g.platformType != generator.Universal && !g.useString {
			cond = g.shortPrefixCheck(unquoted, anchor, lower)
		} else {
			g.regRightPkg()
			cond = srcobj.OperatorAnd(
				srcobj.OperatorGE(
					srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
					srcobj.OperatorAdd(
						srcobj.NewCall("len", srcobj.Raw(constName)),
						srcobj.Literal(lower),
					),
				),
				srcobj.NewCall(
					srcobj.RightPkg(g.useString)+".HasPrefix",
					srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), srcobj.Literal(lower)),
					srcobj.Raw(constName)),
			)
		}
		body.Append(srcobj.If{
			Expr: cond,
			Then: srcobj.Assign("pos", srcobj.Literal(lower)),
			Else: srcobj.Assign("pos", srcobj.Literal(-1)),
		})
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%s%sas %s(%s)", anchor, valueIfTrue(include, " including it "), ccc, name, fieldType)))
		var lookup srcobj.Source
		if close {
			lookup = srcobj.LookupStringShort(g.useString, "pos", rest, srcobj.Raw(constName))
		} else {
			g.regRightPkg()
			var detector srcobj.Source = srcobj.LookupStringLong(g.useString, "pos", rest, srcobj.Raw(constName))
			lookup = srcobj.NewBody(srcobj.Trim(detector), srcobj.Raw("\n"))
		}

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
		body.Append(lookup)
	}

	var alternative srcobj.Source
	if !expand {
		alternative = g.failure(
			"cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m` to bound data for field "+name,
			srcobj.Raw(constName),
			srcobj.Stringify(rest),
		)
	} else {
		if fieldType == "string" {
			alternative = srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: g.varName(item.name),
					Expr:     g.rest(),
				},
				srcobj.LineAssign{
					Receiver: g.curRestVar(),
					Expr: srcobj.SliceFrom(
						srcobj.Raw(g.curRestVar()),
						srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
					),
				},
			)
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
	}

	var offset srcobj.Source
	offset = srcobj.OperatorAdd(
		srcobj.Raw("pos"),
		srcobj.NewCall("len", srcobj.Raw(constName)),
	)
	if lower > 0 && upper != lower {
		offset = srcobj.OperatorAdd(offset, srcobj.Literal(lower))
	}
	var takeOff srcobj.Source
	if include {
		takeOff = offset
	} else {
		takeOff = srcobj.Raw("pos")
		if lower > 0 && upper != lower {
			takeOff = srcobj.OperatorAdd(takeOff, srcobj.Literal(lower))
		}
	}

	var mainPath srcobj.Source
	if fieldType == "string" {
		mainPath = srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "p." + item.name,
				Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), takeOff),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					offset,
				),
			},
		)
	} else {
		mainPath = srcobj.NewBody(
			srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: "tmp",
					Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), takeOff),
				},
				srcobj.LineAssign{
					Receiver: g.curRestVar(),
					Expr: srcobj.SliceFrom(
						srcobj.Raw(g.curRestVar()),
						offset,
					),
				},
			),
		)
	}
	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: mainPath,
		Else: alternative,
	})

	if fieldType != "string" {
		if err := g.regRightVar("tmp"); err != nil {
			return err
		}
		if decoder, ok := g.decoderMap[fieldType]; ok {
			decoder(srcobj.Raw("tmp"), "p."+item.name)
		}
		if decoder, ok := g.decimalDecoderMap[fieldType]; ok {
			decoder(srcobj.Raw("tmp"), "p."+item.name, meta.Precision, meta.Scale)
		}
	}
	return nil
}

func valueIfTrue(flag bool, value string) string {
	if flag {
		return value
	}
	return ""
}

// TakeBeforeChar ...
func (g *Generator) TakeBeforeChar(name, fieldType, char string, meta ast.FieldMeta, lower, upper int, close, expand, include bool) error {
	if err := g.regVar("pos", "int"); err != nil {
		return err
	}
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}

	item := g.fields[g.fullName(name)]
	if err := g.getterGen(name, fieldType); err != nil {
		return err
	}
	g.curField = item
	g.curFieldType = fieldType

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
	ccc := " "
	if expand {
		ccc = " (or all the rest if not found) "
	}
	if lower > 0 && lower == upper {
		body.Append(srcobj.Comment(
			fmt.Sprintf(
				"Take until %s character %sif it is%s%sas %s(%s)",
				numerator(lower), valueIfTrue(include, "including it "), char, ccc, name, fieldType)))
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
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Take until %s%s%sas %s(%s)", char, valueIfTrue(include, " including it"), ccc, name, fieldType)))
		var lookup srcobj.Source
		if close {
			lookup = srcobj.LookupByteShort{
				Var:    "pos",
				Src:    rest,
				Needle: srcobj.Raw(char),
			}
		} else {
			g.regRightPkg()
			var detector srcobj.Source = srcobj.LookupByteLong(g.useString, "pos", rest, srcobj.Raw(char))
			lookup = srcobj.NewBody(srcobj.Trim(detector), srcobj.Raw("\n"))
		}

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
		body.Append(lookup)
	}
	var alternative srcobj.Source
	if !expand {
		alternative = g.failure(
			"cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field "+name,
			srcobj.Raw(char),
			srcobj.Stringify(rest),
		)
	} else {
		if fieldType == "string" {
			alternative = srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: "p." + item.name,
					Expr:     g.rest(),
				},
				srcobj.LineAssign{
					Receiver: g.curRestVar(),
					Expr: srcobj.SliceFrom(
						srcobj.Raw(g.curRestVar()),
						srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
					),
				},
			)
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
	}

	var offset srcobj.Source
	offset = srcobj.OperatorAdd(
		srcobj.Raw("pos"),
		srcobj.Literal(1),
	)
	if lower > 0 && upper != lower {
		offset = srcobj.OperatorAdd(offset, srcobj.Literal(lower))
	}

	var takeOff srcobj.Source
	if include {
		takeOff = offset
	} else {
		takeOff = srcobj.Raw("pos")
		if lower > 0 && upper != lower {
			takeOff = srcobj.OperatorAdd(takeOff, srcobj.Literal(lower))
		}
	}

	var mainPath srcobj.Source
	if fieldType == "string" {
		mainPath = srcobj.NewBody(
			srcobj.LineAssign{
				Receiver: "p." + item.name,
				Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), takeOff),
			},
			srcobj.LineAssign{
				Receiver: g.curRestVar(),
				Expr: srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					offset,
				),
			},
		)
	} else {
		mainPath = srcobj.NewBody(
			srcobj.NewBody(
				srcobj.LineAssign{
					Receiver: "tmp",
					Expr:     srcobj.SliceTo(srcobj.Raw(g.curRestVar()), takeOff),
				},
				srcobj.LineAssign{
					Receiver: g.curRestVar(),
					Expr: srcobj.SliceFrom(
						srcobj.Raw(g.curRestVar()),
						offset,
					),
				},
			),
		)
	}
	body.Append(srcobj.If{
		Expr: srcobj.OperatorGE(srcobj.Raw("pos"), srcobj.Raw("0")),
		Then: mainPath,
		Else: alternative,
	})

	if fieldType != "string" {
		if err := g.regRightVar("tmp"); err != nil {
			return err
		}
		if decoder, ok := g.decoderMap[fieldType]; ok {
			decoder(srcobj.Raw("tmp"), "p."+item.name)
		}
		if decoder, ok := g.decimalDecoderMap[fieldType]; ok {
			decoder(srcobj.Raw("tmp"), "p."+item.name, meta.Precision, meta.Scale)
		}
	}
	return nil
}

// TakeRest ...
func (g *Generator) TakeRest(name, fieldType string, meta ast.FieldMeta) error {
	item := g.fields[g.fullName(name)]
	if err := g.getterGen(name, fieldType); err != nil {
		return err
	}

	body := g.indent()
	body.Append(srcobj.Comment(fmt.Sprintf("Take the rest as %s(%s)", name, fieldType)))

	if fieldType == "string" {
		body.Append(
			srcobj.Assign(
				g.varName(item.name),
				g.rest(),
			),
		)
		body.Append(srcobj.Raw("\n"))
		body.Append(
			srcobj.Assign(
				g.curRestVar(),
				srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					srcobj.NewCall(
						"len",
						srcobj.Raw(g.curRestVar()),
					),
				),
			),
		)
	} else {
		if decoder, ok := g.decoderMap[fieldType]; ok {
			decoder(g.rest(), "p."+item.name)
		}
		if decoder, ok := g.decimalDecoderMap[fieldType]; ok {
			decoder(g.rest(), "p."+item.name, meta.Precision, meta.Scale)
		}
		body.Append(
			srcobj.Assign(
				g.curRestVar(),
				srcobj.SliceFrom(
					srcobj.Raw(g.curRestVar()),
					srcobj.NewCall(
						"len",
						srcobj.Raw(g.curRestVar()),
					),
				),
			),
		)
	}
	return nil
}
