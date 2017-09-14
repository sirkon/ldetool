package gogen

import (
	"strings"

	"github.com/sirkon/ldetool/generator/gogen/srcobj"
	"github.com/sirkon/ldetool/token"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t *token.Token) {
	g.regVar(g.curRestVar(), "[]byte")
	g.namespaces = append(g.namespaces, name)
	g.obj = append(g.obj, g.curObj().AddSubstruct(name))
	g.body.Append(srcobj.LineAssign{
		Receiver: g.curRestVar(),
		Expr:     srcobj.Raw(g.prevRestVar()),
	})
	g.addField(g.namespaces, name, t)
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	g.body.Append(
		srcobj.LineAssign{
			Receiver: g.valid(),
			Expr:     srcobj.Raw("true"),
		},
	)
	g.body.Append(
		srcobj.LineAssign{
			Receiver: g.prevRestVar(),
			Expr:     g.rest(),
		},
	)
	if g.abandoned() {
		scopeLabelName := g.goish.Private(strings.Join(g.namespaces, "_") + "_label")
		g.body.Append(srcobj.OperatorColon(srcobj.Raw(scopeLabelName), srcobj.Raw("")))
	}
	g.obj = g.obj[:len(g.obj)-1]
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
}
