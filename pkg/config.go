package pkg

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Helmfile    []string `yaml:"helmfile"`
	Kustomize   []string `yaml:"kustomize"`
	RawManifest []string `yaml:"raw-manifests"`
}

func LoadConfig(configFileName string) Config {
	yamlFile, err := os.ReadFile(configFileName)
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
