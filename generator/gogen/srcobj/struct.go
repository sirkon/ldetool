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

// Struct represents LDE generated struct
type Struct struct {
	fields []FieldDef
}

// TypeString implementation to satisfy FieldType
func (s *Struct) TypeString() string {
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
func (s *Struct) addPrimitive(fieldName, fieldType string) {
	field := FieldDef{
		Name: fieldName,
		Type: hardToAccessNameYouShouldNotUse(fieldType),
	}
	s.fields = append(s.fields, field)
}

// AddInt8 adds int8 field
func (s *Struct) AddInt8(name string) {
	s.addPrimitive(name, "int8")
}

// AddInt16 adds int16 field
func (s *Struct) AddInt16(name string) {
	s.addPrimitive(name, "int16")
}

// AddInt32 adds int32 field
func (s *Struct) AddInt32(name string) {
	s.addPrimitive(name, "int32")
}

// AddInt64 adds int64 field
func (s *Struct) AddInt64(name string) {
	s.addPrimitive(name, "int64")
}

// AddUint8 adds uint8 field
func (s *Struct) AddUint8(name string) {
	s.addPrimitive(name, "uint8")
}

// AddUint16 adds uint16 field
func (s *Struct) AddUint16(name string) {
	s.addPrimitive(name, "uint16")
}

// AddUint32 adds uint32 field
func (s *Struct) AddUint32(name string) {
	s.addPrimitive(name, "uint32")
}

// AddUint64 adds uint64 field
func (s *Struct) AddUint64(name string) {
	s.addPrimitive(name, "uint64")
}

// AddFloat32 adds float32 field
func (s *Struct) AddFloat32(name string) {
	s.addPrimitive(name, "float32")
}

// AddFloat64 adds float64 field
func (s *Struct) AddFloat64(name string) {
	s.addPrimitive(name, "float64")
}

func (s *Struct) AddString(name string) {
	s.addPrimitive(name, "[]byte")
}

// AddSubstruct add substruct and returns it
func (s *Struct) AddSubstruct(name string) *Struct {
	res := &Struct{}
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
	s    *Struct
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
