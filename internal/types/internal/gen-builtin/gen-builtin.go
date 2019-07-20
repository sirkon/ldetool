package main

import (
	"bytes"
	"os"
	"text/template"

	"github.com/alexflint/go-arg"
	"github.com/sirkon/gosrcfmt"
	"github.com/sirkon/message"
)

func main() {
	var args struct {
		Type       string `arg:"--type,required" help:"field implementation name"`
		Name       string `arg:"--name,required" help:"ldetool field type"`
		Handler    string `arg:"--handler" help:"field registrator function name, equals to Add<--type> by default"`
		GoName     string `arg:"--go-name" help:"what Go type to use to back it in the Go code, <--name> to be used by default"`
		Native     bool   `arg:"--native" help:"set on if this type had direct equivalent in Go type system"`
		Declarable bool   `arg:"--declarable" help:"set on when this type can be used for direct declaration"`
		Decimal    bool   `arg:"--decimal" help:"set on when type is one of decimlas"`
	}
	p := arg.MustParse(&args)

	if len(args.GoName) == 0 {
		if !args.Native {
			p.Fail("either --native or --go-name must be set")
		}
		args.GoName = args.Name
	} else {
		if args.Native {
			p.Fail("--native or --go-name are mutually exclusive, only of of them can be set in the same time")
		}
	}
	if len(args.Handler) == 0 {
		args.Handler = "Add" + args.Type
	}

	dest, err := os.Create("gen_" + args.Name + ".go")
	if err != nil {
		message.Fatalf("create field handler for %s: %s", args.Name, err)
	}
	defer func() {
		if err := dest.Close(); err != nil {
			message.Fatalf("closing %s: %s", dest.Name(), err)
		}
	}()

	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, args); err != nil {
		message.Fatalf("generating %s: %s", dest.Name(), err)
	}

	gosrcfmt.Format(dest, buf.Bytes())
}

const (
	data = `package types

var _ Field = {{ .Type }}("")

// {{ .Type }} represents field of type {{ .Name }}
type {{ .Type }} string

// Name returns field name
func (i {{ .Type }}) Name() string {
	return string(i)
}

// TypeName name of the type
func (i {{ .Type }}) TypeName() string {
	return "{{ .Name }}"
}

// Register registers a field
func (i {{ .Type }}) Register(registrator FieldRegistrator) {
	registrator.{{ .Handler }}(i.Name())
}

// GoName returns Go's representation of this field's type
func (i {{ .Type }}) GoName() string {
    return "{{ .GoName }}"
}

func init() {
    if builtins == nil {
		builtins = map[string]func(name string)Field{}
	}
    builtins["{{ .Name }}"] = func(fieldName string) Field {
        return {{ .Type }}(fieldName) 
    }
    {{ if .Native }}if natives == nil {
        natives = map[string]struct{}{}
    }
    natives["{{ .Name }}"] = struct{}{}{{ end }}
    {{ if .Declarable }}if declarables == nil {
        declarables = map[string]struct{}{}
    }
    declarables["{{ .Name }}"] = struct{}{}{{ end }}
    {{ if .Decimal }}if decimals == nil {
        decimals = map[string]struct{}{}
    }
    decimals["{{ .Name }}"] = struct{}{}{{ end }}
    {{ if not .Native }}if backedBy == nil {
        backedBy = map[string]string{}
    }
    backedBy["{{ .Name }}"] = "{{ .GoName }}"{{ end }}}
`
)

var tmpl *template.Template

func init() {
	t := template.New("template")
	var err error
	tmpl, err = t.Parse(data)
	if err != nil {
		panic(err)
	}
}
