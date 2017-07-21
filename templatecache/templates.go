package templatecache

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"text/template"
)

// TemplateCache is a template storage
type TemplateCache struct {
	root      string
	templates map[string]*template.Template
	rawdata   map[string]string
}

// NewFS TemplateStore constructor. Used for debugging purposes mainly.
// The prefered mode is to use predifined template map
func NewFS(root string) *TemplateCache {
	return &TemplateCache{
		root:      root,
		templates: map[string]*template.Template{},
		rawdata:   map[string]string{},
	}
}

// NewMap constructor
func NewMap(rawdata map[string]string) *TemplateCache {
	templates := map[string]*template.Template{}
	for k, v := range rawdata {
		to := template.New(k)
		tmpl, err := to.Parse(v)
		if err != nil {
			panic(err)
		}
		templates[k] = tmpl
	}
	return &TemplateCache{
		root:      "",
		templates: templates,
		rawdata:   rawdata,
	}
}

func (tc *TemplateCache) getPath(name string) string {
	return filepath.Join(tc.root, name)
}

// add and compile new template
func (tc *TemplateCache) add(name string) (err error) {
	if _, ok := tc.templates[name]; ok {
		panic(fmt.Errorf("Got duplicate template `\033[1m%s\033[0m", name))
	}
	path := tc.getPath(name)
	rawdata, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	t := template.New(name)
	tmpl, err := t.Parse(string(rawdata))
	if err != nil {
		return err
	}
	tc.templates[name] = tmpl
	tc.rawdata[name] = string(rawdata)
	return nil
}

// Get template by name
func (tc *TemplateCache) Get(name string) (*template.Template, error) {
	for {
		if template, ok := tc.templates[name]; !ok {
			if len(tc.root) == 0 {
				return nil, fmt.Errorf("Template `\033[1m%s\033[0m` not found", name)
			}
			if err := tc.add(name); err != nil {
				return nil, err
			}
		} else {
			return template, nil
		}
	}
}

// MustGet get template or raise panic if no such template found
func (tc *TemplateCache) MustGet(name string) *template.Template {
	if res, err := tc.Get(name); err != nil {
		panic(err)
	} else {
		return res
	}
}

// Execute named template
func (tc *TemplateCache) Execute(name string, dest io.Writer, params interface{}) error {
	res, err := tc.Get(name)
	if err != nil {
		return err
	}
	if err = res.Execute(dest, params); err != nil {
		return fmt.Errorf("%s:%s", tc.getPath(name), err)
	}
	return nil
}

// MustExecute named template
func (tc *TemplateCache) MustExecute(name string, dest io.Writer, params interface{}) {
	if err := tc.Execute(name, dest, params); err != nil {
		panic(err)
	}
}

// Dump dumps raw (text) template data as JSON serialized text into writer
func (tc *TemplateCache) Dump(dest io.Writer) error {
	enc := json.NewEncoder(dest)
	return enc.Encode(tc.rawdata)
}

// RawData returns raw template data
func (tc *TemplateCache) RawData() map[string]string {
	return tc.rawdata
}
