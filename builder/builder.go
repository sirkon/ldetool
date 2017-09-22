package builder

import (
	"fmt"
	"io"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/ast"

	"github.com/sirkon/ldetool/generator"
	"github.com/sirkon/message"
)

// Builder creates target sources using Generator object
type Builder struct {
	pkgName      string
	gen          generator.Generator
	dest         io.Writer
	recoverPanic bool
	gotify       *gotify.Gotify

	errToken antlr.Token
}

// NewBuilder constructor
func NewBuilder(pn string, g generator.Generator, d io.Writer, gfy *gotify.Gotify) *Builder {
	return &Builder{
		pkgName:      pn,
		gen:          g,
		dest:         d,
		recoverPanic: true,
		gotify:       gfy,
	}
}

// DontRecover tells not to recover panics
func (b *Builder) DontRecover() {
	b.recoverPanic = false
}

// BuildRule builds shit from the data
func (b *Builder) BuildRule(rule *ast.RuleItem) (err error) {
	if b.recoverPanic {
		defer func() {
			if r := recover(); r != nil {
				var ok bool
				if err, ok = r.(error); !ok {
					panic(r)
				}
			}
		}()
	}
	b.gen.UseRule(rule.Name, rule.NameToken)
	generators, err := b.composeRules(NewPrefix(), b.gen, rule.Actions)
	if err != nil {
		return
	}
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

// checkField if field name is goish public variable
func (b *Builder) checkField(field ast.Field) error {
	if b.gotify.Public(field.Name) != field.Name {
		b.errToken = field.NameToken
		return fmt.Errorf("Wrong taker identifier `%s`, must be %s", field.Name, b.gotify.Public(field.Name))
	}
	return nil
}

func (b *Builder) composeRules(gPrefix Prefix, g generator.Generator, a []*ast.ActionItem) (generators []func(), err error) {
	if len(a) == 0 {
		return
	}
	it := a[0]
	message.Info(it)

	// Set on stress
	if it.ErrorOnMismatch {
		generators = append(generators, g.Stress)
	}

	// TakeUntilOrRest
	if it.TakeUntilOrRest != nil {
		item := it.TakeUntilOrRest
		if err = b.checkField(item.Field); err != nil {
			return
		}
		g.RegGravity(gPrefix.Add(item.Field.Name).String())
		switch item.Limit.Type {
		case ast.String:
			generators = append(generators, func() {
				g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
				g.TakeBeforeString(
					item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper, true)
			})
		case ast.Char:
			generators = append(generators, func() {
				g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
				g.TakeBeforeChar(
					item.Field.Name, item.Field.Type, item.Limit.Value, item.Limit.Lower, item.Limit.Upper, true)
			})
		}
	}

	// TakeUntil
	if it.Take != nil {
		item := it.Take
		if err = b.checkField(item.Field); err != nil {
			return
		}
		g.RegGravity(gPrefix.Add(item.Field.Name).String())
		lower := item.Limit.Lower
		upper := item.Limit.Upper
		if item.Limit.Close {
			lower = -1
			upper = -1
		}
		switch item.Limit.Type {
		case ast.String:
			generators = append(generators, func() {
				g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
				g.TakeBeforeString(
					item.Field.Name, item.Field.Type, item.Limit.Value, lower, upper, false)
			})
		case ast.Char:
			generators = append(generators, func() {
				g.AddField(item.Field.Name, item.Field.Type, item.Field.NameToken)
				g.TakeBeforeChar(
					item.Field.Name, item.Field.Type, item.Limit.Value, lower, upper, false)
			})
		}
	}

	// TakeRest
	if it.TakeRest != nil {
		if err = b.checkField(it.TakeRest.Field); err != nil {
			return
		}
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
			g.HeadString(it.StartWithString.Value, false)
		})
	}

	// Head char
	if it.StartWithChar != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.HeadChar(it.StartWithChar.Value, false)
		})
	}

	// Probably head string
	if it.MayBeStartWithString != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.HeadString(it.MayBeStartWithString.Value, true)
		})
	}

	// Probably head char
	if it.MayBeStartWithChar != nil {
		g.RegGravity(gPrefix.String())
		generators = append(generators, func() {
			g.HeadChar(it.MayBeStartWithChar.Value, true)
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
		var lower int
		var upper int
		if l.Close {
			lower = -1
			upper = -1
		} else {
			lower = l.Lower
			upper = l.Upper
		}

		if lower == upper && lower > 0 {
			// Fixed position check
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupFixedString(l.Value, lower, false)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupFixedChar(l.Value, lower, false)
				})
			}
		} else {
			// It is either short or limited/bounded lookup
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupString(l.Value, lower, upper, false)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupChar(l.Value, lower, upper, false)
				})

			}
		}
	}

	// Passes until
	if it.PassOrIgnore != nil {
		g.RegGravity(gPrefix.String())
		l := it.PassOrIgnore.Limit
		var lower int
		var upper int
		if l.Close {
			lower = -1
			upper = -1
		} else {
			lower = l.Lower
			upper = l.Upper
		}

		if lower == upper && lower > 0 {
			// Fixed position check
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupFixedString(l.Value, lower, true)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupFixedChar(l.Value, lower, true)
				})
			}
		} else {
			// It is either short or limited/bounded lookup
			switch l.Type {
			case ast.String:
				generators = append(generators, func() {
					g.LookupString(l.Value, lower, upper, true)
				})
			case ast.Char:
				generators = append(generators, func() {
					g.LookupChar(l.Value, lower, upper, true)
				})

			}
		}
	}

	// Optional area
	if it.Option != nil {
		if b.gotify.Public(it.Option.Name) != it.Option.Name {
			b.errToken = it.Option.NameToken
			err = fmt.Errorf("Wrong option identifier %s, must be %s", it.Option.Name, b.gotify.Public(it.Option.Name))
			return
		}
		generators = append(generators, func() {
			g.OpenOptionalScope(it.Option.Name, it.Option.NameToken)
		})
		var newgens []func()
		newgens, err = b.composeRules(gPrefix.Add(it.Option.Name), g, it.Option.Actions)
		if err != nil {
			return
		}
		message.Infof("End of option \033[1m%s\033[0m", it.Option.Name)
		generators = append(generators, newgens...)
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

	newgens, err := b.composeRules(gPrefix, g, a[1:])
	if err != nil {
		return
	}
	return append(generators, newgens...), nil
}

// ErrorToken returns token where error happened
func (b *Builder) ErrorToken() antlr.Token {
	return b.errToken
}
