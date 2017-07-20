package main

import (
	"fmt"

	"github.com/glossina/ldetool/templatecache"
	"github.com/urfave/cli"
)

func generateAction(c *cli.Context) error {
	path := c.String("code-source")
	var tc *templatecache.TemplateCache
	if len(path) != 0 {
		tc = templatecache.NewFS(path)
	} else {
		return fmt.Errorf("Dumped cache usage is not implemented yet, use --codes-source parameter")
	}
	if _, err := tc.Get("close_option"); err != nil {
		return err
	}
	return nil
}
