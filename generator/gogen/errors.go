/*
error (mismatch) processing errors
*/

package gogen

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
