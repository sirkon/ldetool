package types

import (
	"io"
)

// FieldRegistrator abstration for field of given type adder
type FieldRegistrator interface {
	AddInt(comment []string, name string)
	AddInt8(comment []string, name string)
	AddInt16(comment []string, name string)
	AddInt32(comment []string, name string)
	AddInt64(comment []string, name string)
	AddUint(comment []string, name string)
	AddUint8(comment []string, name string)
	AddUint16(comment []string, name string)
	AddUint32(comment []string, name string)
	AddUint64(comment []string, name string)
	AddDec128(comment []string, name string)
	AddFloat32(comment []string, name string)
	AddFloat64(comment []string, name string)
	AddString(comment []string, name string)
	AddStr(comment []string, name string)
	AddBool(comment []string, name string)
	AddCustomType(comment []string, name string, info TypeRegistration)
}

// Source copy of srcobj.Source
type Source interface {
	Dump(w io.Writer) error
}

// Field represents a field of given type
type Field interface {
	Name() string
	TypeName() string
	Register(comment []string, registrator FieldRegistrator)
	GoName() string
}
