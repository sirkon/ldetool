package gogen

import "github.com/sirkon/ldetool/internal/generator/gogen/srcobj"

/* decode_int8 sample - other templates are the same
tmp ={{ .Source }};
if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, {{ .Bits }}); err != nil {
    return false, fmt.Errorf("Error parsing as `{{ Type }}` for field {{ .Dest }}: %s", err)
}
p.{{Dest}} = int8(tmpInt)
*/

func (g *Generator) prepNumeric() {
	g.regImport("", "strconv")
	g.regImport("", "fmt")
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

func (g *Generator) decode(src srcobj.Source, tmp, gotype, dest, decoder string, params ...srcobj.Source) {
	g.prepNumeric()

	var p []srcobj.Source
	if g.useString {
		//g.regVar("tmpChar", "byte")
		//g.body.Append(srcobj.Assign("tmpChar", src))
		//g.body.Append(srcobj.Raw("\n"))
		//src = srcobj.Raw("tmpChar")
		p = []srcobj.Source{src}
	} else {
		g.regImport("", "unsafe")
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
			Then: srcobj.ReturnError("Cannot parse `%s`: %s", srcobj.Stringify(src), srcobj.Raw("err")),
		},
	)
	g.body.Append(srcobj.LineAssign{
		Receiver: dest,
		Expr:     srcobj.NewCall(gotype, srcobj.Raw(tmp)),
	})
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
