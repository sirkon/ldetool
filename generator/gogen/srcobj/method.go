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

var go2resultType = map[string]hardToAccessResultType{
	"int8":    hardToAccessResultType("int8"),
	"int16":   hardToAccessResultType("int16"),
	"int32":   hardToAccessResultType("int32"),
	"int64":   hardToAccessResultType("int64"),
	"uint8":   hardToAccessResultType("uint8"),
	"uint16":  hardToAccessResultType("uint16"),
	"uint32":  hardToAccessResultType("uint32"),
	"uint64":  hardToAccessResultType("uint64"),
	"float32": hardToAccessResultType("float32"),
	"float64": hardToAccessResultType("float64"),
	"string":  hardToAccessResultType("[]byte"),
}

func Go2ResultType(goType string) hardToAccessResultType {
	res, ok := go2resultType[goType]
	if !ok {
		panic(fmt.Errorf("unsupported type `\033[1m%s\033[0m`", goType))
	}
	return res
}

// ExtractorResult
type ExtractorResult struct{}

func (e ExtractorResult) ResultType() string {
	return "(bool, error)"
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
		body:    &Body{},
	}
	return res
}

// Body returns method body
func (m *Method) Body() *Body {
	return m.body
}

// Dump implementation of source
func (m *Method) Dump(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "// %s ...\n", m.name); err != nil {
		return err
	}
	if _, err := fmt.Fprintf(w, "func (p *%s) %s(%s) %s {\n", m.objType, m.name, m.params, m.resType.ResultType()); err != nil {
		return err
	}
	if err := m.body.Dump(w); err != nil {
		return err
	}
	_, err := io.WriteString(w, "}\n")
	return err
}
