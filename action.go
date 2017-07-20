package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/glossina/gotify"
	"github.com/glossina/ldetool/ast"
	"github.com/glossina/ldetool/builder"
	"github.com/glossina/ldetool/generator/gogen"
	"github.com/glossina/ldetool/lexer"
	"github.com/glossina/ldetool/parser"
	"github.com/glossina/ldetool/templatecache"
	"github.com/glossina/message"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) (err error) {
	path := c.String("code-source")
	var tc *templatecache.TemplateCache
	if len(path) != 0 {
		tc = templatecache.NewFS(path)
	} else {
		return cli.NewExitError("Dumped cache usage is not implemented yet, use --codes-source parameter", 1)
	}

	input, err := ioutil.ReadFile(c.Args()[0])
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

	for _, rule := range rules {
		gen := gogen.NewGenerator(rule.Name, gotify.New(formatDict), tc)
		fname := fmt.Sprintf("%s_parser.go", rule.Name)
		dest, err := os.Create(fname)
		if err != nil {
			return err
		}
		defer func() {
			if nerr := dest.Close(); nerr != nil {
				message.Fatal(nerr)
			}
		}()
		b := builder.NewBuilder(c.String("package"), gen, dest)
		err = b.BuildRule(rule)
		if err != nil {
			return err
		}
		message.Infof("Rule `\033[1m%s\033[0m` translated into %s", rule.Name, fname)
	}

	return
}
