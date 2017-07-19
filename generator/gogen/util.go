package gogen

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/glossina/ldegen/token"
)

func (g *Generator) lookupPush(name string, lower, upper int) {
	g.lookupStack = append(g.lookupStack, LookupItem{
		Name:  name,
		Lower: lower,
		Upper: upper,
	})
}

func (g *Generator) lookupPop() LookupItem {
	res := g.lookupStack[len(g.lookupStack)-1]
	g.lookupStack = g.lookupStack[:len(g.lookupStack)-1]
	return res
}

// constNameFromContent generates name of the constant based on content
func (g *Generator) constNameFromContent(value string) string {
	w := NewMnemowriter()
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
	return res
}

// regVar registers variable of the given type
func (g *Generator) regVar(name, varType string) {
	if ok, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, name); !ok {
		panic(fmt.Errorf("Wrong variable name `\033[1m%s\033[0m`", name))
	} else if err != nil {
		panic(err)
	}
	if ok, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, varType); !ok {
		panic(fmt.Errorf("Wrong variable type `\033[1m%s\033[0m`", varType))
	} else if err != nil {
		panic(err)
	}

	if oldType, ok := g.vars[name]; ok {
		if oldType != varType {
			panic(fmt.Errorf(
				"local variable \033[1m%s\033[0m has been registered already with type \033[1m%s\033[0m",
				name, varType,
			))
		}
	}
	g.vars[name] = varType
}

func (g *Generator) regImport(importAs, path string) {
	if len(importAs) > 0 {
		if ok, err := regexp.MatchString(`^[a-zA-Z_][a-zA-Z0-9_]*$`, importAs); !ok {
			panic(fmt.Errorf("Wrong import name `\033[1m%s\033[0m`", importAs))
		} else if err != nil {
			panic(err)
		}
	}
	if importedAs, ok := g.imports[path]; ok {
		if importAs != importedAs {
			panic(fmt.Errorf(
				`Attempt to register import of "\033[1m%s\033[0m" as '\033[1m%s\033' while it has already been `+
					`imported as '\033[1m%s\033[0m'`,
				path, importAs, importedAs,
			))
		}
	}
	g.imports[path] = importAs
}

func (g *Generator) gravityTend(pos int) string {
	return ""
}

func (g *Generator) goType(inputType string) string {
	goTypeName, ok := map[string]string{
		"int8":    "int8",
		"int16":   "int16",
		"int32":   "int32",
		"int64":   "int64",
		"uint8":   "uint8",
		"uint16":  "uint16",
		"uint32":  "uint32",
		"uint64":  "uint64",
		"float32": "float32",
		"float64": "float64",
		"string":  "[]byte",
	}[inputType]
	if !ok {
		panic(fmt.Sprintf("Unsupported type `\033[1m%s\033[0m", inputType))
	}
	return goTypeName
}

func (g *Generator) decoderGen(inputType string) func(string, string) string {
	tmpl, ok := map[string]func(string, string) string{
		"int8":    g.dgen.Int8,
		"int16":   g.dgen.Int16,
		"int32":   g.dgen.Int32,
		"int64":   g.dgen.Int64,
		"uint8":   g.dgen.Uint8,
		"uint16":  g.dgen.Uint16,
		"uint32":  g.dgen.Uint32,
		"uint64":  g.dgen.Uint64,
		"float32": g.dgen.Float32,
		"float64": g.dgen.Float64,
		"string":  g.dgen.String,
	}[inputType]
	if !ok {
		panic(fmt.Sprintf("Unsupported type `\033[1m%s\033[0m", inputType))
	}
	return tmpl
}

func (g *Generator) addField(namespace []string, name string, t *token.Token) string {
	namespace = append(namespace, name)
	namespaced := strings.Join(namespace, ".")
	if ppp, ok := g.fields[name]; ok {
		panic(fmt.Sprintf(
			"Field `\033[1m%s\033[0m` redefiniton, previously declared at (%d, %d)",
			name, ppp.token.Line, ppp.token.Column))
	}
	g.fields[name] = Name{
		name:  namespaced,
		token: t,
	}
	return namespaced
}
