package main

//go:generate gocc ldegen_grammar.bnf

import (
	"os"

	"github.com/urfave/cli"

	// These are for testing reasons
	_ "github.com/glossina/ldegen/errors"
	_ "github.com/glossina/ldegen/lexer"
	_ "github.com/glossina/ldegen/parser"
	_ "github.com/glossina/ldegen/token"
	_ "github.com/glossina/ldegen/util"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Text data extraction Go source code generator"
	app.UsageText = "ldegen [--code-source <path:>]"
	app.Commands = []cli.Command{
		cli.Command{
			Name:  "generate",
			Usage: "translate text data extraction rules into source code",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "code-source",
					Value: "",
					Usage: "path for templates folder, the template cache will use file system if set",
				},
			},
			Action: generateAction,
		},
	}
	app.Run(os.Args)
}
