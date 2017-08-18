package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/sirkon/message"
	"github.com/urfave/cli"
)

func yamlSource(path string) map[string]string {
	res := map[string]string{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		message.Criticalf("Cannot read `%s`: %s", path, err)
	}
	if err := yaml.Unmarshal(data, &res); err != nil {
		message.Criticalf("Cannot parse `%s` as YAML file: %s", path, err)
	}
	return res
}

func jsonSource(path string) map[string]string {
	res := map[string]string{}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		message.Criticalf("Cannot read `%s`: %s", path, err)
	}
	if err := json.Unmarshal(data, &res); err != nil {
		message.Criticalf("Cannot parse `%s` as JSON file: %s", path, err)
	}
	return res
}

func getDict(c *cli.Context) (res map[string]string) {
	yamlDict := c.String("yaml-dict")
	jsonDict := c.String("json-dict")
	if len(yamlDict) > 0 {
		res = yamlSource(yamlDict)
	}
	if len(jsonDict) > 0 {
		res = jsonSource(jsonDict)
	}
	return
}
