package gogen

import "strings"

// LookupString ...
func (g *Generator) LookupString(anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupLimitedString ...
func (g *Generator) LookupLimitedString(anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_limited_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupBoundedString ...
func (g *Generator) LookupBoundedString(anchor string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_bounded_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Upper:      upper,
		Lower:      lower,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupChar ...
func (g *Generator) LookupChar(char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_char", g.curBody, TParams{
		Char:       char,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupLimitedChar ...
func (g *Generator) LookupLimitedChar(char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_limited_char", g.curBody, TParams{
		Char:       char,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupBoundedChar ...
func (g *Generator) LookupBoundedChar(char string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.lookupPush("", lower, upper)
	g.tc.MustExecute("lookup_bounded_char", g.curBody, TParams{
		Char:       char,
		Lower:      lower,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// LookupStringOrIgnore ...
func (g *Generator) LookupStringOrIgnore(anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_string_noerror", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// LookupLimitedStringOrIgnore ...
func (g *Generator) LookupLimitedStringOrIgnore(anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_limited_string_noerror", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// LookupBoundedStringOrIgnore ...
func (g *Generator) LookupBoundedStringOrIgnore(anchor string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_bounded_string_noerror", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Upper:      upper,
		Lower:      lower,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// LookupCharOrIgnore ...
func (g *Generator) LookupCharOrIgnore(char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_char_noerror", g.curBody, TParams{
		Char:       char,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// LookupLimitedCharOrIgnore ...
func (g *Generator) LookupLimitedCharOrIgnore(char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_limited_char_noerror", g.curBody, TParams{
		Char:       char,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}

// LookupBoundedCharOrIgnore ...
func (g *Generator) LookupBoundedCharOrIgnore(char string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.lookupPush("", lower, upper)
	g.tc.MustExecute("lookup_bounded_char_noerror", g.curBody, TParams{
		Char:       char,
		Lower:      lower,
		Upper:      upper,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
}
