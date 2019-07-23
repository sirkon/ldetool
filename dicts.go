package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/go-yaml/yaml"
	"github.com/sirkon/message"
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

func getDict(c *runConfig) (res map[string]string) {
	yamlDict := c.YAMLDict
	jsonDict := c.JSONDict
	if len(yamlDict) > 0 {
		res = yamlSource(yamlDict)
	}
	if len(jsonDict) > 0 {
		res = jsonSource(jsonDict)
	}
	return
}
