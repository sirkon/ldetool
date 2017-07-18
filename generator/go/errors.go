/*
error (mismatch) processing errors
*/

package gogen

func (g *Generator) lookupStringError(constName string) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_string_serious", g.body, TParams{
				ConstName: constName,
				Gravity:   g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) lookupLimitedStringError(constName string, upper int) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_limited_string_serious", g.body, TParams{
				ConstName: constName,
				Upper:     upper,
				Gravity:   g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) lookupBoundedStringError(constName string, lower, upper int) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_bounded_string_serious", g.body, TParams{
				ConstName: constName,
				Upper:     upper,
				Lower:     lower,
				Gravity:   g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) lookupCharError(char string) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_char_serious", g.body, TParams{
				Char:    char,
				Gravity: g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) lookupLimitedCharError(char string, upper int) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_limited_char_serious", g.body, TParams{
				Char:    char,
				Upper:   upper,
				Gravity: g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) lookupBoundedCharError(char string, lower, upper int) {
	g.body.WriteString("if pos < 0 {")
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("errors_lookup_bounded_char_serious", g.body, TParams{
				Char:    char,
				Upper:   upper,
				Lower:   lower,
				Gravity: g.gravityTend(g.pos),
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
	g.body.WriteString("};")
}

func (g *Generator) headStringError(constName string) {
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("error_head_string", g.body, TParams{
				ConstName: constName,
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
}

func (g *Generator) headCharError(char string) {
	if len(g.namespaces) > 0 {
		g.ExitOptionalScope()
	} else {
		if g.serious {
			g.regImport("", "fmt")
			g.tc.MustExecute("error_head_char", g.body, TParams{
				Char: char,
			})
		} else {
			g.tc.MustExecute("soft_exit", g.body, nil)
		}
	}
}
