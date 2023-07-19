package main

import (
	"fmt"
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

func getConfig(path string) map[string]string {
	fileData, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var config Config
	err = yaml.Unmarshal(fileData, &config)
	if err != nil {
		log.Fatalf("Failed to unmarshal YAML: %v", err)
	}
	fmt.Println(config)

	out := make(map[string]string)
	for _, block := range config.Blocks {
		for _, source := range block.Sources {
			out[source] = block.Target
		}
	}

	return out
}

//func main() {
//	configPath := "/mnt/ImpSSD/Development/dupliphoto/test.yml"
//	result := getConfig(configPath)
//
//	// Print the resulting map
//	for key, value := range result {
//		fmt.Printf("Source: %s, Target: %s\n", key, value)
//	}
//}
