package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"bytes"

	"io"

	"path/filepath"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/builder"
	"github.com/sirkon/ldetool/generator"
	"github.com/sirkon/ldetool/generator/gogen"
	"github.com/sirkon/ldetool/listener"
	"github.com/sirkon/ldetool/parser"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) (err error) {

	et := NewErrorTranslator()

	var errorToken antlr.Token
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case *ErrorListener:
				err = fmt.Errorf("%d:%d: %s", v.line, v.col, v.msg)
			case string:
				err = errors.New(v)
			default:
				panic(err)
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

	fileName := c.Args()[0]
	input, err := antlr.NewFileStream(fileName)
	if err != nil {
		return
	}
	lexer := parser.NewLDELexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewLDEParser(stream)
	p.RemoveErrorListeners()
	el := &ErrorListener{}
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

	if strings.HasSuffix(fileName, ".lde") {
		fileName = fileName[:len(fileName)-4]
	}
	dirPath, fname := filepath.Split(fileName)
	fname = fmt.Sprintf("%s_lde.go", strings.Replace(fname, ".", "_", -1))
	tmpDest := &bytes.Buffer{}
	gfy := gotify.New(formatDict)
	gen := gogen.NewGenerator(c.Bool("go-string"), gfy)
	if c.Bool("little-endian") {
		gen.PlatformType(generator.LittleEndian)
	} else if c.Bool("big-endian") {
		gen.PlatformType(generator.BigEndian)
	}
	b := builder.NewBuilder(c.String("package"), gen, tmpDest, gfy)
	b.DontRecover()
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

	dest, err := os.Create(filepath.Join(dirPath, fname))
	if err != nil {
		message.Fatal(err)
	}
	defer func() {
		if nerr := dest.Close(); nerr != nil {
			message.Error(nerr)
		}
		if err != nil {
			if nerr := os.Remove(fname); nerr != nil {
				message.Warning(nerr)
			}
		}
	}()
	io.Copy(dest, tmpDest)

	return
}
