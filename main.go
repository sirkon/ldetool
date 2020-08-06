package main

import (
	"bytes"
	"errors"
	"fmt"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/sirkon/gotify"
	"github.com/urfave/cli"

	"github.com/sirkon/ldetool/internal/ast"
	"github.com/sirkon/ldetool/internal/generator"
	"github.com/sirkon/ldetool/internal/generator/gogen"
	"github.com/sirkon/ldetool/internal/listener"
	parser2 "github.com/sirkon/ldetool/internal/parser"
	"github.com/sirkon/ldetool/internal/srcbuilder"
	"github.com/sirkon/ldetool/internal/types"

	// These are for testing reasons
	_ "github.com/sirkon/ldetool/internal/parser"

	"github.com/sirkon/message"
)

type runConfig struct {
	YAMLDict     string `arg:"--yaml-dict" help:"YAML-formatted gotifying dictionary"`
	JSONDict     string `arg:"--json-dict" help:"JSON-formatted gotifying dictionary"`
	Package      string `arg:"--package" help:"package name for generated files"`
	BigEndian    bool   `arg:"--big-endian" help:"target architecture is big endian"`
	LittleEndian bool   `arg:"--little-endian" help:"target architecture is little endian"`
	GoString     bool   `arg:"--go-string" help:"treat strings as go string"`
	Version      bool   `arg:"-v" help:"prints version and stop"`

	File []string `arg:"positional" help:"file to process"`
}

func main() {
	var cfg runConfig
	p := arg.MustParse(&cfg)

	if len(cfg.JSONDict) > 0 && len(cfg.YAMLDict) > 0 {
		p.Fail("--yaml-dict and --json-dict are mutually exclusive")
	}
	if cfg.BigEndian && cfg.LittleEndian {
		p.Fail("--big-endian and --little-endian are mutually exclusive")
	}
	if len(cfg.File) >= 1 {
		if cfg.File[0] == "generate" {
			message.Warningf("subcommand `generate` of ldetool is considered abundant. It is deprecated, please remove it")
			for i, item := range cfg.File[1:] {
				cfg.File[i] = item
			}
			cfg.File = cfg.File[:len(cfg.File)-1]
		}
	}
	if cfg.Version {
		var version string
		info, ok := debug.ReadBuildInfo()
		if !ok {
			version = "(devel)"
		} else {
			version = info.Main.Version
		}

		message.Info("ldetool", "version", version)
		return
	}
	switch len(cfg.File) {
	case 0:
		p.Fail("missing file name with LDE rules")
	case 1:
	default:
		p.Fail(fmt.Sprintf("ldetool only take 1 file, got %d", len(cfg.File)))
	}

	resolvePackageName(p, &cfg)
	if err := generate(&cfg); err != nil {
		message.Error(err)
	}
}

func resolvePackageName(p *arg.Parser, c *runConfig) {
	outputFileName := getOutputFileName(c.File[0])
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
		if len(c.Package) == 0 {
			p.Fail("no existing Go files found in the directory, option --package must be set")
		}
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
	var pkgNames []string
	for pkgName := range pkgs {
		pkgNames = append(pkgNames, pkgName)
	}
	sort.Strings(pkgNames)
	if len(pkgs) >= 2 && pkgNames[1] != pkgNames[0]+"_test" {
		message.Fatalf("located %d packages in `%s`, thus cannot set up a package name", len(pkgs), dirPath)
	}
	if len(pkgs) == 0 {
		return
	}
	if len(c.Package) == 0 {
		c.Package = pkgNames[0]
	}
	for _, pkgName := range pkgNames {
		if c.Package == pkgName || c.Package == pkgName+"_test" {
			return
		}
	}
	message.Fatalf("invalid package name `%s` set for code: there's package name `%s` in the directory", c.Package, pkgNames[0])
}

func getOutputFileName(ruleFile string) string {
	dirPath, fname := filepath.Split(ruleFile)
	if strings.HasSuffix(fname, ".lde") {
		fname = fname[:len(fname)-4]
	}
	fname = fmt.Sprintf("%s_lde.go", strings.Replace(fname, ".", "_", -1))
	return path.Join(dirPath, fname)
}

func generate(c *runConfig) (err error) {
	et := NewErrorTranslator()

	var errorToken antlr.Token
	defer func() {
		if r := recover(); r != nil {
			switch v := r.(type) {
			case *ast.ErrorListener:
				err = fmt.Errorf("%d:%d: %s", v.Line, v.Col+1, v.Msg)
			case string:
				err = errors.New(v)
			default:
				panic(r)
			}
		}
		if err != nil {
			if errorToken != nil {
				err = cli.NewExitError(
					fmt.Sprintf(
						"%s:%d:%d: %s",
						c.File[0],
						errorToken.GetLine(),
						errorToken.GetColumn(),
						err),
					1,
				)
			} else {
				err = et.Translate(err)
				err = cli.NewExitError(fmt.Sprintf("%s:%s", c.File[0], err), 1)
			}
		}
	}()

	ruleFileName := c.File[0]
	input, err := antlr.NewFileStream(ruleFileName)
	if err != nil {
		return
	}
	lexer := parser2.NewLDELexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewLDEParser(stream)
	p.RemoveErrorListeners()
	el := &ast.ErrorListener{}
	p.AddErrorListener(el)
	tree := p.Rules()
	walker := antlr.NewParseTreeWalker()
	l := listener.New(lexer.Comments())

	eh := antlr.NewBailErrorStrategy()
	p.RemoveErrorListeners()
	p.SetErrorHandler(eh)

	walker.Walk(l, tree)

	rules := l.Rules()
	formatDict := getDict(c)

	tmpDest := &bytes.Buffer{}
	gfy := gotify.New(formatDict)
	externalTypes := l.Types().Types()
	if externalTypes == nil {
		externalTypes = map[string]types.TypeRegistration{}
	}
	gen := gogen.NewGenerator(c.GoString, gfy, externalTypes)
	if c.LittleEndian {
		gen.PlatformType(generator.LittleEndian)
	} else if c.BigEndian {
		gen.PlatformType(generator.BigEndian)
	}
	b := srcbuilder.New(c.Package, gen, tmpDest, gfy)
	b.DontRecover()
	if err := b.DispatchTypeRegistration(l.Types()); err != nil {
		return err
	}
	for _, rule := range rules {
		if gfy.Public(rule.Name) != rule.Name {
			errorToken = rule.NameToken
			return fmt.Errorf("wrong rule name %s, must be %s", rule.Name, gfy.Public(rule.Name))
		}
		message.Infof("\nRule `\033[1m%s\033[0m`: processing", rule.Name)
		err = b.BuildRule(rule)
		if err != nil {
			errorToken = b.ErrorToken()
			return err
		}
		message.Infof("Rule `\033[1m%s\033[0m`: done", rule.Name)
		gen.Relax()
	}
	if err = b.Build(); err != nil {
		return
	}

	destFileName := getOutputFileName(ruleFileName)
	dest, err := os.Create(destFileName)
	if err != nil {
		message.Fatal(err)
	}
	defer func() {
		if nerr := dest.Close(); nerr != nil {
			message.Error(nerr)
		}
		if err != nil {
			if nerr := os.Remove(destFileName); nerr != nil {
				message.Warning(nerr)
			}
		}
	}()
	if _, err := io.Copy(dest, tmpDest); err != nil {
		message.Fatal("generation error: %s", err)
	}

	return
}
