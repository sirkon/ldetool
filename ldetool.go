package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/urfave/cli"

	// These are for testing reasons
	_ "github.com/sirkon/ldetool/internal/parser"

	"github.com/sirkon/message"
)

func main() {
	app := cli.NewApp()
	app.Usage = "Text data extraction Go source code generator"
	app.Version = ldetoolVersion
	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "translate text data extraction rules into source code",
			Flags: []cli.Flag{
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
				cli.BoolFlag{
					Name:  "big-endian",
					Usage: "Target architecture is big endian",
				},
				cli.BoolFlag{
					Name:  "little-endian",
					Usage: "Target architecture is little endian",
				},
				cli.BoolFlag{
					Name:  "go-string",
					Usage: "Treat strings as go strings",
				},
			},
			UsageText: "ldetool generate [command options] <lde file path>",
			Action:    generate,
			Before: func(c *cli.Context) error {
				if c.NArg() != 1 {
					return cli.NewExitError("There must be one and only one lde rule file", 1)
				}
				c.Args().First()
				if len(c.String("yaml-dict")) > 0 && len(c.String("json-dict")) > 0 {
					return cli.NewExitError(
						"Cannot use yaml and json formatting dictionaries at the same time, choose one of them", 1)
				}
				resolvePackageName(c)
				if len(c.String("package")) == 0 {
					return cli.NewExitError("Package name is required", 1)
				}
				if c.Bool("big-endian") && c.Bool("little-endian") {
					return cli.NewExitError("Target architecture cannot be both little and big endian", 1)
				}
				return nil
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		message.Critical(err)
	}
}

func resolvePackageName(c *cli.Context) {
	outputFileName := getOutputFileName(c.Args().First())
	dirPath, outFileBaseName := path.Split(outputFileName)
	if len(dirPath) == 0 {
		dirPath = "."
	}
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		message.Fatalf("failed to read `%s`: %s", dirPath, err)
	}

	var packageDetected bool
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") || file.Name() != outFileBaseName {
			packageDetected = true
			break
		}
	}
	if !packageDetected {
		return
	}

	// parse package to extract package name
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dirPath, func(info os.FileInfo) bool {
		return info.Name() != outFileBaseName
	}, parser.PackageClauseOnly)
	if err != nil {
		message.Warningf("failed to parse one of Go source files, thus cannot resolve package name: %s", err)
		return
	}
	if len(pkgs) > 1 {
		message.Warningf("located %d packages in `%s`, thus cannot set up a package name", len(pkgs), dirPath)
	}
	if len(pkgs) == 0 {
		return
	}
	for pkgName := range pkgs {
		if err := c.Set("package", pkgName); err != nil {
			message.Warningf("failed to set up package name: %s", err)
		}
		return
	}
}
