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
	"int":     hardToAccessResultType("int"),
	"int8":    hardToAccessResultType("int8"),
	"int16":   hardToAccessResultType("int16"),
	"int32":   hardToAccessResultType("int32"),
	"int64":   hardToAccessResultType("int64"),
	"uint8":   hardToAccessResultType("uint8"),
	"uint":    hardToAccessResultType("uint"),
	"uint16":  hardToAccessResultType("uint16"),
	"uint32":  hardToAccessResultType("uint32"),
	"uint64":  hardToAccessResultType("uint64"),
	"float32": hardToAccessResultType("float32"),
	"float64": hardToAccessResultType("float64"),
	"string":  hardToAccessResultType("[]byte"),
}

func Go2ResultType(useString bool, goType string) (hardToAccessResultType, error) {
	var res hardToAccessResultType
	var ok bool
	if useString && goType == "string" {
		res = hardToAccessResultType("string")
	} else {
		res, ok = go2resultType[goType]
		if !ok {
			return res, fmt.Errorf("unsupported type `\033[1m%s\033[0m`", goType)
		}
	}
	return res, nil
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
	objType   string
	name      string
	params    string // Not to be edited
	resType   ResultType
	body      *Body
	useString bool
}

// NewExtractor creates extractor definition
func NewExtractor(useString bool, objType string) *Method {
	res := &Method{
		objType:   objType,
		name:      "Extract",
		params:    "line " + RightType(useString),
		resType:   ExtractorResult{},
		body:      &Body{},
		useString: useString,
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
