package srcobj

import (
	"fmt"
	"io"

	"github.com/sirkon/ldetool/internal/types"
)

// ///////////////// special hard to access type ///////////////////
type hardToAccessResultType string

func (h hardToAccessResultType) ResultType() string {
	return fmt.Sprintf("(res %s)", string(h))
}

func Go2ResultType(extTypes map[string]types.TypeRegistration, useString bool, goType string) (hardToAccessResultType, error) {
	var res hardToAccessResultType
	if goType == "string" {
		if useString {
			return hardToAccessResultType("string"), nil
		}
		return hardToAccessResultType("[]byte"), nil
	} else {
		if types.IsBuiltin(goType) {
			fieldType := types.Builtin("", goType)
			return hardToAccessResultType(fieldType.Native()), nil
		}
		ext, ok := extTypes[goType]
		if !ok {
			return res, fmt.Errorf("unsupported type `\033[1m%s\033[0m`", goType)
		}
		return hardToAccessResultType(ext.String()), nil
	}
}

// ExtractorResult
type ExtractorResult struct{}

func (e ExtractorResult) ResultType() string {
	return "(bool, error)"
}

// /////////////////////////////////////////////////////////////////

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
