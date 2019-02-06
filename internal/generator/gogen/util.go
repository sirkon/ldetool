package gogen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/mnemo"
)

// constNameFromContent generates name of the constant based on content
func (g *Generator) constNameFromContent(value string) string {
	w := mnemo.New()
	for _, r := range []rune(value) {
		_, _ = w.WriteRune(r)
	}
	_ = w.Flush()
	res := w.String()

	if ok, err := regexp.MatchString(`^\d.*$`, res); ok {
		res = "const_" + res
	} else if err != nil {
		panic(err)
	}
	res = g.goish.Private(res)
	newRes := res
	i := 2
	for {
		if cst, ok := g.consts[newRes]; !ok || (cst == value) {
			res = newRes
			break
		}
		newRes = g.goish.Private(fmt.Sprintf("%s_case_%d", res, i))
		i++
	}
	g.consts[res] = value
	g.file.AddConst(res, value)
	return res
}

// regVar registers variable of the given type
func (g *Generator) regVar(name, varType string) error {
	if name == "p.Rest" {
		return nil
	}
	if ok, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name); !ok {
		return fmt.Errorf("Wrong variable name `\033[1m%s\033[0m`", name)
	} else if err != nil {
		return err
	}
	if ok, err := regexp.MatchString(`^(?:\[\])?[a-zA-Z_][a-zA-Z0-9_]*$`, varType); !ok {
		return fmt.Errorf("Wrong variable type `\033[1m%s\033[0m`", varType)
	} else if err != nil {
		return err
	}

	if oldType, ok := g.vars[name]; ok {
		if oldType != varType {
			return fmt.Errorf(
				"local variable \033[1m%s\033[0m has been registered already with type \033[1m%s\033[0m",
				name, varType,
			)
		}
	}
	g.vars[name] = varType
	g.vargen.Declare(name, varType)
	return nil
}

func (g *Generator) regImport(importAs, path string) error {
	if len(importAs) > 0 {
		if ok, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, importAs); !ok {
			return fmt.Errorf("Wrong import name `\033[1m%s\033[0m`", importAs)
		} else if err != nil {
			return err
		}
	}
	if importedAs, ok := g.imports[path]; ok {
		if importAs != importedAs {
			return fmt.Errorf(
				`Attempt to register import of "\033[1m%s\033[0m" as '\033[1m%s\033' while it has already been `+
					`imported as '\033[1m%s\033[0m'`,
				path, importAs, importedAs,
			)
		}
	}
	g.imports[path] = importAs
	g.file.AddNamedImport(importAs, path)
	return nil
}

func (g *Generator) gravityTend(pos int) string {
	return ""
}

func (g *Generator) tmpSuspectancy(inputType string) bool {
	suspected, ok := map[string]bool{
		"int":     true,
		"int8":    true,
		"int16":   true,
		"int32":   true,
		"int64":   true,
		"uint":    true,
		"uint8":   true,
		"uint16":  true,
		"uint32":  true,
		"uint64":  true,
		"hex":     true,
		"hex8":    true,
		"hex16":   true,
		"hex32":   true,
		"hex64":   true,
		"oct":     true,
		"oct8":    true,
		"oct16":   true,
		"oct32":   true,
		"oct64":   true,
		"dec16":   true,
		"dec32":   true,
		"dec64":   true,
		"float32": true,
		"float64": true,
		"string":  false,
		"str":     true,
	}[inputType]
	if !ok {
		panic(fmt.Errorf("Unsupported type `\033[1m%s\033[0m`", inputType))
	}
	return suspected
}

func (g *Generator) addField(namespace []string, name string, t antlr.Token) string {
	namespace = append(namespace, name)
	namespaced := strings.Join(namespace, ".")
	if ppp, ok := g.fields[g.fullName(name)]; ok {
		panic(fmt.Sprintf(
			"%d:%d: Field `\033[1m%s\033[0m` redefiniton, previously declared at (%d, %d)",
			t.GetLine(), t.GetColumn(),
			name, ppp.token.GetLine(), ppp.token.GetColumn()))
	}
	g.fields[g.fullName(name)] = Name{
		name:  namespaced,
		token: t,
	}
	return namespaced
}

func (g *Generator) fullName(name string) string {
	return strings.Join(append(g.namespaces, name), "")
}

func (g *Generator) getAccessName() string {
	return strings.Join(g.namespaces, ".")
}

func (g *Generator) abandon() {
	if len(g.namespaces) > 0 {
		g.scopeAbandoned[g.getAccessName()] = true
	}
}

func (g *Generator) abandoned() bool {
	_, ok := g.scopeAbandoned[g.getAccessName()]
	return ok
}

func (g *Generator) curRestVar() string {
	if len(g.namespaces) == 0 {
		return "p.Rest"
	}
	return g.goish.Private(fmt.Sprintf("rest%d", len(g.namespaces)))
}

func (g *Generator) prevRestVar() string {
	if len(g.namespaces) <= 1 {
		return "p.Rest"
	}
	return g.goish.Private(strings.Join(g.namespaces[:len(g.namespaces)-1], "_") + "_rest")
}
