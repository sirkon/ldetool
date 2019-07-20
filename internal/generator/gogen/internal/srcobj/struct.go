package srcobj

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/types"
)

// ///////////////// special hard to access type ///////////////////
type hardToAccessNameYouShouldNotUse string

func (h hardToAccessNameYouShouldNotUse) TypeString() string {
	return string(h)
}

// /////////////////////////////////////////////////////////////////

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
	gen       generator.Generator
}

// Struct creates Strct for external consumption
func Struct(useString bool, g generator.Generator) *Strct {
	return &Strct{
		useString: useString,
		gen:       g,
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

// AddDec128 adds a couple of two elements in a structure emulating uint128 type
func (s *Strct) AddDec128(name string) {
	res := Struct(s.useString, s.gen)
	res.addPrimitive("Lo", "uint64")
	res.addPrimitive("Hi", "uint64")
	s.fields = append(s.fields, FieldDef{
		Name: name,
		Type: res,
	})
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

// AddStr adds native string field
func (s *Strct) AddStr(name string) {
	s.addPrimitive(name, "string")
}

// AddSubstruct add substruct and returns it
func (s *Strct) AddSubstruct(name string) *Strct {
	res := Struct(s.useString, s.gen)
	res.addPrimitive("Valid", "bool")
	s.fields = append(s.fields, FieldDef{
		Name: name,
		Type: res,
	})
	return res
}

// AddCustomType add custom type
func (s *Strct) AddCustomType(name string, fieldType types.TypeRegistration) {
	switch v := fieldType.(type) {
	case types.ImportedType:
		s.addPrimitive(name, v.Name)
		if err := s.gen.RegImport(strings.Split(strings.TrimLeft(v.Name, "*"), ".")[0], strings.Trim(v.ImportPath, `"`)); err != nil {
			panic(err)
		}
	case types.LocalType:
		s.addPrimitive(name, v.Name)
	}
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
