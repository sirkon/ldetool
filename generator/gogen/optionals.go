package gogen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t antlr.Token) {
	g.regVar(g.curRestVar(), "[]byte")
	g.namespaces = append(g.namespaces, name)
	if !g.anonymous() {
		g.obj = append(g.obj, g.curObj().AddSubstruct(name))
	}
	g.body.Append(srcobj.LineAssign{
		Receiver: g.curRestVar(),
		Expr:     srcobj.Raw(g.prevRestVar()),
	})
	g.addField(g.namespaces, name, t)
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() {
	if !g.anonymous() {
		g.body.Append(
			srcobj.LineAssign{
				Receiver: g.valid(),
				Expr:     srcobj.Raw("true"),
			},
		)
	}
	g.body.Append(
		srcobj.LineAssign{
			Receiver: g.prevRestVar(),
			Expr:     g.rest(),
		},
	)
	if g.abandoned() {
		scopeLabelName := g.label()
		g.body.Append(srcobj.OperatorColon(srcobj.Raw(scopeLabelName), srcobj.Raw("")))
		g.indent()
	}
	if !g.anonymous() {
		g.obj = g.obj[:len(g.obj)-1]
	}
	g.namespaces = g.namespaces[:len(g.namespaces)-1]
}
