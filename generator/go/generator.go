package gogen

import (
	"bytes"

	"github.com/DenisCheremisov/gotify"
	"github.com/DenisCheremisov/ldegen/token"
)

// Name provides a link between token and string
type Name struct {
	name  string
	token *token.Token
}

// Generator implementation of generator.Generator
// for Go target
type Generator struct {
	consts map[string]string
	fields map[string]Name

	namespaces  []string
	lookupStack []string

	serious bool
	vars    map[string]string
	goish   *gotify.Gotify

	gravity []string
	pos     int

	body *bytes.Buffer
}

// NewGenerator constructor
func NewGenerator(goish *gotify.Gotify) *Generator {
	return &Generator{
		consts:      map[string]string{},
		fields:      map[string]Name{},
		namespaces:  nil,
		lookupStack: nil,

		serious: false,
		vars:    map[string]string{},
		goish:   goish,
		gravity: nil,
		pos:     0,
		body:    &bytes.Buffer{},
	}
}

func (g *Generator) PassFoundString(anchor string) {
	panic("not implemented")
}

func (g *Generator) PassFoundChar() {
	panic("not implemented")
}

func (g *Generator) PassBoundedFoundString(anchor string) {
	panic("not implemented")
}

func (g *Generator) PassBoundedFoundChar(anchor string) {
	panic("not implemented")
}

func (g *Generator) PassN(n int) {
	panic("not implemented")
}

func (g *Generator) HeadString(anchor string) {
	panic("not implemented")
}

func (g *Generator) HeadChar(anchor string) {
	panic("not implemented")
}

func (g *Generator) AddField(name string, fieldType string) {
	panic("not implemented")
}

func (g *Generator) ConsumeField(name string) {
	panic("not implemented")
}

func (g *Generator) OpenOptionalScope(name string) {
	panic("not implemented")
}

func (g *Generator) ExitOptionalScope() {
	panic("not implemented")
}

func (g *Generator) Stress() {
	panic("not implemented")
}

func (g *Generator) Generate() error {
	panic("not implemented")
}
