package gogen

import (
	"fmt"
	"strings"

	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

/* decode_int8 sample - other templates are the same
tmp ={{ .Source }};
if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, {{ .Bits }}); err != nil {
    return false, fmt.Errorf("Error parsing as `{{ Type }}` for field {{ .Dest }}: %s", err)
}
p.{{Dest}} = int8(tmpInt)
*/

func (g *Generator) prepNumeric() {
	g.RegImport("", "strconv")
	g.RegImport("", "fmt")
	g.regVar("err", "error")
}

func (g *Generator) prepInt() {
	g.prepNumeric()
	g.regVar("tmpInt", "int64")
}

func (g *Generator) prepUint() {
	g.prepNumeric()
	g.regVar("tmpUint", "uint64")
}

func (g *Generator) prepFloat() {
	g.prepNumeric()
	g.regVar("tmpFloat", "float64")
}

func (g *Generator) returnError(srcName srcobj.Source) srcobj.Source {
	if g.silenceDepth > 0 {
		g.abandon()
		var pre srcobj.Source
		if g.anonymous() {
			pre = srcobj.Raw("")
		} else {
			pre = srcobj.Assign(g.valid(), srcobj.False)
		}
		return srcobj.NewBody(
			pre,
			srcobj.Semicolon,
			srcobj.Goto(g.label()),
		)
	}
	if !g.useString {
		srcName = srcobj.NewCall("string", srcName)
	}
	if err := g.RegImport("", "fmt"); err != nil {
		panic("failed to register fmt import")
	}
	return srcobj.ReturnError("parsing `%s` into field "+g.curField.name+"("+strings.TrimLeft(g.curFieldType, "$")+"): %s",
		srcName,
		srcobj.Raw("err"),
	)
}

func (g *Generator) decode(src srcobj.Source, tmp, gotype, dest, decoder string, params ...srcobj.Source) {
	g.prepNumeric()

	var p []srcobj.Source
	if g.useString {
		// g.regVar("tmpChar", "byte")
		// g.body.Append(srcobj.Assign("tmpChar", src))
		// g.body.Append(srcobj.Raw("\n"))
		// src = srcobj.Raw("tmpChar")
		p = []srcobj.Source{src}
	} else {
		g.RegImport("", "unsafe")
		p = []srcobj.Source{
			srcobj.Deref(
				srcobj.NewCall(
					"(*string)",
					srcobj.NewCall(
						"unsafe.Pointer",
						srcobj.Ref(src),
					),
				),
			),
		}
	}
	p = append(p, params...)
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorSemicolon(
				srcobj.OperatorAssign(
					srcobj.OperatorComma(
						srcobj.Raw(tmp),
						srcobj.Raw("err"),
					),
					srcobj.NewCall(
						decoder, // func name, strconv.ParseInt, for instance
						p...,
					),
				),
				srcobj.OperatorNEq(srcobj.Raw("err"), srcobj.Raw("nil")),
			),
			Then: g.returnError(p[0]),
		},
	)
	g.body.Append(srcobj.LineAssign{
		Receiver: dest,
		Expr:     srcobj.NewCall(gotype, srcobj.Raw(tmp)),
	})
}

func (g *Generator) decodeDirect(src srcobj.Source, dest, decoder string, params ...srcobj.Source) {
	if err := g.regErr(); err != nil {
		panic("duplicate err declaration")
	}
	var p []srcobj.Source
	p = append(p, src)
	p = append(p, params...)
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorSemicolon(
				srcobj.OperatorAssign(
					srcobj.OperatorComma(srcobj.Raw(dest), srcobj.Raw("err")),
					srcobj.NewCall(
						decoder, // func name, strconv.ParseInt, for instance
						p...,
					),
				),
				srcobj.OperatorNEq(srcobj.Raw("err"), srcobj.Raw("nil")),
			),
			Then: g.returnError(p[0]),
		},
	)
}

func (g *Generator) decodeInt(src srcobj.Source, dest string) {
	g.prepInt()
	g.decode(src, "tmpInt", "int", dest, "strconv.ParseInt", srcobj.Literal(10), srcobj.Literal(64))
}

func (g *Generator) decodeInt8(src srcobj.Source, dest string) {
	g.prepInt()
	g.decode(src, "tmpInt", "int8", dest, "strconv.ParseInt", srcobj.Literal(10), srcobj.Literal(8))
}

func (g *Generator) decodeInt16(src srcobj.Source, dest string) {
	g.prepInt()
	g.decode(src, "tmpInt", "int16", dest, "strconv.ParseInt", srcobj.Literal(10), srcobj.Literal(16))
}

func (g *Generator) decodeInt32(src srcobj.Source, dest string) {
	g.prepInt()
	g.decode(src, "tmpInt", "int32", dest, "strconv.ParseInt", srcobj.Literal(10), srcobj.Literal(32))
}

func (g *Generator) decodeInt64(src srcobj.Source, dest string) {
	g.prepInt()
	g.decode(src, "tmpInt", "int64", dest, "strconv.ParseInt", srcobj.Literal(10), srcobj.Literal(64))
}

func (g *Generator) decodeUint(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint", dest, "strconv.ParseUint", srcobj.Literal(10), srcobj.Literal(64))
}

func (g *Generator) decodeUint8(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint8", dest, "strconv.ParseUint", srcobj.Literal(10), srcobj.Literal(8))
}

func (g *Generator) decodeUint16(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint16", dest, "strconv.ParseUint", srcobj.Literal(10), srcobj.Literal(16))
}

func (g *Generator) decodeUint32(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint32", dest, "strconv.ParseUint", srcobj.Literal(10), srcobj.Literal(32))
}

