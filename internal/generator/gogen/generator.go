package gogen

import (
	"fmt"
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"

	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/generator/gogen/internal/srcobj"
)

// Name provides a link between token and string
type Name struct {
	name  string
	token antlr.Token
}

var _ generator.Generator = &Generator{}

// Generator implementation of generator.Generator
// for Go target
type Generator struct {
	useString      bool
	rules          map[string]antlr.Token // Register rule names to check duplication
	consts         map[string]string      // string constants for reuse
	fields         map[string]Name        // Field names obviously should have different names
	vars           map[string]string      // Function local variables
	imports        map[string]string      // Set of import paths
	scopeAbandoned map[string]bool        // Set check if the current scope was abandoned due to mismatch

	labels    map[string]struct{} // Label name must be unique
	labelName string              // Current label name, set on scope enter and add field actions

	namespaces []string // Stack of namespaces (each item is a name of optional area)

	critical bool           // Treat mismatch errors as critical
	goish    *gotify.Gotify // identifier gotification service

	platformType generator.PlatformType

	// gravity is not used yet, planned for mismatch position  clarification against the rule, something like
	// "Could not find string `name=` right before Parameter.Name field" or
	// "Could not find string 'FETCH' between Time and Timeout fields", etc
	gravity []string

	// Source object representation primitives
	file              *srcobj.File    // File image
	body              *srcobj.Body    // Current method image
	obj               []*srcobj.Strct // Image of the structure
	optgetters        *srcobj.Body    // Option getters for current structure
	vargen            *srcobj.Vars    // Function variables image
	decoderMap        map[string]func(src srcobj.Source, dest string)
	decimalDecoderMap map[string]func(src srcobj.Source, dest string, precision, scale int)

	ruleName string // Name of currently processing rule

	curField     Name
	curFieldType string

	rulePassCounter int // how many passes there in current rule. To be reset on a new one.
}

// PlatformType holds an information what type of platform to generate code for:
//    1. Little endiant
//    2. Big endian
//    3. Crossplatform
// There's a difference for fast short (up to 8 bytes) prefix checks in strings.
func (g *Generator) PlatformType(t generator.PlatformType) error {
	g.platformType = t
	return nil
}

// ErrorToken message
func (g *Generator) ErrorToken(token antlr.Token, format string, params ...interface{}) error {
	return fmt.Errorf(
		"%d:%d: %s",
		token.GetLine(),
		token.GetColumn(),
		fmt.Sprintf(format, params...),
	)
}

// NewGenerator constructor
func NewGenerator(useString bool, goish *gotify.Gotify) *Generator {
	res := &Generator{
		useString: useString,
		rules:     map[string]antlr.Token{},
		consts:    map[string]string{},
		imports:   map[string]string{},

		critical: false,
		vars:     map[string]string{},
		goish:    goish,
		gravity:  nil,

		file:     srcobj.NewFile(useString),
		ruleName: "",

		platformType: generator.Universal,
	}

	res.decoderMap = map[string]func(src srcobj.Source, dest string){
		"int":     res.decodeInt,
		"int8":    res.decodeInt8,
		"int16":   res.decodeInt16,
		"int32":   res.decodeInt32,
		"int64":   res.decodeInt64,
		"uint":    res.decodeUint,
		"uint8":   res.decodeUint8,
		"uint16":  res.decodeUint16,
		"uint32":  res.decodeUint32,
		"uint64":  res.decodeUint64,
		"hex":     res.decodeHex,
		"hex8":    res.decodeHex8,
		"hex16":   res.decodeHex16,
		"hex32":   res.decodeHex32,
		"hex64":   res.decodeHex64,
		"oct":     res.decodeOct,
		"oct8":    res.decodeOct8,
		"oct16":   res.decodeOct16,
		"oct32":   res.decodeOct32,
		"oct64":   res.decodeOct64,
		"float32": res.decodeFloat32,
		"float64": res.decodeFloat64,
		"string":  res.decodeString,
		"str":     res.decodeStr,
	}
	res.decimalDecoderMap = map[string]func(src srcobj.Source, dest string, precision, scale int){
		"dec32":  res.decodeDec32,
		"dec64":  res.decodeDec64,
		"dec128": res.decodeDec128,
	}
	return res
}

