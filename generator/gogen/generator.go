package gogen

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/glossina/gosrcfmt"
	"github.com/glossina/gotify"
	"github.com/glossina/ldetool/templatecache"
	"github.com/glossina/ldetool/token"
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
	rules          map[string]*token.Token
	consts         map[string]string
	fields         map[string]Name
	vars           map[string]string
	imports        map[string]string
	scopeAbandoned map[string]bool

	namespaces  []string
	lookupStack []LookupItem

	serious bool
	goish   *gotify.Gotify

	gravity []string
	pos     int

	tc        *templatecache.TemplateCache
	obj       *bytes.Buffer
	body      *bytes.Buffer
	curObj    *bytes.Buffer
	curBody   *bytes.Buffer
	opgetters *bytes.Buffer

	parserName string

	dgen         *DecoderGen
	tmpSuspected bool
}

// NewGenerator constructor
func NewGenerator(goish *gotify.Gotify, tc *templatecache.TemplateCache) *Generator {
	res := &Generator{
		rules:          map[string]*token.Token{},
		consts:         map[string]string{},
		fields:         map[string]Name{},
		imports:        map[string]string{},
		scopeAbandoned: map[string]bool{},
		namespaces:     nil,
		lookupStack:    nil,

		serious:    false,
		vars:       map[string]string{},
		goish:      goish,
		gravity:    nil,
		pos:        0,
		curObj:     &bytes.Buffer{},
		curBody:    &bytes.Buffer{},
		obj:        &bytes.Buffer{},
		body:       &bytes.Buffer{},
		opgetters:  &bytes.Buffer{},
		tc:         tc,
		parserName: "",
	}
	res.dgen = &DecoderGen{g: res}
	return res
}

// UseRule ...
func (g *Generator) UseRule(name string, t *token.Token) {
	if len(g.parserName) != 0 {
		panic(fmt.Errorf("Attempt to use rule `%s` while the previous one (%s) was not pushed", name, g.parserName))
	}
	if prev, ok := g.rules[name]; ok {
		panic(fmt.Errorf("%d: redeclaration of rule `%s` which has already been defined at line %d", t.Line, name, prev.Line))
	}
	g.rules[name] = t
	g.parserName = name
}

// AddField ...
func (g *Generator) AddField(name string, fieldType string, t *token.Token) {
	g.addField(g.namespaces, name, t)
	goType := g.goType(fieldType)
	g.tc.MustExecute("struct_field", g.curObj, TParams{
		Name: name,
		Type: goType,
	})
	return
}

// PassN passes first N characters if they are there, otherwise signal a error
func (g *Generator) PassN(n int) {
	g.tc.MustExecute("pass_n_items", g.curBody, TParams{
		Rest:       g.curRestVar(),
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

	buf := &bytes.Buffer{}
	g.tc.MustExecute("parser_code", buf, ParserParams{
		Imports: imports,
		Consts:  g.consts,
		Struct:  g.obj.String(),
		Parser:  g.body.String(),
		Getters: g.opgetters.String(),
		PkgName: pkgName,
	})
	gosrcfmt.FormatReader(dest, buf)
}

// Push pushes data
func (g *Generator) Push() {
	if len(g.parserName) == 0 {
		panic(fmt.Errorf("No rule has been set up to push it now"))
	}
	g.tc.MustExecute("struct_body", g.body, ParserParams{
		Struct:     g.curObj.String(),
		ParserName: g.parserName,
	})
	g.curObj.Reset()

	var vars VarSeq
	for name, varType := range g.vars {
		vars = append(vars, Var{Name: name, Type: varType})
	}
	sort.Sort(vars)
	g.tc.MustExecute("parser_body", g.body, ParserParams{
		Parser:     g.curBody.String(),
		ParserName: g.parserName,
		Vars:       vars,
	})
	g.curBody.Reset()
	_, _ = io.Copy(g.body, g.opgetters)
	g.opgetters.Reset()

	g.serious = false
	g.lookupStack = nil
	g.vars = map[string]string{}
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.parserName = ""
}

// RegGravity registers center of gravity
func (g *Generator) RegGravity(name string) {
	g.gravity = append(g.gravity, name)
}

// AtEnd checks if the rest is empty
func (g *Generator) AtEnd() {
	g.tc.MustExecute("at_end", g.curBody, TParams{
		Serious: g.serious,
	})
}

// TmpOn sets tmpSuspected state to on
func (g *Generator) TmpOn() {
	g.tmpSuspected = true
}

// TmpOff sets tmpSuspected state to off
func (g *Generator) TmpOff() {
	g.tmpSuspected = false
}
