package types

import (
	"fmt"
	"regexp"
	"strings"
)

var builtins map[string]func(name string) Field
var natives map[string]struct{}
var declarables map[string]struct{}
var decimals map[string]struct{}
var backedBy map[string]string

// IsDecimal check if this type name is one of decimals
func IsDecimal(typeName string) bool {
	_, ok := decimals[typeName]
	return ok
}

// IsDeclarable check if type name can be declared explicitly (dec32, dec64 and dec128 can't)
func IsDeclarable(typeName string) bool {
	_, ok := declarables[typeName]
	return ok
}

// Declarables return list of builtin declarables type except decX.Y, which should be handled separately
func Declarables() []string {
	var res []string
	for name := range declarables {
		res = append(res, name)
	}
	return res
}

// IsNative check if type name is native
func IsNative(typeName string) bool {
	_, ok := natives[typeName]
	return ok
}

// IsBuiltin checks if given name is builtin
func IsBuiltin(typeName string) bool {
	_, ok := builtins[typeName]
	return ok
}

// Builtins return list of builtins
func Builtins() []string {
	var bs []string
	for b := range builtins {
		bs = append(bs, b)
	}
	return bs
}

// Builtin generates a field representation with given name and builtin type
func Builtin(fieldName, typeName string) Field {
	res, ok := builtins[typeName]
	if !ok {
		return nil
	}
	return res(fieldName)
}

// NeedCustomUnmarshaler checks if given typeName is a native type which needs custom unmarshaler, e.g. $int, $float64,
// etc
func NeedCustomUnmarshaler(typeName string) (ok bool, err error) {
	if !strings.HasPrefix(typeName, "$") {
		return false, nil
	}
	realName := typeName[1:]
	if IsNative(realName) {
		return true, nil
	}
	if IsDeclarable(realName) {
		return false, fmt.Errorf("cannot use %s with custom unmarshaling, use %s instead", realName, backedBy[realName])
	}
	isDecimal, err := regexp.MatchString(`dec\d+.\d+`, realName)
	if err != nil {
		panic(err)
	}
	if IsDecimal(realName) || isDecimal {
		return false, fmt.Errorf("use custom types for decimals")
	}
	return false, fmt.Errorf("using $ for unsupported type %s does not make a sense", realName)
}
