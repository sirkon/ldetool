package main

//go:generate gocc lde_grammar.bnf

import (
	"os"

	"github.com/urfave/cli"

	// These are for testing reasons
	_ "github.com/glossina/ldetool/errors"
	_ "github.com/glossina/ldetool/lexer"
	_ "github.com/glossina/ldetool/parser"
	_ "github.com/glossina/ldetool/token"
	_ "github.com/glossina/ldetool/util"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Text data extraction Go source code generator"
	app.UsageText = "ldetool [--code-source <path:>]"
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
