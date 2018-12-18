package srcobj

import (
	"bytes"
	"fmt"
	"io"
)

/////////////////// special hard to access type ///////////////////
type hardToAccessNameYouShouldNotUse string

func (h hardToAccessNameYouShouldNotUse) TypeString() string {
	return string(h)
}

///////////////////////////////////////////////////////////////////

// FieldType represents LDE generated struct's field
type FieldType interface {
	TypeString() string
}

// FieldDef describes LDE generated field definition
type FieldDef struct {
	Name string
	Type FieldType
}

// Strct represents LDE generated struct
type Strct struct {
	useString bool
	fields    []FieldDef
}

// Struct creates Strct for external consumption
func Struct(useString bool) *Strct {
	return &Strct{
		useString: useString,
	}
}

// TypeString implementation to satisfy FieldType
func (s *Strct) TypeString() string {
	res := &bytes.Buffer{}
	res.WriteString("struct {\n")
	for _, field := range s.fields {
		res.WriteString(field.Name)
		res.WriteByte(' ')
		res.WriteString(field.Type.TypeString())
		res.WriteByte('\n')
	}
	res.WriteByte('}')
	return res.String()
}

// addPrimitive ...
func (s *Strct) addPrimitive(fieldName, fieldType string) {
	field := FieldDef{
		Name: fieldName,
		Type: hardToAccessNameYouShouldNotUse(fieldType),
	}
	s.fields = append(s.fields, field)
}

// AddInt8 adds int field
func (s *Strct) AddInt(name string) {
	s.addPrimitive(name, "int")
}

// AddInt8 adds int8 field
func (s *Strct) AddInt8(name string) {
	s.addPrimitive(name, "int8")
}

// AddInt16 adds int16 field
func (s *Strct) AddInt16(name string) {
	s.addPrimitive(name, "int16")
}

// AddInt32 adds int32 field
func (s *Strct) AddInt32(name string) {
	s.addPrimitive(name, "int32")
}

// AddInt64 adds int64 field
func (s *Strct) AddInt64(name string) {
	s.addPrimitive(name, "int64")
}

// AddUint adds uint field
func (s *Strct) AddUint(name string) {
	s.addPrimitive(name, "uint")
}

// AddUint8 adds uint8 field
func (s *Strct) AddUint8(name string) {
	s.addPrimitive(name, "uint8")
}

// AddUint16 adds uint16 field
func (s *Strct) AddUint16(name string) {
	s.addPrimitive(name, "uint16")
}

// AddUint32 adds uint32 field
func (s *Strct) AddUint32(name string) {
	s.addPrimitive(name, "uint32")
}

// AddUint64 adds uint64 field
func (s *Strct) AddUint64(name string) {
	s.addPrimitive(name, "uint64")
}

// AddHex add hex field
func (s *Strct) AddHex(name string) {
	s.addPrimitive(name, "uint")
}

// AddHex8 adds hex8 field
func (s *Strct) AddHex8(name string) {
	s.addPrimitive(name, "uint8")
}

// AddHex16 adds hex16 field
func (s *Strct) AddHex16(name string) {
	s.addPrimitive(name, "uint16")
}

// AddHex32 adds hex32 field
func (s *Strct) AddHex32(name string) {
	s.addPrimitive(name, "uint32")
}

// AddHex64 adds hex64 field
func (s *Strct) AddHex64(name string) {
	s.addPrimitive(name, "uint64")
}

// AddOct add oct field
func (s *Strct) AddOct(name string) {
	s.addPrimitive(name, "uint")
}

// AddOct8 adds oct8 field
func (s *Strct) AddOct8(name string) {
	s.addPrimitive(name, "uint8")
}

// AddOct16 adds oct16 field
func (s *Strct) AddOct16(name string) {
	s.addPrimitive(name, "uint16")
}

// AddOct32 adds oct32 field
func (s *Strct) AddOct32(name string) {
	s.addPrimitive(name, "uint32")
}

// AddOct64 adds oct64 field
func (s *Strct) AddOct64(name string) {
	s.addPrimitive(name, "uint64")
}

// AddFloat32 adds float32 field
func (s *Strct) AddFloat32(name string) {
	s.addPrimitive(name, "float32")
}

// AddFloat64 adds float64 field
func (s *Strct) AddFloat64(name string) {
	s.addPrimitive(name, "float64")
}

// AddString adds string field
func (s *Strct) AddString(name string) {
	s.addPrimitive(name, RightType(s.useString))
}

// AddSubstruct add substruct and returns it
func (s *Strct) AddSubstruct(name string) *Strct {
	res := Struct(s.useString)
	res.addPrimitive("Valid", "bool")
	s.fields = append(s.fields, FieldDef{
		Name: name,
		Type: res,
	})
	return res
}

// structType ...
type structType struct {
	name string
	s    *Strct
}

// Dump source implementation
func (s structType) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "// %s ...\n", s.name); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "type %s %s", s.name, s.s.TypeString()); err != nil {
		return err
	}
	return nil
}
