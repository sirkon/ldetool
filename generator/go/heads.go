package gogen

// HeadString checks if the rest starts with the given string and passes it
func (g *Generator) HeadString(anchor string) {
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("head_string", g.body, TParams{
		ConstName: constName,
	})
	g.headStringError(constName)
}

// MayBeHeadString checks if the rest starts with the given string and passes it if yes. Otherwise do nothing
func (g *Generator) MayBeHeadString(anchor string) {
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("head_string", g.body, TParams{
		ConstName: constName,
	})
}

// HeadChar checks if rest starts with the given char
func (g *Generator) HeadChar(char string) {
	g.tc.MustExecute("head_char", g.body, TParams{
		Char: char,
	})
	g.headCharError(char)
}

// MayBeHeadChar checks if rest starts with the given char
func (g *Generator) MayBeHeadChar(char string) {
	g.tc.MustExecute("head_char", g.body, TParams{
		Char: char,
	})
}
