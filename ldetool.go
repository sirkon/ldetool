package main

//go:generate gocc lde_grammar.bnf
//go:generate ldetool sync --code-source generator/gogen/template_data --package main static_template_data.go

import (
	"os"

	"github.com/urfave/cli"

	// These are for testing reasons
	_ "github.com/sirkon/ldetool/errors"
	_ "github.com/sirkon/ldetool/lexer"
	_ "github.com/sirkon/ldetool/parser"
	_ "github.com/sirkon/ldetool/token"
	_ "github.com/sirkon/ldetool/util"
	"github.com/sirkon/message"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Text data extraction Go source code generator"
	app.UsageText = "ldetool [--code-source <path:>]"
	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "translate text data extraction rules into source code",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "code-source",
					Value: "",
					Usage: "path for templates folder, template cache will use file system if set",
				},
				cli.StringFlag{
					Name:  "yaml-dict",
					Value: "",
					Usage: "YAML-formatted gotifying dictionary",
				},
				cli.StringFlag{
					Name:  "json-dict",
					Value: "",
					Usage: "JSON-formatted gotifying dictionary",
				},
				cli.StringFlag{
					Name:  "package",
					Value: "",
					Usage: "Package name for generated files",
				},
			},
			UsageText: "ldetool generate [command options] <lde file path>",
			Action:    generateAction,
			Before: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.NewExitError("There must be one and only one lde rule file", 1)
				}
				if len(c.String("yaml-dict")) > 0 && len(c.String("json-dict")) > 0 {
					return cli.NewExitError(
						"Cannot use yaml and json formatting dictionaries at the same time, choose one of them", 1)
				}
				if len(c.String("package")) == 0 {
					return cli.NewExitError("Package name is required", 1)
				}
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		message.Critical(err)
	}
}
