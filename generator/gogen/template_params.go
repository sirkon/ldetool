/*
parameter structs
*/

package gogen

// TParams describes template parameters
type TParams struct {
	ConstName  string
	ConstValue string
	Char       string
	Upper      int
	Lower      int

	Name string
	Type string

	Namespace  string
	ScopeLabel string
	Gravity    string
	Serious    bool
	Dest       string
	Decoder    func(string, string) string
}

// DParams describes parameters for decoder generators
type DParams struct {
	Source  string
	Dest    string
	Type    string
	Serious bool
	Bits    int
}

// GParams describes paramers for optional values
type GParams struct {
	Accesses   []string
	Access     string
	LongName   string
	Name       string
	Type       string
	ParserName string
}

// Import desribes package import
type Import struct {
	Name string
	Path string
}

// ImportSeq is a slice of imports
type ImportSeq []Import

func (is ImportSeq) Len() int {
	return len(is)
}

func (is ImportSeq) Less(i int, j int) bool {
	return is[i].Name < is[j].Name || is[i].Path < is[j].Path
}

func (is ImportSeq) Swap(i int, j int) {
	is[i], is[j] = is[j], is[i]
}

// Var describes variable
type Var struct {
	Name string
	Type string
}

// VarSeq is a slice of vars
type VarSeq []Var

// Len ...
func (vs VarSeq) Len() int {
	return len(vs)
}

// Less ...
func (vs VarSeq) Less(i int, j int) bool {
	return vs[i].Name < vs[j].Name
}

// Swap ...
func (vs VarSeq) Swap(i int, j int) {
	vs[i], vs[j] = vs[j], vs[i]
}

// ParserParams describes parameters for parser generation template
type ParserParams struct {
	Imports    ImportSeq
	Consts     map[string]string
	Vars       VarSeq
	Struct     string
	Parser     string
	Getters    string
	ParserName string
	PkgName    string
}
