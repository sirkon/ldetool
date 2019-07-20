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
		Type    string `arg:"--type,required" help:"field implementation name"`
		Name    string `arg:"--name,required" help:"ldetool field type"`
		Handler string `arg:"--handler" help:"field registrator function name, equals to Add<--type> by default"`
		Native  string `arg:"--native" help:"what Go type to use to back it in the Go code, <--name> to be used by default"`
	}
	arg.MustParse(&args)

	if len(args.Native) == 0 {
		args.Native = args.Name
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

// Native returns Go's representation of this field's type
func (i {{ .Type }}) Native() string {
    return "{{ .Native }}"
}

func init() {
    if builtins == nil {
		builtins = map[string]func(name string)Field{}
	}
    builtins["{{ .Name }}"] = func(fieldName string) Field {
        return {{ .Type }}(fieldName) 
    }
}
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