func (g *Generator) valid() string {
	return "p." + strings.Join(g.namespaces, ".") + ".Valid"
}

func (g *Generator) regLabel() {
	namespaces := make([]string, len(g.namespaces))
	for i, chunk := range g.namespaces {
		if len(chunk) == 0 {
			namespaces[i] = "AnonymousArea"
		} else {
			namespaces[i] = chunk
		}
	}
	suffix := ""
	i := 1
	for {
		labelName := g.goish.Private(strings.Join(namespaces, "_")+"_label") + suffix
		labelName = g.goish.Private(g.goish.Package(g.ruleName) + "_" + labelName)
		if _, ok := g.labels[labelName]; !ok {
			g.labelName = labelName
			break
		}
		i++
		suffix = fmt.Sprintf("%d", i)
	}
	g.labels[g.labelName] = struct{}{}
}

func (g *Generator) label() string {
	return g.labelName
}

func (g *Generator) curObj() *srcobj.Strct {
	return g.obj[len(g.obj)-1]
}

func (g *Generator) varName(name string) string {
	return "p." + strings.Join(append(g.namespaces, name), ".")
}

func (g *Generator) anonymous() bool {
	return len(g.namespaces) > 0 && len(g.namespaces[len(g.namespaces)-1]) == 0
}

// UseRule ...
func (g *Generator) UseRule(name string, t antlr.Token) error {
	if len(g.ruleName) != 0 {
		return fmt.Errorf("attempt to use rule `%s` while the previous one (%s) was not pushed", name, g.ruleName)
	}
	if prev, ok := g.rules[name]; ok {
		return fmt.Errorf("%d: redeclaration of rule `%s` which has already been defined at line %d", t.GetLine(), name, prev)
	}
	g.rules[name] = t
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.labels = map[string]struct{}{}
	g.vars = map[string]string{}
	g.namespaces = nil
	g.ruleName = name
	g.obj = []*srcobj.Strct{g.file.AddExtractor(name)}
	g.curObj().AddString("Rest")
	g.body = g.file.AddExtract(name).Body()
	g.body.Append(srcobj.LineAssign{
		Receiver: "p.Rest",
		Expr:     srcobj.Raw("line"),
	})
	g.optgetters = srcobj.NewBody()
	g.vargen = srcobj.NewVars()
	g.body.Append(g.vargen)
	g.rulePassCounter = 0
	return nil
}

// AddField ...
func (g *Generator) AddField(name string, fieldType string, t antlr.Token) error {
	g.addField(g.namespaces, name, t)
	s := g.curObj()
	fieldGen, ok := map[string]func(name string){
		"int":     s.AddInt,
		"int8":    s.AddInt8,
		"int16":   s.AddInt16,
		"int32":   s.AddInt32,
		"int64":   s.AddInt64,
		"uint":    s.AddUint,
		"uint8":   s.AddUint8,
		"uint16":  s.AddUint16,
		"uint32":  s.AddUint32,
		"uint64":  s.AddUint64,
		"hex":     s.AddUint,
		"hex8":    s.AddUint8,
		"hex16":   s.AddUint16,
		"hex32":   s.AddUint32,
		"hex64":   s.AddUint64,
		"oct":     s.AddUint,
		"oct8":    s.AddUint8,
		"oct16":   s.AddUint16,
		"oct32":   s.AddUint32,
		"oct64":   s.AddUint64,
		"dec32":   s.AddInt32,
		"dec64":   s.AddInt64,
		"dec128":  s.AddDec128,
		"float32": s.AddFloat32,
		"float64": s.AddFloat64,
		"string":  s.AddString,
		"str":     s.AddStr,
	}[fieldType]
	if !ok {
		fieldNames := []string{
			"int8", "int16", "int32", "int64",
			"uint8", "uint16", "uint32", "uint64",
			"decX.Y",
			"float32", "float64",
			"string", "str",
		}
		for i, fieldName := range fieldNames {
			fieldNames[i] = fmt.Sprintf("\033[1m%s\033[0m", fieldName)
		}
		return g.ErrorToken(t, "unsupported type `\033[1m%s\033[0m`, must be one of %s",
			fieldType, strings.Join(fieldNames, ", "))
	}
	fieldGen(name)
	return nil
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

func (g *Generator) RestLengthCheck(operator string, length int) error {
	var operatorAction func(srcobj.Source, srcobj.Source) srcobj.Source
	var errorFormat string
	switch operator {
	case "<":
		operatorAction = srcobj.OperatorGE
		errorFormat = "rest is longer than required (%d symbols)"
	case "==":
		operatorAction = srcobj.OperatorNEq
		errorFormat = "rest is not %d symbols long"
	case ">":
		operatorAction = srcobj.OperatorLE
		errorFormat = "rest is shorter than required (%d symbols)"
	}
	g.body.Append(
		srcobj.If{
			Expr: operatorAction(
				srcobj.NewCall("len", g.rest()),
				srcobj.Literal(length),
			),
			Then: g.failure(
				errorFormat,
				srcobj.Literal(length),
			),
		},
	)
	return nil
}

// PassN passes first N characters if they are there, otherwise signal a error
func (g *Generator) PassN(n int) error {
	g.body.Append(srcobj.Literal("\n"))
	g.body.Append(srcobj.Comment(fmt.Sprintf("Pass first N symbols in the rest")))
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
				"cannot skip first %d symbols: only %d left in the rest",
				srcobj.Literal(n),
				srcobj.NewCall("len", g.rest()),
			),
		},
	)
	return nil
}

