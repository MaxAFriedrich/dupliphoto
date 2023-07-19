package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Blocks []Block `yaml:"blocks"`
}

type Block struct {
	Target  string   `yaml:"target"`
	Sources []string `yaml:"sources"`
}

func getConfig(path string) Config {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}

	return config
}
