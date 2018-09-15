package gogen

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

// OpenOptionalScope ...
func (g *Generator) OpenOptionalScope(name string, t antlr.Token) error {
	if err := g.regRightVar(g.curRestVar()); err != nil {
		return err
	}
	g.namespaces = append(g.namespaces, name)
	g.regLabel()
	if !g.anonymous() {
		g.obj = append(g.obj, g.curObj().AddSubstruct(name))
	}
	g.body.Append(srcobj.LineAssign{
		Receiver: g.curRestVar(),
		Expr:     srcobj.Raw(g.prevRestVar()),
	})
	if len(name) > 0 {
		g.addField(g.namespaces, "", t)
	}
	return nil
}

// CloseOptionalScope ...
func (g *Generator) CloseOptionalScope() error {
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
	return nil
}
