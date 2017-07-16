package gogen

import (
	"fmt"
	"regexp"
)

// constNameFromContent generates name of the constant based on content
func (g *Generator) constNameFromContent(value string) string {
	w := NewMnemowriter()
	for _, r := range []rune(value) {
		w.WriteRune(r)
	}
	w.Flush()
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

// LookupString ...
func (g *Generator) LookupString(anchor string) {
}

// LookupLimitedString ...
func (g *Generator) LookupLimitedString(anchor string) {
	panic("not implemented")
}

// LookupBoundedString ...
func (g *Generator) LookupBoundedString(anchor string) {
	panic("not implemented")
}

// LookupChar ...
func (g *Generator) LookupChar(anchor string) {
	panic("not implemented")
}

// LookupLimitedChar ...
func (g *Generator) LookupLimitedChar(anchor string) {
	panic("not implemented")
}

// LookupBoundedChar ...
func (g *Generator) LookupBoundedChar(anchor string) {
	panic("not implemented")
}
