package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"bytes"

	"io"

	"path/filepath"

	"time"

	"bufio"

	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/ast"
	"github.com/sirkon/ldetool/builder"
	"github.com/sirkon/ldetool/generator/gogen"
	"github.com/sirkon/ldetool/lexer"
	"github.com/sirkon/ldetool/parser"
	"github.com/sirkon/ldetool/token"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) (err error) {
	inputSource := c.Args()[0]
	input, err := ioutil.ReadFile(inputSource)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	et := NewErrorTranslator()

	var errorToken *token.Token
	defer func() {
		if r := recover(); r != nil {
			if v, ok := r.(string); !ok && v != "finish" {
				panic(r)
			}
		}
		if err != nil {
			if errorToken != nil {
				err = cli.NewExitError(fmt.Sprintf("%s:%d:%d: %s", c.Args()[0], errorToken.Line, errorToken.Column, err), 1)
			} else {
				err = et.Translate(err)
				err = cli.NewExitError(fmt.Sprintf("%s:%s", c.Args()[0], err), 1)
			}
		}
	}()

	////////////////////////////
	// Unfortunately, generated parser may hang if no expected element was not found
	// Thus the workaround, where we just expect parsing to be completed within 2 seconds
	resultChan := make(chan interface{})
	go func() {
		lex := lexer.NewLexer(input)
		p := parser.NewParser()
		w, err := p.Parse(lex)
		if err != nil {
			return
		}
		resultChan <- w
	}()
	var w interface{}
	select {
	case kkk := <-resultChan:
		w = kkk
	case <-time.After(time.Second * 2):
		var count int
		var column int
		scanner := bufio.NewScanner(bytes.NewBuffer(input))
		i := 0
		for scanner.Scan() {
			if len(scanner.Bytes()) > 0 {
				count = i + 1
				column = len(scanner.Bytes()) + 1
			}
			i++
		}
		err = fmt.Errorf("Probably, `\033[1m;\033[0m` missed")
		errorToken = &token.Token{
			Pos: token.Pos{
				Line:   count,
				Column: column,
			},
		}
		return
	}
	///////////////

	rules, ok := w.([]ast.RuleItem)
	if !ok {
		return fmt.Errorf("not a parsing scripts file")
	}
	formatDict := getDict(c)

	if strings.HasSuffix(inputSource, ".lde") {
		inputSource = inputSource[:len(inputSource)-4]
	}
	dirPath, fname := filepath.Split(inputSource)
	fname = fmt.Sprintf("%s_lde.go", strings.Replace(fname, ".", "_", -1))
	tmpDest := &bytes.Buffer{}
	gfy := gotify.New(formatDict)
	gen := gogen.NewGenerator(gfy)
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
