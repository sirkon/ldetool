package gogen

import (
	"strings"

	"github.com/glossina/ldegen/token"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t *token.Token) {
	g.namespaces = append(g.namespaces, name)
	g.tc.MustExecute("open_option", g.obj, TParams{
		Name: name,
	})
	g.addField(g.namespaces, name, t)
}

// ExitOptionalScope ...
func (g *Generator) ExitOptionalScope() {
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("exit_option_scope", g.obj, TParams{
		Name:       validPoint,
		ScopeLabel: scopeLabelName,
	})
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	g.tc.MustExecute("close_scope", g.obj, nil)
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("close_option_scope", g.body, TParams{
		Name:       validPoint, // this is true
		ScopeLabel: scopeLabelName,
	})
}
