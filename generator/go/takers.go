package gogen

import "strings"

/* take_before_string
if pos = bytes.Index(p.rest, {{ .ConstName }}); pos >= 0 {
    {{ call .Decoder "p.rest[:pos]" .Dest }}
} else {
    return false, {{ if .Serious }}fmt.Errorf("Can't find a string limiting a value of {{ .Dest }}{{ else }}nil{{end}}
}
*/

// getterGen generates optional getter
func (g *Generator) getterGen(name, fieldType string) {
	if len(g.namespaces) == 0 {
		return
	}
	goType := g.goType(fieldType)
	g.tc.MustExecute("getter", g.opgetters, GParams{
		Access: strings.Join(g.namespaces, "."),
		Name:   name,
		Type:   goType,
	})
}

// TakeBeforeString ...
func (g *Generator) TakeBeforeString(name, fieldType, anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_string", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Serious:   g.serious,
		Dest:      item.name,
		Decoder:   method,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedString ...
func (g *Generator) TakeBeforeLimitedString(name, fieldType, anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_limited_string", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Serious:   g.serious,
		Dest:      item.name,
		Decoder:   method,
		Upper:     upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedString ...
func (g *Generator) TakeBeforeBoundedString(name, fieldType, anchor string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_bounded_string", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Serious:   g.serious,
		Dest:      item.name,
		Decoder:   method,
		Lower:     lower,
		Upper:     upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeChar ...
func (g *Generator) TakeBeforeChar(name, fieldType, char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_char", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Serious: g.serious,
		Dest:    item.name,
		Decoder: method,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedChar ...
func (g *Generator) TakeBeforeLimitedChar(name, fieldType, char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_limited_char", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Serious: g.serious,
		Dest:    item.name,
		Decoder: method,
		Upper:   upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedChar ...
func (g *Generator) TakeBeforeBoundedChar(name, fieldType, char string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_bounded_string", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Serious: g.serious,
		Dest:    item.name,
		Decoder: method,
		Lower:   lower,
		Upper:   upper,
	})
	g.getterGen(name, fieldType)
}

// TakeRest ...
func (g *Generator) TakeRest(name, fieldType string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_char", g.body, TParams{
		Name:    item.name,
		Serious: g.serious,
		Dest:    item.name,
		Decoder: method,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeStringOrRest ...
func (g *Generator) TakeBeforeStringOrRest(name, fieldType, anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_string_or_rest", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Dest:      item.name,
		Decoder:   method,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedStringOrRest ...
func (g *Generator) TakeBeforeLimitedStringOrRest(name, fieldType, anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_limited_string_or_rest", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Dest:      item.name,
		Decoder:   method,
		Upper:     upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedStringOrRest ...
func (g *Generator) TakeBeforeBoundedStringOrRest(name, fieldType, anchor string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_bounded_string_or_rest", g.body, TParams{
		ConstName: constName,
		Name:      item.name,
		Dest:      item.name,
		Decoder:   method,
		Lower:     lower,
		Upper:     upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeCharOrRest ...
func (g *Generator) TakeBeforeCharOrRest(name, fieldType, char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_string_or_rest", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Dest:    item.name,
		Decoder: method,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedCharOrRest ...
func (g *Generator) TakeBeforeLimitedCharOrRest(name, fieldType, char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_limited_string_or_rest", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Dest:    item.name,
		Decoder: method,
		Upper:   upper,
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedCharOrRest ...
func (g *Generator) TakeBeforeBoundedCharOrRest(name, fieldType, char string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_bounded_string_or_rest", g.body, TParams{
		Char:    char,
		Name:    item.name,
		Dest:    item.name,
		Decoder: method,
		Lower:   lower,
		Upper:   upper,
	})
	g.getterGen(name, fieldType)
}
