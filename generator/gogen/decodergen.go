package gogen

import (
	"bytes"

	"fmt"

	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

// DecoderGen generates code for value decoding
type DecoderGen struct {
	g *Generator
}

/* decode_int8 sample - other templates are the same
tmp ={{ .Source }};
if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, {{ .Bits }}); err != nil {
    return false, fmt.Errorf("Error parsing as `{{ Type }}` for field {{ .Dest }}: %s", err)
}
p.{{Dest}} = int8(tmpInt)
*/

// Int8 decoder
func (dg *DecoderGen) Int8(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpInt", "int64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_int", buf, DParams{
		Source: src,
		Type:   "int8",
		Dest:   dest,
		Bits:   8,
	})
	return buf.String()
}

// Int16 decoder
func (dg *DecoderGen) Int16(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpInt", "int64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_int", buf, DParams{
		Source: src,
		Type:   "int16",
		Dest:   dest,
		Bits:   16,
	})
	return buf.String()
}

// Int32 decoder
func (dg *DecoderGen) Int32(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpInt", "int64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_int", buf, DParams{
		Source: src,
		Type:   "int32",
		Dest:   dest,
		Bits:   32,
	})
	return buf.String()
}

// Int64 decoder
func (dg *DecoderGen) Int64(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpInt", "int64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_int", buf, DParams{
		Source: src,
		Type:   "int64",
		Dest:   dest,
		Bits:   64,
	})
	return buf.String()
}

// Uint8 decoder
func (dg *DecoderGen) Uint8(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpUint", "uint64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_uint", buf, DParams{
		Source: src,
		Type:   "uint8",
		Dest:   dest,
		Bits:   8,
	})
	return buf.String()
}

// Uint16 decoder
func (dg *DecoderGen) Uint16(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpUint", "uint64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_uint", buf, DParams{
		Source: src,
		Type:   "uint16",
		Dest:   dest,
		Bits:   16,
	})
	return buf.String()
}

// Uint32 decoder
func (dg *DecoderGen) Uint32(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpUint", "uint64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_uint", buf, DParams{
		Source: src,
		Type:   "uint32",
		Dest:   dest,
		Bits:   32,
	})
	return buf.String()
}

// Uint64 decoder
func (dg *DecoderGen) Uint64(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpUint", "uint64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_uint", buf, DParams{
		Source: src,
		Type:   "uint64",
		Dest:   dest,
		Bits:   64,
	})
	return buf.String()
}

// Float32 decoder
func (dg *DecoderGen) Float32(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpFloat", "float64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_float", buf, DParams{
		Source: src,
		Type:   "float32",
		Dest:   dest,
		Bits:   32,
	})
	return buf.String()
}

// Float64 decoder
func (dg *DecoderGen) Float64(src, dest string) string {
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "fmt")
	dg.g.regVar("tmpFloat", "float64")
	dg.g.regVar("err", "error")

	buf := &bytes.Buffer{}
	dg.g.tc.MustExecute("decode_float", buf, DParams{
		Source: src,
		Type:   "float64",
		Dest:   dest,
		Bits:   64,
	})
	return buf.String()
}

// String decoder
func (dg *DecoderGen) String(src, dest string) string {
	buf := &bytes.Buffer{}
	dg.g.TmpOff()
	dg.g.tc.MustExecute("decode_string", buf, DParams{
		Source: src,
		Dest:   dest,
	})
	return buf.String()
}

// Source decoder
func (dg *DecoderGen) Source(dest string, src srcobj.Source, fieldType string) (res srcobj.Source) {
	var unknownType bool
	defer func() {
		if unknownType {
			panic(fmt.Errorf("unsupported type %s for field %s", fieldType, dest))
		}
	}()
	failure := srcobj.ReturnError(
		"error parsing `\033[1m%s\033]0m` value as "+fieldType+" for field `\033[1m"+dest[2:]+"\033[0m`: %s",
		src, srcobj.Raw("err"))

	if fieldType == "string" {
		return srcobj.LineAssign{Receiver: dest, Expr: src}
	}

	dg.g.regVar("err", "error")
	dg.g.regVar("tmp", "[]byte")
	dg.g.regImport("", "unsafe")
	dg.g.regImport("", "strconv")
	dg.g.regImport("", "fmt")
	var decoder srcobj.Call
	unsafeDeref := srcobj.NewCall("*(*string)", srcobj.NewCall("unsafe.Pointer", srcobj.Ref(srcobj.Raw("tmp"))))

	var tmpDest string
	var conv string
	switch fieldType {
	case "int8":
		dg.g.regVar("tmpInt", "int64")
		decoder = srcobj.NewCall("strconv.ParseInt", unsafeDeref, srcobj.Literal(10), srcobj.Literal(8))
		tmpDest = "tmpInt"
		conv = "int8"
	case "int16":
		dg.g.regVar("tmpInt", "int64")
		decoder = srcobj.NewCall("strconv.ParseInt", unsafeDeref, srcobj.Literal(10), srcobj.Literal(16))
		tmpDest = "tmpInt"
		conv = "int16"
	case "int32":
		dg.g.regVar("tmpInt", "int64")
		decoder = srcobj.NewCall("strconv.ParseInt", unsafeDeref, srcobj.Literal(10), srcobj.Literal(32))
		tmpDest = "tmpInt"
		conv = "int32"
	case "int64":
		dg.g.regVar("tmpInt", "int64")
		decoder = srcobj.NewCall("strconv.ParseInt", unsafeDeref, srcobj.Literal(10), srcobj.Literal(64))
		tmpDest = "tmpInt"
		conv = "int64"
	case "uint8":
		dg.g.regVar("tmpUint", "uint64")
		decoder = srcobj.NewCall("strconv.ParseUint", unsafeDeref, srcobj.Literal(10), srcobj.Literal(8))
		tmpDest = "tmpUint"
		conv = "uint8"
	case "uint16":
		dg.g.regVar("tmpUint", "uint64")
		decoder = srcobj.NewCall("strconv.ParseUint", unsafeDeref, srcobj.Literal(10), srcobj.Literal(16))
		tmpDest = "tmpUint"
		conv = "uint16"
	case "uint32":
		dg.g.regVar("tmpUint", "uint64")
		decoder = srcobj.NewCall("strconv.ParseUint", unsafeDeref, srcobj.Literal(10), srcobj.Literal(32))
		tmpDest = "tmpUint"
		conv = "uint32"
	case "uint64":
		dg.g.regVar("tmpUint", "uint64")
		decoder = srcobj.NewCall("strconv.ParseUint", unsafeDeref, srcobj.Literal(10), srcobj.Literal(64))
		tmpDest = "tmpUint"
		conv = "uint64"
	case "float32":
		dg.g.regVar("tmpFloat", "float64")
		decoder = srcobj.NewCall("strconv.ParseFloat", unsafeDeref, srcobj.Literal(32))
		tmpDest = "tmpFloat"
		conv = "float32"
	case "float64":
		dg.g.regVar("tmpFloat", "float64")
		decoder = srcobj.NewCall("strconv.ParseFloat", unsafeDeref, srcobj.Literal(64))
		tmpDest = "tmpFloat"
		conv = "float64"
	default:
		unknownType = true
		return
	}

	return srcobj.NewBody(
		srcobj.Decode(tmpDest, decoder, failure),
		srcobj.LineAssign{
			dest,
			srcobj.NewCall(conv, srcobj.Raw(tmpDest)),
		},
	)
}
