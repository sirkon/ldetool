package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/sirkon/gotify"
	"github.com/sirkon/ldetool/ast"
	"github.com/sirkon/ldetool/builder"
	"github.com/sirkon/ldetool/generator/gogen"
	"github.com/sirkon/ldetool/lexer"
	"github.com/sirkon/ldetool/parser"
	"github.com/sirkon/ldetool/templatecache"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) (err error) {
	path := c.String("code-source")
	var tc *templatecache.TemplateCache
	if len(path) != 0 {
		tc = templatecache.NewFS(path)
	} else {
		tc = templatecache.NewMap(staticTemplatesData)
	}

	inputSource := c.Args()[0]
	input, err := ioutil.ReadFile(inputSource)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	defer func() {
		if err != nil {
			err = cli.NewExitError(fmt.Sprintf("%s: %s", c.Args()[0], err), 1)
		}
	}()
	lex := lexer.NewLexer(input)
	p := parser.NewParser()
	w, err := p.Parse(lex)
	if err != nil {
		return
	}
	rules, ok := w.([]ast.RuleItem)
	if !ok {
		return fmt.Errorf("Not an parsing scripts file")
	}
	formatDict := getDict(c)

	if strings.HasSuffix(inputSource, ".lde") {
		inputSource = inputSource[:len(inputSource)-4]
	}
	fname := fmt.Sprintf("%s_lde.go", strings.Replace(inputSource, ".", "_", -1))
	dest, err := os.Create(fname)
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
	gen := gogen.NewGenerator(gotify.New(formatDict), tc)
	b := builder.NewBuilder(c.String("package"), gen, dest)
	for _, rule := range rules {
		err = b.BuildRule(rule)
		if err != nil {
			return err
		}
		if err != nil {
			return err
		}
		message.Infof("Rule `\033[1m%s\033[0m` translated", rule.Name)
	}
	err = b.Build()
	return
}
