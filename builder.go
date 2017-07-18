package main

import (
	"fmt"
	"io"

	"github.com/DenisCheremisov/ldegen/ast"
	"github.com/DenisCheremisov/ldegen/generator"
	"github.com/DenisCheremisov/ldegen/token"
	"github.com/DenisCheremisov/message"
)

// Builder creates target sources using Generator object
type Builder struct {
	pkgName   string
	prevRules map[string]*token.Token
	gFactory  func() generator.Generator
	dFactory  func(pkgName, name string) io.Writer
}

// NewBuilder consturcot
func NewBuilder(pn string, dg func() generator.Generator, df func(name, pkg string) io.Writer) *Builder {
	return &Builder{
		prevRules: map[string]*token.Token{},
		pkgName:   pn,
		gFactory:  dg,
		dFactory:  df,
	}
}

// BuildRule builds shit from the data
func (b *Builder) BuildRule(pkgName string, rule ast.RuleItem) error {
	defer func() {
		if r := recover(); r != nil {
			message.Critical(r)
		}
	}()
	if t, ok := b.prevRules[rule.Name]; ok {
		return fmt.Errorf(
			"%d: Rule `\033[1m%s\033[0m` has already been defined at line %d", rule.NameToken.Line, rule.Name, t.Line)
	}
	b.prevRules[rule.Name] = rule.NameToken
	g := b.gFactory()
	d := b.dFactory(b.pkgName, rule.Name)
	generators := b.composeRules(g, &rule.Actions)
	for _, item := range generators {
		func() {
			item()
		}()
	}
	g.Generate(pkgName, rule.Name, d)
	return nil
}

func (b *Builder) composeRules(g generator.Generator, a *ast.ActionSequence) (generators []func()) {
	if a == nil {
		return
	}
	it := a.Head
	// TakeUntilOrRest
	if it.TakeUntilOrRest != nil {
		item := it.TakeUntilOrRest
		g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
		if item.Limit.Lower > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeBoundedStringOrRest(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeBoundedCharOrRest(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			}
		} else if item.Limit.Upper > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeLimitedStringOrRest(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeLimitedCharOrRest(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			}
		} else {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeStringOrRest(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeCharOrRest(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			}
		}
	}

	// TakeUntil
	if it.Take != nil {
		item := it.Take
		g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
		if item.Limit.Lower > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeBoundedString(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeBoundedChar(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			}
		} else if item.Limit.Upper > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeLimitedString(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeLimitedChar(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			}
		} else {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.TakeBeforeString(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.TakeBeforeChar(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			}
		}
	}

	// TakeRest
	if it.TakeRest != nil {
		g.AddField(it.TakeRest.Field.Name, it.TakeRest.Field.Type, it.TakeRest.Field.NameToken)
		generators = append(generators, func() {
			g.TakeRest(it.TakeRest.Field.Name, it.TakeRest.Field.Type)
		})
	}

	// Head string
	if it.StartWithString != nil {
		generators = append(generators, func() {
			g.HeadString(it.StartWithString.Value)
		})
	}

	// Head char
	if it.StartWithChar != nil {
		generators = append(generators, func() {
			g.HeadString(it.StartWithChar.Value)
		})
	}

	// Probably head string
	if it.MayBeStartWithString != nil {
		generators = append(generators, func() {
			g.MayBeHeadString(it.MayBeStartWithString.Value)
		})
	}

	// Probably head char
	if it.MayBeStartWithChar != nil {
		generators = append(generators, func() {
			g.MayBeHeadChar(it.MayBeStartWithChar.Value)
		})
	}

	// Passes first N symbols
	if it.PassFirst != nil {
		generators = append(generators, func() {
			g.PassN(int(*it.PassFirst))
		})
	}

	// Passes until
	if it.Pass != nil {
		l := it.Pass.Limit
		if l.Lower > 0 {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupBoundedString(l.Value, l.Upper, l.Lower)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupBoundedChar(l.Value, l.Lower, l.Upper)
				})
			}
		} else if l.Upper > 0 {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupLimitedString(l.Value, l.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupLimitedChar(l.Value, l.Upper)
				})
			}
		} else {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupString(l.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupChar(l.Value)
				})
			}
		}
	}

	// Optional area
	if it.Option != nil {
		generators = append(generators, func() {
			g.OpenOptionalScope(it.Option.Name)
		})
		generators = append(generators, b.composeRules(g, &it.Option.Actions)...)
		generators = append(generators, func() {
			g.CloseOptionalScope()
		})
	}

	return generators
}
