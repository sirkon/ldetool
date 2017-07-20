package gogen

import "bytes"

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
	dg.g.regVar("tmpUint", "uint64")

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
	dg.g.regVar("tmpUint", "float64")

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
	dg.g.regVar("tmpUint", "float64")

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
