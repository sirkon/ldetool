package srcbuilder

import (
	"fmt"
	"github.com/sirkon/message"
	"io"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/internal/ast"
	"github.com/sirkon/ldetool/internal/types"

	"github.com/sirkon/ldetool/internal/generator"
)

var _ ast.ActionDispatcher = &SrcBuilder{}

// SrcBuilder creates target sources using Generator object
type SrcBuilder struct {
	pkgName         string
	gen             generator.Generator
	dest            io.Writer
	recoverPanic    bool
	gotify          *gotify.Gotify
	generators      []func() error
	prefixList      []string
	registeredTypes map[string]types.TypeRegistration

	anonDepth int

	errToken antlr.Token
}

// New constructor for source builder
func New(pn string, g generator.Generator, d io.Writer, gfy *gotify.Gotify) *SrcBuilder {
	return &SrcBuilder{
		pkgName:      pn,
		gen:          g,
		dest:         d,
		recoverPanic: true,
		gotify:       gfy,
	}
}

func (sb *SrcBuilder) prefixTie(item string) {
	sb.prefixList = append(sb.prefixList, item)
}

func (sb *SrcBuilder) prefixUntie() {
	sb.prefixList = sb.prefixList[:len(sb.prefixList)-1]
}

func (sb *SrcBuilder) prefixCur() string {
	return strings.Join(sb.prefixList, ".")
}

// DontRecover tells not to recover panics
func (sb *SrcBuilder) DontRecover() {
	sb.recoverPanic = false
}

// BuildRule builds code from the data
func (sb *SrcBuilder) BuildRule(rule *ast.Rule) (err error) {
	// fill register types
	if err := rule.Accept(sb); err != nil {
		return err
	}
	generators := sb.generators
	sb.generators = nil
	for _, item := range generators {
		if err := item(); err != nil {
			return err
		}
	}
	if err := sb.gen.Push(); err != nil {
		return err
	}
	return nil
}

// Build full source file
func (sb *SrcBuilder) Build() (err error) {
	sb.gen.Generate(sb.pkgName, sb.dest)
	return nil
}

// checkField if field name is goish public variable
func (sb *SrcBuilder) checkField(field ast.Field) error {
	if sb.gotify.Public(field.Name) != field.Name {
		sb.errToken = field.NameToken
		return fmt.Errorf("Wrong taker identifier `%s`, must be %s", field.Name, sb.gotify.Public(field.Name))
	}
	return nil
}

func (sb *SrcBuilder) composeRule(actions []ast.Action) error {
	for _, action := range actions {
		message.Info(action)
		if err := action.Accept(sb); err != nil {
			return fmt.Errorf("%s: %s", action, err)
		}
	}
	return nil
}

func (sb *SrcBuilder) appendGens(gens ...func() error) {
	sb.generators = append(sb.generators, gens...)
}

// errorToken returns token where error happened
func (sb *SrcBuilder) ErrorToken() antlr.Token {
	return sb.errToken
}
