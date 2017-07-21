package main

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/glossina/ldetool/templatecache"
	"github.com/stretchr/testify/require"
)

func TestSyncing(t *testing.T) {
	path := "generator/gogen/template_data"
	tc1 := templatecache.NewFS(path)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		t.Fatal(err)
	}
	for _, file := range files {
		_, name := filepath.Split(file.Name())
		_, err := tc1.Get(name)
		if err != nil {
			t.Fatal(err)
		}
	}

	tc2 := templatecache.NewMap(staticTemplatesData)

	require.Equal(t, tc1.RawData(), tc2.RawData())
}
