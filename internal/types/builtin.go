package types

var builtins map[string]func(name string) Field

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