func (g *Generator) decodeUint64(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint64", dest, "strconv.ParseUint", srcobj.Literal(10), srcobj.Literal(64))
}

func (g *Generator) decodeHex(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint", dest, "strconv.ParseUint", srcobj.Literal(16), srcobj.Literal(64))
}

func (g *Generator) decodeHex8(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint8", dest, "strconv.ParseUint", srcobj.Literal(16), srcobj.Literal(8))
}

func (g *Generator) decodeHex16(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint16", dest, "strconv.ParseUint", srcobj.Literal(16), srcobj.Literal(16))
}

func (g *Generator) decodeHex32(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint32", dest, "strconv.ParseUint", srcobj.Literal(16), srcobj.Literal(32))
}

func (g *Generator) decodeHex64(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint64", dest, "strconv.ParseUint", srcobj.Literal(16), srcobj.Literal(64))
}

func (g *Generator) decodeOct(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint", dest, "strconv.ParseUint", srcobj.Literal(8), srcobj.Literal(64))
}

func (g *Generator) decodeOct8(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint8", dest, "strconv.ParseUint", srcobj.Literal(8), srcobj.Literal(8))
}

func (g *Generator) decodeOct16(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint16", dest, "strconv.ParseUint", srcobj.Literal(8), srcobj.Literal(8))
}

func (g *Generator) decodeOct32(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint32", dest, "strconv.ParseUint", srcobj.Literal(8), srcobj.Literal(32))
}

func (g *Generator) decodeOct64(src srcobj.Source, dest string) {
	g.prepUint()
	g.decode(src, "tmpUint", "uint64", dest, "strconv.ParseUint", srcobj.Literal(8), srcobj.Literal(64))
}

func (g *Generator) prepDec() {
	g.regVar("err", "error")
	g.RegImport("", "fmt")
	g.RegImport("", "github.com/sirkon/decconv")
}

func (g *Generator) decodeSmallDecimal(src srcobj.Source, dest, decoder string, precision, scale int) {
	g.RegImport("", "fmt")
	p := []srcobj.Source{
		srcobj.Literal(precision),
		srcobj.Literal(scale),
	}
	if g.useString {
		p = append(p, srcobj.NewCall("[]byte", src))
	} else {
		p = append(p, src)
	}
	var printSrc srcobj.Source
	if g.useString {
		printSrc = src
	} else {
		printSrc = srcobj.NewCall("string", src)
	}
	g.curFieldType = fmt.Sprintf("dec%d.%d", precision, scale)
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorSemicolon(
				srcobj.OperatorAssign(
					srcobj.OperatorComma(
						srcobj.Raw(dest),
						srcobj.Raw("err"),
					),
					srcobj.NewCall(
						decoder, // func name, strconv.ParseInt, for instance
						p...,
					),
				),
				srcobj.OperatorNEq(srcobj.Raw("err"), srcobj.Raw("nil")),
			),
			Then: g.returnError(printSrc),
		},
	)
}

func (g *Generator) decodeDec32(src srcobj.Source, dest string, precision, scale int) {
	g.decodeSmallDecimal(src, dest, "decconv.Decode32", precision, scale)
}

func (g *Generator) decodeDec64(src srcobj.Source, dest string, precision, scale int) {
	g.decodeSmallDecimal(src, dest, "decconv.Decode64", precision, scale)
}

func (g *Generator) decodeDec128(src srcobj.Source, dest string, precision, scale int) {
	g.prepDec()
	p := []srcobj.Source{
		srcobj.Literal(precision),
		srcobj.Literal(scale),
	}
	if g.useString {
		p = append(p, srcobj.NewCall("[]byte", src))
	} else {
		p = append(p, src)
	}
	var printSrc srcobj.Source
	if g.useString {
		printSrc = src
	} else {
		printSrc = srcobj.NewCall("string", src)
	}
	g.curFieldType = fmt.Sprintf("dec%d.%d", precision, scale)
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorSemicolon(
				srcobj.OperatorAssign(
					srcobj.OperatorComma(
						srcobj.OperatorComma(
							srcobj.Raw(dest+".Lo"),
							srcobj.Raw(dest+".Hi"),
						),
						srcobj.Raw("err"),
					),
					srcobj.NewCall(
						"decconv.Decode128", // func name, strconv.ParseInt, for instance
						p...,
					),
				),
				srcobj.OperatorNEq(srcobj.Raw("err"), srcobj.Raw("nil")),
			),
			Then: g.returnError(printSrc),
		},
	)
}

func (g *Generator) decodeFloat32(src srcobj.Source, dest string) {
	g.prepFloat()
	g.decode(src, "tmpFloat", "float32", dest, "strconv.ParseFloat", srcobj.Literal(32))
}

func (g *Generator) decodeFloat64(src srcobj.Source, dest string) {
	g.prepFloat()
	g.decode(src, "tmpFloat", "float64", dest, "strconv.ParseFloat", srcobj.Literal(64))
}

func (g *Generator) decodeString(src srcobj.Source, dest string) {
	g.body.Append(srcobj.LineAssign{
		Receiver: dest,
		Expr:     src,
	})
}
func (g *Generator) decodeStr(src srcobj.Source, dest string) {
	if g.useString {
		g.body.Append(srcobj.LineAssign{
			Receiver: dest,
			Expr:     src,
		})
	} else {
		g.body.Append(srcobj.LineAssign{
			Receiver: dest,
			Expr:     srcobj.NewCall("string", src),
		})
	}
}

func (g *Generator) decodeCustomType(src srcobj.Source, dest string, decName string) {
	g.decodeDirect(src, dest, "p."+g.goish.Private("unmarshal"+decName))
}
