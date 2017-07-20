package gogen

import (
	"bytes"
	"io"
	"sort"
	"strings"

	"github.com/glossina/gosrcfmt"
	"github.com/glossina/gotify"
	"github.com/glossina/ldegen/templatecache"
	"github.com/glossina/ldegen/token"
)

// Name provides a link between token and string
type Name struct {
	name  string
	token *token.Token
}

// LookupItem keeps last lookup cooridinates
type LookupItem struct {
	Name  string
	Lower int
	Upper int
}

// Generator implementation of generator.Generator
// for Go target
type Generator struct {
	consts  map[string]string
	fields  map[string]Name
	vars    map[string]string
	imports map[string]string

	namespaces  []string
	lookupStack []LookupItem

	serious bool
	goish   *gotify.Gotify

	gravity []string
	pos     int

	tc        *templatecache.TemplateCache
	obj       *bytes.Buffer
	body      *bytes.Buffer
	opgetters *bytes.Buffer

	parserName string

	dgen *DecoderGen
}

// NewGenerator constructor
func NewGenerator(parserName string, goish *gotify.Gotify, tc *templatecache.TemplateCache) *Generator {
	res := &Generator{
		consts:      map[string]string{},
		fields:      map[string]Name{},
		imports:     map[string]string{},
		namespaces:  nil,
		lookupStack: nil,

		serious:    false,
		vars:       map[string]string{},
		goish:      goish,
		gravity:    nil,
		pos:        0,
		obj:        &bytes.Buffer{},
		body:       &bytes.Buffer{},
		opgetters:  &bytes.Buffer{},
		tc:         tc,
		parserName: parserName,
	}
	res.dgen = &DecoderGen{g: res}
	return res
}

// AddField ...
func (g *Generator) AddField(name string, fieldType string, t *token.Token) {
	g.addField(g.namespaces, name, t)
	goType := g.goType(fieldType)
	g.tc.MustExecute("struct_field", g.obj, TParams{
		Name: name,
		Type: goType,
	})
	return
}

// PassN passes first N characters if they are there, otherwise signal a error
func (g *Generator) PassN(n int) {
	g.tc.MustExecute("pass_n_items", g.body, TParams{
		Upper:      n,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// Stress mismatches should be treated as serious errors
func (g *Generator) Stress() {
	g.serious = true
}

// Generate writes into io.Writer
func (g *Generator) Generate(pkgName string, dest io.Writer) {
	var imports ImportSeq
	for path, name := range g.imports {
		imports = append(imports, Import{Name: name, Path: path})
	}
	sort.Sort(imports)

	var vars VarSeq
	for name, varType := range g.vars {
		vars = append(vars, Var{Name: name, Type: varType})
	}
	sort.Sort(vars)

	buf := &bytes.Buffer{}
	g.tc.MustExecute("parser_code", buf, ParserParams{
		Imports:    imports,
		Consts:     g.consts,
		Vars:       vars,
		Struct:     g.obj.String(),
		Parser:     g.body.String(),
		Getters:    g.opgetters.String(),
		ParserName: g.parserName,
		PkgName:    pkgName,
	})
	gosrcfmt.FormatReader(dest, buf)
}

// RegGravity registers center of gravity
func (g *Generator) RegGravity(name string) {
	g.gravity = append(g.gravity, name)
}

// AtEnd checks if the rest is empty
func (g *Generator) AtEnd() {
	g.tc.MustExecute("at_end", g.body, TParams{
		Serious: g.serious,
	})
}
