package types

var _ TypeRegistration = LocalType{}

// LocalType local package type view
type LocalType struct {
	Name string
}

func (t LocalType) String() string {
	return t.Name
}

func (LocalType) typeRegistration() {}