func (g *Generator) PassHeadCharacters(char string) error {
	const (
		counter = "headPassCounter"
		value   = "headPassValue"
	)
	if err := g.regVar(counter, "int"); err != nil {
		return err
	}
	itemType := "byte"
	if g.useString {
		itemType = "rune"
	}
	if err := g.regVar(value, itemType); err != nil {
		return err
	}
	g.body.Append(srcobj.Literal("\n"))
	g.body.Append(srcobj.Comment(fmt.Sprintf("Pass all characters %s at the rest start", char)))
	if g.rulePassCounter > 0 {
		g.body.Append(srcobj.Assign(counter, srcobj.Literal(0)))
		g.body.Append(srcobj.Literal("\n"))
	}
	g.rulePassCounter++
	rest := g.rest()
	if g.useString {
		rest = srcobj.NewCall("string", rest)
	}
	g.body.Append(
		srcobj.For{
			I:         counter,
			Value:     value,
			Container: rest,
			Body: srcobj.If{
				Expr: srcobj.OperatorNEq(srcobj.Raw(value), srcobj.Raw(char)),
				Then: srcobj.Break,
			},
			DontAssign: true,
		},
	)
	g.body.Append(srcobj.If{
		Expr: srcobj.OperatorGT(srcobj.Raw(counter), srcobj.Literal(0)),
		Then: srcobj.OperatorAssign(g.rest(), srcobj.SliceFrom(g.rest(), srcobj.Raw(counter))),
	})
	return nil
}

// Stress mismatches should be treated as critical errors
func (g *Generator) Stress() error {
	g.critical = true
	return nil
}

// Relax ...
func (g *Generator) Relax() error {
	g.critical = false
	return nil
}

// Generate writes into io.Writer
func (g *Generator) Generate(pkgName string, dest io.Writer) error {
	g.file.PkgName(pkgName)
	if err := g.file.Dump(dest); err != nil {
		return err
	}
	return nil
}

// Push pushes data
func (g *Generator) Push() error {
	if len(g.ruleName) == 0 {
		return fmt.Errorf("no rule has been set up to push it now")
	}

	g.indent()
	g.body.Append(srcobj.ReturnOK)
	g.file.Append(g.optgetters)

	if err := g.Relax(); err != nil {
		return err
	}
	g.critical = false
	g.vars = map[string]string{}
	g.fields = map[string]Name{}
	g.scopeAbandoned = map[string]bool{}
	g.ruleName = ""
	return nil
}

// RegGravity registers center of gravity
func (g *Generator) RegGravity(name string) error {
	g.gravity = append(g.gravity, name)
	return nil
}

// AtEnd checks if the rest is empty
func (g *Generator) AtEnd() error {
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
	return nil
}
