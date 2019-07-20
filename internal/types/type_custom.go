package types

var _ Field = FieldCustom{}

// FieldCustom represents a field of custom pre-registered type
type FieldCustom struct {
	FieldName string
	Type      TypeRegistration
}

// Name of a field with custom type
func (f FieldCustom) Name() string {
	return f.FieldName
}

// TypeName name of a custom type
func (f FieldCustom) TypeName() string {
	return f.Type.String()
}

// Register registers custom type
func (f FieldCustom) Register(registrator FieldRegistrator) {
	registrator.AddCustomType(f.FieldName, f.Type)
}

// GoName name of registered type
func (f FieldCustom) GoName() string {
	return f.Type.String()
}
