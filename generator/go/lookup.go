package gogen

// LookupString ...
func (g *Generator) LookupString(anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_string", g.body, TParams{
		ConstName: constName,
	})
	g.lookupStringError(constName)
}

// LookupLimitedString ...
func (g *Generator) LookupLimitedString(anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_limited_string", g.body, TParams{
		ConstName: constName,
		Upper:     upper,
	})
	g.lookupLimitedStringError(constName, upper)
}

// LookupBoundedString ...
func (g *Generator) LookupBoundedString(anchor string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("lookup_bounded_string", g.body, TParams{
		ConstName: constName,
		Upper:     upper,
		Lower:     lower,
	})
	g.lookupBoundedStringError(constName, lower, upper)
}

// LookupChar ...
func (g *Generator) LookupChar(char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_char", g.body, TParams{
		Char: char,
	})
	g.lookupCharError(char)
}

// LookupLimitedChar ...
func (g *Generator) LookupLimitedChar(char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.tc.MustExecute("lookup_limited_char", g.body, TParams{
		Char:  char,
		Upper: upper,
	})
	g.lookupLimitedCharError(char, upper)
}

// LookupBoundedChar ...
func (g *Generator) LookupBoundedChar(char string, lower, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	g.lookupPush("", lower, upper)
	g.tc.MustExecute("lookup_bounded_char", g.body, TParams{
		Char:  char,
		Lower: lower,
		Upper: upper,
	})
	g.lookupBoundedCharError(char, lower, upper)
}
