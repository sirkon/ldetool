package gogen

import "strings"

// HeadString checks if the rest starts with the given string and passes it
func (g *Generator) HeadString(anchor string) {
	g.regVar(g.curRestVar(), "[]byte")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("head_string", g.curBody, TParams{
		Rest:       g.curRestVar(),
		ConstName:  constName,
		ConstValue: anchor,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// MayBeHeadString checks if the rest starts with the given string and passes it if yes. Otherwise do nothing
func (g *Generator) MayBeHeadString(anchor string) {
	g.regVar(g.curRestVar(), "[]byte")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("head_string_maybe", g.curBody, TParams{
		Rest:       g.curRestVar(),
		ConstName:  constName,
		ConstValue: anchor,
	})
}

// HeadChar checks if rest starts with the given char
func (g *Generator) HeadChar(char string) {
	g.regVar(g.curRestVar(), "[]byte")
	g.tc.MustExecute("head_char", g.curBody, TParams{
		Rest:       g.curRestVar(),
		Char:       char,
		Serious:    g.serious,
		Namespace:  strings.Join(g.namespaces, "."),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
	})
	g.abandon()
}

// MayBeHeadChar checks if rest starts with the given char
func (g *Generator) MayBeHeadChar(char string) {
	g.regVar(g.curRestVar(), "[]byte")
	g.tc.MustExecute("head_char_maybe", g.curBody, TParams{
		Rest: g.curRestVar(),
		Char: char,
	})
}
