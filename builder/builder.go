package builder

import (
	"io"

	"github.com/sirkon/ldetool/ast"
	"github.com/sirkon/ldetool/generator"
)

// Builder creates target sources using Generator object
type Builder struct {
	pkgName      string
	gen          generator.Generator
	dest         io.Writer
	recoverPanic bool
}

// NewBuilder consturcot
func NewBuilder(pn string, g generator.Generator, d io.Writer) *Builder {
	return &Builder{
		pkgName:      pn,
		gen:          g,
		dest:         d,
		recoverPanic: true,
	}
}

// DontRecover tells not to recover panics
func (b *Builder) DontRecover() {
	b.recoverPanic = false
}

// BuildRule builds shit from the data
func (b *Builder) BuildRule(rule ast.RuleItem) (err error) {
	defer func() {
		if b.recoverPanic {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					panic(r)
				}
			}
		}
	}()
	b.gen.UseRule(rule.Name, rule.NameToken)
	generators := b.composeRules(NewPrefix(), b.gen, &rule.Actions)
	for _, item := range generators {
		func() {
			item()
		}()
	}
	b.gen.Push()
	return nil
}

// Build full source file
func (b *Builder) Build() (err error) {
	b.gen.Generate(b.pkgName, b.dest)
	return nil
}

func (b *Builder) composeRules(gPrefix Prefix, g generator.Generator, a *ast.ActionSequence) (generators []func()) {
	if a == nil {
		return
	}
	it := a.Head

	// Set on stress
	if a.ErrorOnMismatch {
		generators = append(generators, g.Stress)
	}

	// TakeUntilOrRest
	if it.TakeUntilOrRest != nil {
		item := it.TakeUntilOrRest
		g.RegGravity(gPrefix.Add(item.Field.Name).String())
		if item.Limit.Lower > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeBoundedStringOrRest(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeBoundedCharOrRest(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			}
		} else if item.Limit.Upper > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeLimitedStringOrRest(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeLimitedCharOrRest(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			}
		} else {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeStringOrRest(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeCharOrRest(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			}
		}
	}

	// TakeUntil
	if it.Take != nil {
		item := it.Take
		g.RegGravity(gPrefix.Add(item.Field.Name).String())
		if item.Limit.Lower > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeBoundedString(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeBoundedChar(
						item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper)
				})
			}
		} else if item.Limit.Upper > 0 {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeLimitedString(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeLimitedChar(item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Upper)
				})
			}
		} else {
			switch item.Limit.Type {
			case ast.String:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeString(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
					g.TakeBeforeChar(item.Field.Name, item.Field.Type, item.Limit.Value)
				})
			}
		}
	}

	// TakeRest
	if it.TakeRest != nil {
		g.RegGravity(gPrefix.Add(it.TakeRest.Field.Name).String())
		generators = append(generators, func() {
			g.AddField(it.TakeRest.Field.Name, it.TakeRest.Field.Type, it.TakeRest.Field.NameToken)
			g.TakeRest(it.TakeRest.Field.Name, it.TakeRest.Field.Type)
		})
	}

	// Head string
	if it.StartWithString != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.HeadString(it.StartWithString.Value)
		})
	}

	// Head char
	if it.StartWithChar != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.HeadChar(it.StartWithChar.Value)
		})
	}

	// Probably head string
	if it.MayBeStartWithString != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.MayBeHeadString(it.MayBeStartWithString.Value)
		})
	}

	// Probably head char
	if it.MayBeStartWithChar != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.MayBeHeadChar(it.MayBeStartWithChar.Value)
		})
	}

	// Passes first N symbols
	if it.PassFirst != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.PassN(int(*it.PassFirst))
		})
	}

	// Passes until
	if it.Pass != nil {
		g.RegGravity(gPrefix.String())
		l := it.Pass.Limit
		if l.Lower > 0 {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupBoundedString(l.Value, l.Lower, l.Upper)
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

	// Passes until
	if it.PassOrIgnore != nil {
		g.RegGravity(gPrefix.String())
		l := it.PassOrIgnore.Limit
		if l.Lower > 0 {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupBoundedStringOrIgnore(l.Value, l.Lower, l.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupBoundedCharOrIgnore(l.Value, l.Lower, l.Upper)
				})
			}
		} else if l.Upper > 0 {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupLimitedStringOrIgnore(l.Value, l.Upper)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupLimitedCharOrIgnore(l.Value, l.Upper)
				})
			}
		} else {
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupStringOrIgnore(l.Value)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupCharOrIgnore(l.Value)
				})
			}
		}
	}

	// Optional area
	if it.Option != nil {
		generators = append(generators, func() {
			g.OpenOptionalScope(it.Option.Name, it.Option.NameToken)
		})
		generators = append(generators, b.composeRules(gPrefix.Add(it.Option.Name), g, &it.Option.Actions)...)
		generators = append(generators, func() {
			g.CloseOptionalScope()
		})
	}

	// AtEnd
	if it.End != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.AtEnd()
		})
	}

	return append(generators, b.composeRules(gPrefix, g, a.Tail)...)
}
