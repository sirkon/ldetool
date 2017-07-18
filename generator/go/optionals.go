package gogen

import (
	"strings"

	"github.com/DenisCheremisov/ldegen/token"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t *token.Token) {
	g.addField(g.namespaces, name, t)
	g.tc.MustExecute("open_option", g.obj, TParams{
		Name: name,
	})
}

// ExitOptionalScope ...
func (g *Generator) ExitOptionalScope() {
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("exit_option_scope", g.obj, TParams{
		Name:   validPoint,
		Extra1: scopeLabelName,
	})
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	g.tc.MustExecute("close_scope", g.obj, nil)
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("close_option", g.obj, nil)
	g.tc.MustExecute("close_option_scope", g.obj, TParams{
		Name:   validPoint, // this is true
		Extra1: scopeLabelName,
	})
}
