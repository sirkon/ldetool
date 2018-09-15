package gogen

import (
	"encoding/json"
	"fmt"

	"encoding/binary"

	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

func (g *Generator) shortPrefixCheck(unquoted, anchor string, offset int) srcobj.Source {
	if !g.useString {
		g.regImport("", "unsafe")
	}
	var mask uint64
	var byteMask = make([]byte, 8)
	for i := 0; i < len(unquoted); i++ {
		byteMask[i] = 255
	}
	if g.platformType == generator.LittleEndian {
		mask = binary.LittleEndian.Uint64(byteMask)
	} else {
		mask = binary.BigEndian.Uint64(byteMask)
	}
	tmp := make([]byte, 8)
	copy(tmp, unquoted)
	var prefix uint64
	if g.platformType == generator.LittleEndian {
		prefix = binary.LittleEndian.Uint64(tmp)
	} else {
		prefix = binary.BigEndian.Uint64(tmp)
	}
	var lengthCheck srcobj.Source
	if offset > 0 {
		lengthCheck = srcobj.OperatorGE(
			srcobj.OperatorSub(
				srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
				srcobj.Literal(offset),
			),
			srcobj.Literal(len(unquoted)),
		)
	} else {
		lengthCheck = srcobj.OperatorGE(
			srcobj.NewCall("len", srcobj.Raw(g.curRestVar())),
			srcobj.Literal(len(unquoted)),
		)
	}

	return srcobj.OperatorAnd(
		lengthCheck,
		srcobj.OperatorEq(
			srcobj.OperatorBitAnd(
				srcobj.Deref(
					srcobj.NewCall(
						"(*uint64)",
						srcobj.NewCall(
							"unsafe.Pointer",
							srcobj.Ref(
								srcobj.Index{
									Src:   g.rest(),
									Index: srcobj.Literal(offset),
								},
							),
						),
					),
				),
				srcobj.HexU64(mask),
			),
			srcobj.HexU64(prefix),
		),
	)
}

func (g *Generator) checkStringPrefix(anchor string, offset int, ignore bool) error {
	var unquoted string
	if err := json.Unmarshal([]byte(anchor), &unquoted); err != nil {
		return fmt.Errorf("cannot unqouote \033[1m%s\033[0m: %s", anchor, err)
	}

	body := g.body
	body.Append(srcobj.Raw("\n"))
	if offset > 0 {
		body.Append(
			srcobj.Comment(fmt.Sprintf("Checks if rest[%d:] starts with `%s` and pass it", offset, anchor)))
	} else {
		body.Append(srcobj.Comment(fmt.Sprintf("Checks if the rest starts with `%s` and pass it", anchor)))
	}

	var rest = g.rest()
	if offset != 0 {
		rest = srcobj.SliceFrom(rest, srcobj.Literal(offset))
	}

	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`",
			srcobj.Stringify(rest),
			srcobj.Raw(anchor),
		)
	}

	var shift srcobj.Source = srcobj.Literal(len(unquoted) + offset)
	var code srcobj.Source

	if len(unquoted) <= 8 && g.platformType != generator.Universal && !g.useString {
		code = g.shortPrefixCheck(unquoted, anchor, offset)
	} else {
		g.regRightVar(g.curRestVar())
		g.regRightPkg()
		constName := g.constNameFromContent(anchor)

		shift = srcobj.NewCall("len", srcobj.Raw(constName))
		if offset != 0 {
			shift = srcobj.OperatorAdd(shift, srcobj.Literal(offset))
		}

		code = srcobj.NewCall(srcobj.RightPkg(g.useString)+".HasPrefix", rest, srcobj.Raw(constName))
		if offset > 0 {
			code = srcobj.OperatorAnd(
				srcobj.OperatorGE(
					srcobj.NewCall("len", rest),
					srcobj.OperatorAdd(
						srcobj.Literal(offset),
						srcobj.NewCall("len", srcobj.Raw(constName)),
					),
				),
				code,
			)
		}
	}

	body.Append(srcobj.If{
		Expr: code,
		Then: srcobj.LineAssign{
			Receiver: g.curRestVar(),
			Expr:     srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), shift),
		},
		Else: failure,
	})
	return nil
}

// HeadString checks if the rest starts with the given string and passes it
func (g *Generator) HeadString(anchor string, ignore bool) error {
	return g.checkStringPrefix(anchor, 0, ignore)
}

func (g *Generator) checkCharPrefix(char string, offset int, ignore bool) error {
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}

	var rest srcobj.Source = srcobj.Raw(g.curRestVar())

	var shift srcobj.Source = srcobj.Literal(1)
	if offset != 0 {
		shift = srcobj.OperatorAdd(srcobj.Literal(offset), shift)
	}

	var failure srcobj.Source
	if !ignore {
		failure = g.failure(
			"`\033[1m%s\033[0m)` is expected to start with \033[1m%s\033[0m",
			srcobj.Stringify(rest),
			srcobj.DrawChar(char),
		)
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
		Then: srcobj.LineAssign{
			Receiver: g.curRestVar(),
			Expr:     srcobj.SliceFrom(srcobj.Raw(g.curRestVar()), shift),
		},
		Else: failure,
	})
	g.body.Append(body)
	return nil
}

// HeadChar checks if rest starts with the given char
func (g *Generator) HeadChar(char string, ignore bool) error {
	return g.checkCharPrefix(char, 0, false)
}
