package gogen

import (
	"strings"

	"github.com/glossina/ldetool/token"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t *token.Token) {
	g.namespaces = append(g.namespaces, name)
	g.tc.MustExecute("open_option", g.curObj, TParams{
		Name: name,
	})
	g.addField(g.namespaces, name, t)
}

// ExitOptionalScope ...
func (g *Generator) ExitOptionalScope() {
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("exit_option_scope", g.curObj, TParams{
		Name:       validPoint,
		ScopeLabel: scopeLabelName,
	})
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	g.tc.MustExecute("close_scope", g.curObj, nil)
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	wasabandoned := g.abandoned()
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
	g.tc.MustExecute("close_option_scope", g.curBody, TParams{
		Name:         validPoint, // this is true
		ScopeLabel:   scopeLabelName,
		WasAbandoned: wasabandoned,
	})
}
