package gogen

import (
	"fmt"
	"io"

	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/generator/gogen/srcobj"
)

// Name provides a link between token and string
type Name struct {
	name  string
	token antlr.Token
}

// Generator implementation of generator.Generator
// for Go target
type Generator struct {
	rules          map[string]antlr.Token // Register rule names to check duplication
	consts         map[string]string      // string constants for reuse
	fields         map[string]Name        // Field names obviously should have different names
	vars           map[string]string      // Function local variables
	imports        map[string]string      // Set of import paths
	scopeAbandoned map[string]bool        // Set check if the current scope was abandoned due to mismatch

	namespaces []string // Stack of namespaces (each item is a name of optional area)

	critical bool           // Treat mismatch errors as critical
	goish    *gotify.Gotify // identifier gotification service

	// gravity is not used yet, planned to clarify mismatch position against the rule, something like
	// "Could not find string `name=` right before Parameter.Name field" or
	// "Could not find string 'FETCH' between Time and Timeout fields", etc
	gravity []string

	// Source object representation primitives
	file       *srcobj.File     // File image
	body       *srcobj.Body     // Current method image
	obj        []*srcobj.Struct // Image of the structure
	optgetters *srcobj.Body     // Option getters for current structure
	vargen     *srcobj.Vars     // Function variables image
	decoderMap map[string]func(src srcobj.Source, dest string)

	ruleName string // Name of currently processing rule
}

// ErrorToken message
func (g *Generator) ErrorToken(token antlr.Token, format string, params ...interface{}) {
	panic(fmt.Sprintf(
		"%d:%d: %s",
		token.GetLine(),
		token.GetColumn(),
		fmt.Sprintf(format, params...),
	))
}

// NewGenerator constructor
func NewGenerator(goish *gotify.Gotify) *Generator {
	res := &Generator{
		rules:   map[string]antlr.Token{},
		consts:  map[string]string{},
		imports: map[string]string{},

		critical: false,
		vars:     map[string]string{},
		goish:    goish,
		gravity:  nil,

		file:     srcobj.NewFile(),
		ruleName: "",
	}

	res.decoderMap = map[string]func(src srcobj.Source, dest string){
		"int8":    res.decodeInt8,
		"int16":   res.decodeInt16,
		"int32":   res.decodeInt32,
		"int64":   res.decodeInt64,
		"uint8":   res.decodeUint8,
		"uint16":  res.decodeUint16,
		"uint32":  res.decodeUint32,
		"uint64":  res.decodeUint64,
		"float32": res.decodeFloat32,
		"float64": res.decodeFloat64,
		"string":  res.decodeString,
	}
	return res
}

func (g *Generator) valid() string {
	return "p." + strings.Join(g.namespaces, ".") + ".Valid"
}

func (g *Generator) label() string {
	namespaces := make([]string, len(g.namespaces))
	for i, chunk := range g.namespaces {
		if len(chunk) == 0 {
			namespaces[i] = "AnonymousArea"
		} else {
			namespaces[i] = chunk
		}
	}
	return g.goish.Private(strings.Join(namespaces, "_") + "_label")
}

func (g *Generator) curObj() *srcobj.Struct {
	return g.obj[len(g.obj)-1]
}

func (g *Generator) varName(name string) string {
	return "p." + strings.Join(append(g.namespaces, name), ".")
}

func (g *Generator) anonymous() bool {
	return len(g.namespaces) > 0 && len(g.namespaces[len(g.namespaces)-1]) == 0
}

// UseRule ...
func (g *Generator) UseRule(name string, t antlr.Token) {
	if len(g.ruleName) != 0 {
		panic(fmt.Errorf("attempt to use rule `%s` while the previous one (%s) was not pushed", name, g.ruleName))
	}
	if prev, ok := g.rules[name]; ok {
		panic(fmt.Errorf("%d: redeclaration of rule `%s` which has already been defined at line %d", t.GetLine(), name, prev))
	}
	g.rules[name] = t
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.vars = map[string]string{}
	g.namespaces = nil
	g.ruleName = name
	g.obj = []*srcobj.Struct{g.file.AddExtractor(name)}
	g.curObj().AddString("rest")
	g.body = g.file.AddExtract(name).Body()
	g.body.Append(srcobj.LineAssign{
		Receiver: "p.rest",
		Expr:     srcobj.Raw("line"),
	})
	g.optgetters = srcobj.NewBody()
	g.vargen = srcobj.NewVars()
	g.body.Append(g.vargen)
}

