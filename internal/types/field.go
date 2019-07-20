package types

import (
	"io"
)

// FieldRegistrator abstration for field of given type adder
type FieldRegistrator interface {
	AddInt(name string)
	AddInt8(name string)
	AddInt16(name string)
	AddInt32(name string)
	AddInt64(name string)
	AddUint(name string)
	AddUint8(name string)
	AddUint16(name string)
	AddUint32(name string)
	AddUint64(name string)
	AddDec128(name string)
	AddFloat32(name string)
	AddFloat64(name string)
	AddString(name string)
	AddStr(name string)
	AddCustomType(name string, info TypeRegistration)
}

// Source copy of srcobj.Source
type Source interface {
	Dump(w io.Writer) error
}

// Field represents a field of given type
type Field interface {
	Name() string
	TypeName() string
	Register(registrator FieldRegistrator)
	Native() string
}
