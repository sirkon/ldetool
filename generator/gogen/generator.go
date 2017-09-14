package gogen

import (
	"fmt"
	"io"

	"strings"

	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/generator/gogen/srcobj"
	"github.com/sirkon/ldetool/token"
)

// Name provides a link between token and string
type Name struct {
	name  string
	token *token.Token
}

// LookupItem keeps last lookup cooridinates
type LookupItem struct {
	Name  string
	Lower int
	Upper int
}

// Generator implementation of generator.Generator
// for Go target
type Generator struct {
	rules          map[string]*token.Token
	consts         map[string]string
	fields         map[string]Name
	vars           map[string]string
	imports        map[string]string
	scopeAbandoned map[string]bool

	namespaces  []string
	lookupStack []LookupItem

	serious bool
	goish   *gotify.Gotify

	gravity []string
	pos     int

	file       *srcobj.File
	body       *srcobj.Body
	obj        []*srcobj.Struct
	optgetters *srcobj.Body
	vargen     *srcobj.Vars

	decoderMap map[string]func(src srcobj.Source, dest string)

	parserName string

	tmpSuspected bool
}

// NewGenerator constructor
func NewGenerator(goish *gotify.Gotify) *Generator {
	res := &Generator{
		rules:          map[string]*token.Token{},
		consts:         map[string]string{},
		fields:         map[string]Name{},
		imports:        map[string]string{},
		scopeAbandoned: map[string]bool{},
		namespaces:     nil,
		lookupStack:    nil,

		serious: false,
		vars:    map[string]string{},
		goish:   goish,
		gravity: nil,
		pos:     0,

		file:       srcobj.NewFile(),
		parserName: "",
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

func (g *Generator) curObj() *srcobj.Struct {
	return g.obj[len(g.obj)-1]
}

func (g *Generator) varName(name string) string {
	return "p." + strings.Join(append(g.namespaces, name), ".")
}

// UseRule ...
func (g *Generator) UseRule(name string, t *token.Token) {
	if len(g.parserName) != 0 {
		panic(fmt.Errorf("attempt to use rule `%s` while the previous one (%s) was not pushed", name, g.parserName))
	}
	if prev, ok := g.rules[name]; ok {
		panic(fmt.Errorf("%d: redeclaration of rule `%s` which has already been defined at line %d", t.Line, name, prev.Line))
	}
	g.rules[name] = t
	g.parserName = name
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
func (g *Generator) AddField(name string, fieldType string, t *token.Token) {
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
		panic(fmt.Errorf("unsupported type %s", fieldType))
	}
	fieldGen(name)
	return
}

func (g *Generator) failure(format string, params ...srcobj.Source) (res srcobj.Source) {
	if len(g.namespaces) > 0 {
		g.abandon()
		res = srcobj.NewBody(
			srcobj.Assign(g.valid(), srcobj.False),
			srcobj.Semicolon,
			srcobj.Goto(g.label()),
		)
	} else if g.serious {
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

// Stress mismatches should be treated as serious errors
func (g *Generator) Stress() {
	g.serious = true
}

// Relax ...
func (g *Generator) Relax() {
	g.serious = false
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
	if len(g.parserName) == 0 {
		panic(fmt.Errorf("no rule has been set up to push it now"))
	}

	g.indent()
	g.body.Append(srcobj.ReturnOK)
	g.file.Append(g.optgetters)

	g.Relax()
	g.serious = false
	g.lookupStack = nil
	g.vars = map[string]string{}
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.parserName = ""
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

// TmpOn sets tmpSuspected state to on
func (g *Generator) TmpOn() {
	g.tmpSuspected = true
}

// TmpOff sets tmpSuspected state to off
func (g *Generator) TmpOff() {
	g.tmpSuspected = false
}
