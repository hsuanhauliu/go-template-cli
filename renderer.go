package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type ConfigSchema struct {
	Input   string            `yaml:input`
	Output  string            `yaml:output`
	Enabled bool              `yaml:enabled`
	Mode    string            `yaml:mode`
	Values  map[string]string `yaml:values`
}

type CliConfigSchema struct {
	DefaultValues map[string]string `yaml:"values"`
	Configs       []ConfigSchema    `yaml:"configs"`
}

func main() {
	yfile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	data := CliConfigSchema{}

	err = yaml.Unmarshal(yfile, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)
}
