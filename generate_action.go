package main

import (
	"errors"
	"fmt"
	"github.com/sirkon/ldetool/internal/ast"
	"os"
	"path"
	"strings"

	"bytes"

	"io"

	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/generator/gogen"
	"github.com/sirkon/ldetool/internal/listener"
	"github.com/sirkon/ldetool/internal/parser"
	"github.com/sirkon/ldetool/internal/srcbuilder"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
)

func getOutputFileName(ruleFile string) string {
	dirPath, fname := filepath.Split(ruleFile)
	if strings.HasSuffix(fname, ".lde") {
		fname = fname[:len(fname)-4]
	}
	fname = fmt.Sprintf("%s_lde.go", strings.Replace(fname, ".", "_", -1))
	return path.Join(dirPath, fname)
}

func generateAction(c *cli.Context) (err error) {

	et := NewErrorTranslator()

	var errorToken antlr.Token
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case *ast.ErrorListener:
				err = fmt.Errorf("%d:%d: %s", v.Line, v.Col+1, v.Msg)
			case string:
				err = errors.New(v)
			default:
				panic(r)
			}
		}
		if err != nil {
			if errorToken != nil {
				err = cli.NewExitError(
					fmt.Sprintf(
						"%s:%d:%d: %s",
						c.Args()[0],
						errorToken.GetLine(),
						errorToken.GetColumn(),
						err),
					1,
				)
			} else {
				err = et.Translate(err)
				err = cli.NewExitError(fmt.Sprintf("%s:%s", c.Args()[0], err), 1)
			}
		}
	}()

	ruleFileName := c.Args().First()
	input, err := antlr.NewFileStream(ruleFileName)
	if err != nil {
		return
	}
	lexer := parser.NewLDELexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLDEParser(stream)
	p.RemoveErrorListeners()
	el := &ast.ErrorListener{}
	p.AddErrorListener(el)
	tree := p.Rules()
	walker := antlr.NewParseTreeWalker()
	l := listener.New()

	eh := antlr.NewBailErrorStrategy()
	p.RemoveErrorListeners()
	p.SetErrorHandler(eh)

	walker.Walk(l, tree)

	rules := l.Rules()
	formatDict := getDict(c)

	tmpDest := &bytes.Buffer{}
	gfy := gotify.New(formatDict)
	gen := gogen.NewGenerator(c.Bool("go-string"), gfy, l.Types().Types())
	if c.Bool("little-endian") {
		gen.PlatformType(generator.LittleEndian)
	} else if c.Bool("big-endian") {
		gen.PlatformType(generator.BigEndian)
	}
	b := srcbuilder.New(c.String("package"), gen, tmpDest, gfy)
	b.DontRecover()
	if err := b.DispatchTypeRegistration(l.Types()); err != nil {
		return err
	}
	for _, rule := range rules {
		if gfy.Public(rule.Name) != rule.Name {
			errorToken = rule.NameToken
			return fmt.Errorf("wrong rule name %s, must be %s", rule.Name, gfy.Public(rule.Name))
		}
		message.Infof("\nRule `\033[1m%s\033[0m`: processing", rule.Name)
		err = b.BuildRule(rule)
		if err != nil {
			errorToken = b.ErrorToken()
			return err
		}
		message.Infof("Rule `\033[1m%s\033[0m`: done", rule.Name)
		gen.Relax()
	}
	if err = b.Build(); err != nil {
		return
	}

	destFileName := getOutputFileName(ruleFileName)
	dest, err := os.Create(destFileName)
	if err != nil {
		message.Fatal(err)
	}
	defer func() {
		if nerr := dest.Close(); nerr != nil {
			message.Error(nerr)
		}
		if err != nil {
			if nerr := os.Remove(destFileName); nerr != nil {
				message.Warning(nerr)
			}
		}
	}()
	io.Copy(dest, tmpDest)

	return
}
