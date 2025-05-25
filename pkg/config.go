// Package pkg contains the core logic for Hekura,
// including configuration loading, manifest building, and diffing.
package pkg

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Helmfile    []string `yaml:"helmfile"`
	Kustomize   []string `yaml:"kustomize"`
	RawManifest []string `yaml:"raw-manifests"`
}

// LoadConfig reads a YAML configuration file specified by configFileName
// and unmarshals it into the provided Config struct.
// It returns an error if the file cannot be read or if unmarshalling fails.
func LoadConfig(configFileName string, config *Config) error {
	yamlFile, err := os.ReadFile(configFileName)
	if err != nil {
		return fmt.Errorf("error reading YAML file %s: %w", configFileName, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML from %s: %w", configFileName, err)
	}
	return nil
}
