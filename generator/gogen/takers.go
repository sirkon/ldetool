package gogen

import (
	"fmt"
	"strings"
)

/* take_before_string
if pos = bytes.Index(p.rest, {{ .ConstName }}); pos >= 0 {
    {{ call .Decoder "p.rest[:pos]" .Dest }}
} else {
    return false, {{ if .Serious }}fmt.Errorf("Can't find a string limiting a value of {{ .Dest }}{{ else }}nil{{end}}
}
*/

// getterGen generates optional getter
func (g *Generator) getterGen(name, fieldType string) {
	if len(g.parserName) == 0 {
		panic(fmt.Errorf("Rule set up required"))
	}
	if len(g.namespaces) == 0 {
		return
	}
	goType := g.goType(fieldType)
	accesses := []string{}
	for i := 1; i <= len(g.namespaces); i++ {
		accesses = append(accesses, strings.Join(g.namespaces[:i], "."))
	}

	g.tc.MustExecute("getter", g.opgetters, GParams{
		Accesses:   accesses,
		Access:     strings.Join(g.namespaces, "."),
		LongName:   g.goish.Public(strings.Join(append(g.namespaces, name), "_")),
		ParserName: g.parserName,
		Name:       name,
		Type:       goType,
	})
}

// TakeBeforeString ...
func (g *Generator) TakeBeforeString(name, fieldType, anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeBeforeLimitedString ...
func (g *Generator) TakeBeforeLimitedString(name, fieldType, anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_limited_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeBeforeBoundedString ...
func (g *Generator) TakeBeforeBoundedString(name, fieldType, anchor string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_bounded_string", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Lower:      lower,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeBeforeChar ...
func (g *Generator) TakeBeforeChar(name, fieldType, char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_char", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeBeforeLimitedChar ...
func (g *Generator) TakeBeforeLimitedChar(name, fieldType, char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_limited_char", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeBeforeBoundedChar ...
func (g *Generator) TakeBeforeBoundedChar(name, fieldType, char string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	if g.tmpSuspectancy(fieldType) {
		g.regVar("tmp", "[]byte")
	}

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_bounded_char", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Lower:      lower,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
	g.abandon()
}

// TakeRest ...
func (g *Generator) TakeRest(name, fieldType string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_rest", g.curBody, TParams{
		Name:       item.name,
		Type:       g.goType(fieldType),
		Serious:    g.serious,
		UseTmp:     g.tmpSuspected,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeStringOrRest ...
func (g *Generator) TakeBeforeStringOrRest(name, fieldType, anchor string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_string_or_rest", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Name:       item.name,
		Type:       g.goType(fieldType),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedStringOrRest ...
func (g *Generator) TakeBeforeLimitedStringOrRest(name, fieldType, anchor string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_limited_string_or_rest", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Name:       item.name,
		Type:       g.goType(fieldType),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedStringOrRest ...
func (g *Generator) TakeBeforeBoundedStringOrRest(name, fieldType, anchor string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	constName := g.constNameFromContent(anchor)
	g.tc.MustExecute("take_before_bounded_string_or_rest", g.curBody, TParams{
		ConstName:  constName,
		ConstValue: anchor,
		Name:       item.name,
		Type:       g.goType(fieldType),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Lower:      lower,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeCharOrRest ...
func (g *Generator) TakeBeforeCharOrRest(name, fieldType, char string) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_char_or_rest", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeLimitedCharOrRest ...
func (g *Generator) TakeBeforeLimitedCharOrRest(name, fieldType, char string, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_limited_char_or_rest", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Dest:       item.name,
		Decoder:    method,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}

// TakeBeforeBoundedCharOrRest ...
func (g *Generator) TakeBeforeBoundedCharOrRest(name, fieldType, char string, lower int, upper int) {
	g.regVar("pos", "int")
	g.regImport("", "bytes")
	g.regVar("tmp", "[]byte")

	item := g.fields[name]
	method := g.decoderGen(fieldType)

	g.tc.MustExecute("take_before_bounded_char_or_rest", g.curBody, TParams{
		Char:       char,
		Name:       item.name,
		Type:       g.goType(fieldType),
		Dest:       item.name,
		ScopeLabel: g.goish.Private(strings.Join(g.namespaces, "_") + "_label"),
		Decoder:    method,
		Lower:      lower,
		Upper:      upper,
		Namespace:  strings.Join(g.namespaces, "."),
	})
	g.getterGen(name, fieldType)
}
