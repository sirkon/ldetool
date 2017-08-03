package gogen

import (
	"strings"

	"github.com/sirkon/ldetool/token"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t *token.Token) {
	g.regVar(g.curRestVar(), "[]byte")
	g.namespaces = append(g.namespaces, name)
	g.tc.MustExecute("open_option", g.curObj, TParams{
		Name:  name,
		Rest:  g.curRestVar(),
		PRest: g.prevRestVar(),
	})
	g.tc.MustExecute("open_option_scope", g.curBody, TParams{
		Rest:  g.curRestVar(),
		PRest: g.prevRestVar(),
	})
	g.addField(g.namespaces, name, t)
}

// ExitOptionalScope ...
func (g *Generator) ExitOptionalScope() {
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	g.tc.MustExecute("exit_option_scope", g.curObj, TParams{
		Rest:       g.curRestVar(),
		PRest:      g.prevRestVar(),
		Name:       validPoint,
		ScopeLabel: scopeLabelName,
	})
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	g.tc.MustExecute("close_scope", g.curObj, nil)
	scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
	validPoint := strings.Join(append(g.namespaces, "Valid"), ".")
	wasabandoned := g.abandoned()
	g.tc.MustExecute("close_option_scope", g.curBody, TParams{
		Rest:         g.curRestVar(),
		PRest:        g.prevRestVar(),
		Name:         validPoint, // this is true
		ScopeLabel:   scopeLabelName,
		WasAbandoned: wasabandoned,
	})
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
}
