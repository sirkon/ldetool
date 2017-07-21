package main

import (
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/glossina/ldetool/templatecache"
	"github.com/glossina/message"
	"github.com/urfave/cli"
)

const (
	staticTemplateProgram = `
package {{.PackageName}}

var staticTemplatesData = map[string]string{
{{ range $key, $value := .Templates }}"{{$key}}": "{{ call $.Echo $value }}",
{{end}}
}

`
)

type templateParams struct {
	PackageName string
	Templates   map[string]string
	Echo        func(string) string
}

func echo(data string) string {
	replaceMap := map[string]string{
		"\t":    "\\t",
		"\n":    "\\n",
		"\r":    "\\r",
		"\a":    "\\a",
		"\b":    "\\b",
		"\"":    "\\\"",
		"\\033": "\\\\033",
	}
	for orig, replacement := range replaceMap {
		data = strings.Replace(data, orig, replacement, -1)
	}
	return data
}

func syncAction(c *cli.Context) error {
	path := c.String("code-source")
	tc := templatecache.NewFS(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		message.Critical(err)
	}

	for _, file := range files {
		_, name := filepath.Split(file.Name())
		_, err := tc.Get(name)
		if err != nil {
			message.Critical(err)
		}
	}

	t := template.New("tsdfsdf")
	tmpl, err := t.Parse(staticTemplateProgram)
	if err != nil {
		message.Critical(err)
	}

	var dest io.WriteCloser
	if c.NArg() == 1 {
		dest, err = os.Create(c.Args()[0])
		if err != nil {
			message.Critical(err)
		}
	} else {
		dest = os.Stdout
	}
	defer dest.Close()

	raw := tc.RawData()

	err = tmpl.Execute(dest, templateParams{
		PackageName: c.String("package"),
		Templates:   raw,
		Echo:        echo,
	})
	if err != nil {
		message.Critical(err)
	}
	return nil
}
