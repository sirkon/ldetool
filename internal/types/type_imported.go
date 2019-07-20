package types

var _ TypeRegistration = ImportedType{}

// ImportedType imported type view
type ImportedType struct {
	Name       string
	ImportPath string
}

func (it ImportedType) String() string {
	return it.Name
}

func (it ImportedType) typeRegistration() {}
