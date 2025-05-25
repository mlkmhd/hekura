package pkg

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestLoadConfig_Success(t *testing.T) {
	content := `
helmfile:
  - "path/to/helmfile1"
kustomize:
  - "path/to/kustomize1"
  - "path/to/kustomize2"
raw-manifests:
  - "path/to/raw-manifests1"
`
	expectedConfig := Config{
		Helmfile:    []string{"path/to/helmfile1"},
		Kustomize:   []string{"path/to/kustomize1", "path/to/kustomize2"},
		RawManifest: []string{"path/to/raw-manifests1"},
	}

	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "hekura.yaml")
	err := os.WriteFile(configFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary config file: %v", err)
	}

	var loadedConfig Config
	err = LoadConfig(configFile, &loadedConfig)
	if err != nil {
		t.Errorf("LoadConfig() error = %v, wantErr false", err)
	}

	if !reflect.DeepEqual(loadedConfig, expectedConfig) {
		t.Errorf("LoadConfig() loadedConfig = %v, want %v", loadedConfig, expectedConfig)
	}
}

func TestLoadConfig_FileNotFound(t *testing.T) {
	var config Config
	err := LoadConfig("non_existent_config.yaml", &config)
	if err == nil {
		t.Errorf("LoadConfig() error = nil, wantErr true for non-existent file")
	} else {
		// Check if the error message indicates file reading issue
		if !strings.Contains(err.Error(), "error reading YAML file") {
			t.Errorf("LoadConfig() error message = %q, want to contain %q", err.Error(), "error reading YAML file")
		}
	}
}

func TestLoadConfig_InvalidYAML(t *testing.T) {
	content := `
helmfile:
  - "path/to/helmfile1"
kustomize:
  - "path/to/kustomize1"
  - "path/to/kustomize2"
raw-manifests:
  - "path/to/raw-manifests1"
this: is: invalid: yaml
`
	tmpDir := t.TempDir()
	configFile := filepath.Join(tmpDir, "invalid_hekura.yaml")
	err := os.WriteFile(configFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create temporary invalid config file: %v", err)
	}

	var loadedConfig Config
	err = LoadConfig(configFile, &loadedConfig)
	if err == nil {
		t.Errorf("LoadConfig() error = nil, wantErr true for invalid YAML")
	} else {
		// Check if the error message indicates unmarshalling issue
		if !strings.Contains(err.Error(), "error unmarshalling YAML") {
			t.Errorf("LoadConfig() error message = %q, want to contain %q", err.Error(), "error unmarshalling YAML")
		}
	}
}
