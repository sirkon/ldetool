package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/glossina/gotify"
	"github.com/glossina/ldetool/ast"
	"github.com/glossina/ldetool/builder"
	"github.com/glossina/ldetool/generator/gogen"
	"github.com/glossina/ldetool/lexer"
	"github.com/glossina/ldetool/parser"
	"github.com/glossina/ldetool/templatecache"
)

func parse(input string) (res ast.RuleItem, err error) {
	lex := lexer.NewLexer([]byte(input))
	p := parser.NewParser()
	w, err := p.Parse(lex)
	if err != nil {
		panic(err)
	}
	rules, ok := w.([]ast.RuleItem)
	if !ok {
		panic(fmt.Errorf("Not a set of rules"))
	}
	if len(rules) != 1 {
		return res, fmt.Errorf("We only consume 1 rule per input in testing")
	}
	return rules[0], nil
}

func compile(input string) (string, error) {
	res, err := parse(input)
	if err != nil {
		return "", err
	}
	gopath := os.Getenv("GOPATH")
	root := filepath.Join(gopath, "src/github.com/glossina/ldetool/generator/gogen/template_data")
	tc := templatecache.NewFS(root)
	gen := gogen.NewGenerator(gotify.New(nil), tc)
	buf := &bytes.Buffer{}
	b := builder.NewBuilder("main", gen, buf)
	err = b.BuildRule(res)
	if err != nil {
		return "", err
	}
	err = b.Build()
	if err != nil {
		return "", err
	}
	return buf.String(), err
}

// test lookups
func TestLookupGeneration(t *testing.T) {
	data, err := compile(`rule = _ "bound";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ "bound";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _"bound"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _"bound"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _"bound"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _"bound"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _'@';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _'@';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _'@'[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _'@'[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _'@'[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _'@'[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

// test lookups
func TestOptionalLookupGeneration(t *testing.T) {
	data, err := compile(`rule = ?Opt ( _ "bound" );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _ "bound" );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( _"bound"[:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _"bound"[:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( _"bound"[12:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _"bound"[12:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( _'@' );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _'@' );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( _'@'[:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _'@'[:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( _'@'[12:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ?Opt ( ! _'@'[12:40] );`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}

// test lookups with optional stay here decision
func TestIgnorableLookupGeneration(t *testing.T) {
	data, err := compile(`rule = _ ?? "bound";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? "bound";`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _ ?? "bound"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? "bound"[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _ ?? "bound"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? "bound"[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _ ?? '@';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? '@';`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _ ?? '@'[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? '@'[:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = _ ?? '@'[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)

	data, err = compile(`rule = ! _ ?? '@'[12:40];`)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(data)
}
