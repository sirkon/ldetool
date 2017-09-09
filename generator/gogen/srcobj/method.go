package srcobj

import (
	"fmt"
	"io"
)

/////////////////// special hard to access type ///////////////////
type hardToAccessResultType string

func (h hardToAccessResultType) ResultType() string {
	return fmt.Sprintf("(res %s)", string(h))
}

const (
	Int8Result  = hardToAccessResultType("int8")
	Int16Result = hardToAccessResultType("int16")
	Int32Result = hardToAccessResultType("int32")
	Int64Result = hardToAccessResultType("int64")

	Uint8Result  = hardToAccessResultType("uint8")
	Uint16Result = hardToAccessResultType("uint16")
	Uint32Result = hardToAccessResultType("uint32")
	Uint64Result = hardToAccessResultType("uint64")

	Float32Result = hardToAccessResultType("float32")
	Float64Result = hardToAccessResultType("float64")
)

// ExtractorResult
type ExtractorResult struct{}

func (e ExtractorResult) ResultType() string {
	return "(ok bool, err error)"
}

///////////////////////////////////////////////////////////////////

type ResultType interface {
	ResultType() string
}

// Method describes LDE generated method of extractor
type Method struct {
	objType string
	name    string
	params  string // Not to be edited
	resType ResultType
	body    *Body
}

// NewExtractor creates extractor definition
func NewExtractor(objType string) *Method {
	res := &Method{
		objType: objType,
		name:    "Extract",
		params:  "line []byte",
		resType: ExtractorResult{},
		body:    &Body{},
	}
	return res
}

// NewAccessor creates accessor definition
func NewAccessor(objType, name string, resType hardToAccessResultType) *Method {
	res := &Method{
		objType: objType,
		name:    name,
		params:  "",
		resType: resType,
	}
	return res
}

// Append appends new piece of source to the body
func (m *Method) Append(s Source) {
	m.body.Append(s)
}

// Dump implementation of source
func (m *Method) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "// %s ...\n", m.name); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "func (e *%s) %s(%s) %s {\n", m.objType, m.name, m.params, m.resType.ResultType()); err != nil {
		return err
	}
	if err := m.body.Dump(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}\n")
	return err
}
