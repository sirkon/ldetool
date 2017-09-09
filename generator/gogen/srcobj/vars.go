package srcobj

import (
	"fmt"
	"io"
	"sort"
)

// Vars for in-func variable declarations
type Vars struct {
	m map[string]string
}

// NewVars constructor
func NewVars() *Vars {
	return &Vars{m: map[string]string{}}
}

func (v *Vars) Declare(varName, varType string) error {
	prevType, ok := v.m[varName]
	if ok {
		if prevType != varType {
			return fmt.Errorf("attempt to redeclare variable %s from %s to %s", varName, prevType, varType)
		}
	}
	v.m[varName] = varType
	return nil
}

// Pos declare `var pos int` variable
func (v *Vars) Pos() {
	v.Declare("pos", "int")
}

// Dump implementation
func (v *Vars) Dump(w io.Writer) error {
	keys := make([]string, 0, len(v.m))
	for k := range v.m {
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))
	for _, variable := range keys {
		if _, err := fmt.Fprintf(w, "var %s %s", variable, v.m[variable]); err != nil {
			return err
		}
	}
	return nil
}