// AddField ...
func (g *Generator) AddField(name string, fieldType string, t antlr.Token) {
	g.addField(g.namespaces, name, t)
	s := g.curObj()
	fieldGen, ok := map[string]func(name string){
		"int8":    s.AddInt8,
		"int16":   s.AddInt16,
		"int32":   s.AddInt32,
		"int64":   s.AddInt64,
		"uint8":   s.AddUint8,
		"uint16":  s.AddUint16,
		"uint32":  s.AddUint32,
		"uint64":  s.AddUint64,
		"float32": s.AddFloat32,
		"float64": s.AddFloat64,
		"string":  s.AddString,
	}[fieldType]
	if !ok {
		fieldNames := []string{
			"int8", "int16", "int32", "int64",
			"uint8", "uint16", "uint32", "uint64",
			"float32", "float64",
			"string",
		}
		for i, fieldName := range fieldNames {
			fieldNames[i] = fmt.Sprintf("\033[1m%s\033[0m", fieldName)
		}
		g.ErrorToken(t, "unsupported type `\033[1m%s\033[0m`, must be one of %s",
			fieldType, strings.Join(fieldNames, ", "))
		panic(fmt.Errorf("unsupported type %s", fieldType))
	}
	fieldGen(name)
	return
}

func (g *Generator) failure(format string, params ...srcobj.Source) (res srcobj.Source) {
	if len(g.namespaces) > 0 {
		g.abandon()
		var pre srcobj.Source
		if g.anonymous() {
			pre = srcobj.Raw("")
		} else {
			pre = srcobj.Assign(g.valid(), srcobj.False)
		}
		res = srcobj.NewBody(
			pre,
			srcobj.Semicolon,
			srcobj.Goto(g.label()),
		)
	} else if g.critical {
		g.regImport("", "fmt")
		res = srcobj.ReturnError(format, params...)
	} else {
		res = srcobj.ReturnFail
	}
	return
}

func (g *Generator) rest() srcobj.Source {
	return srcobj.Raw(g.curRestVar())
}

func (g *Generator) indent() *srcobj.Body {
	g.body.Append(srcobj.Raw("\n"))
	return g.body
}

// PassN passes first N characters if they are there, otherwise signal a error
func (g *Generator) PassN(n int) {
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorGE(
				srcobj.NewCall(
					"len",
					g.rest(),
				),
				srcobj.Literal(n),
			),
			Then: srcobj.Assign(
				g.curRestVar(),
				srcobj.SliceFrom(g.rest(), srcobj.Literal(n)),
			),
			Else: g.failure(
				"Cannot skip first %d symbols: only %d left in the rest",
				srcobj.Literal(n),
				srcobj.NewCall("len", g.rest()),
			),
		},
	)
}

// Stress mismatches should be treated as critical errors
func (g *Generator) Stress() {
	g.critical = true
}

// Relax ...
func (g *Generator) Relax() {
	g.critical = false
}

// Generate writes into io.Writer
func (g *Generator) Generate(pkgName string, dest io.Writer) {
	g.file.PkgName(pkgName)
	if err := g.file.Dump(dest); err != nil {
		panic(err)
	}
}

// Push pushes data
func (g *Generator) Push() {
	if len(g.ruleName) == 0 {
		panic(fmt.Errorf("no rule has been set up to push it now"))
	}

	g.indent()
	g.body.Append(srcobj.ReturnOK)
	g.file.Append(g.optgetters)

	g.Relax()
	g.critical = false
	g.vars = map[string]string{}
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.ruleName = ""
}

// RegGravity registers center of gravity
func (g *Generator) RegGravity(name string) {
	g.gravity = append(g.gravity, name)
}

// AtEnd checks if the rest is empty
func (g *Generator) AtEnd() {
	g.body.Append(
		srcobj.If{
			Expr: srcobj.OperatorNEq(
				srcobj.NewCall("len", g.rest()),
				srcobj.Literal(0),
			),
			Then: g.failure(
				"The rest is not empty: %s",
				srcobj.NewCall("string", g.rest()),
			),
		},
	)
}
