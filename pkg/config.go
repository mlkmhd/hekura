package pkg

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Helmfile    []string `yaml:"helmfile"`
	Kustomize   []string `yaml:"kustomize"`
	RawManifest []string `yaml:"raw-manifest"`
}

func LoadConfig(configFileName string) Config {
	yamlFile, err := ioutil.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v; %v", configFileName, err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Error unmarshalling YAML: %v", err)
	}

	// Return the configuration
	return config
}
