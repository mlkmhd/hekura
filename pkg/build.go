package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

// Build processes the Helmfile, Kustomize, and raw manifest configurations
// defined in the Config struct. It generates Kubernetes manifests by:
// 1. Running `helmfile template` for each Helmfile configuration.
// 2. Running `kustomize build` for each Kustomize configuration, using the output from the previous step.
// 3. Appending all raw manifests.
// The combined manifests are written to a temporary file, and the path to this file is returned.
// The function uses fatal logging for errors encountered during the build process, exiting the application.
func Build(config *Config) string {

	rootDir, err := os.Getwd()
	if err != nil {
		Logger.Fatalf("Error getting current working directory: %v", err)
	}
	tempDir, err := os.MkdirTemp("", "template")
	if err != nil {
		Logger.Fatalf("Error creating temp directory: %v", err)
	}

	for _, element := range config.Helmfile {
		if err := os.Chdir(element); err != nil {
			Logger.Fatalf("Error changing directory to %s: %v", element, err)
		}

		command := exec.Command("helmfile", "template", "-q")

		// Capture output
		output, cmdErr := command.CombinedOutput()
		if cmdErr != nil {
			fmt.Println("Error executing command:", cmdErr)
			os.Exit(1)
		}
		if err := WriteToFile(tempDir+"/resources.yaml", string(output)); err != nil {
			Logger.Fatalf("Error writing helmfile output to %s: %v", tempDir+"/resources.yaml", err)
		}
		if err := os.Chdir(rootDir); err != nil {
			Logger.Fatalf("Error changing directory to %s: %v", rootDir, err)
		}
	}

	for _, element := range config.Kustomize {
		if _, statErr := os.Stat(element); os.IsNotExist(statErr) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			if err := os.Chdir(element); err != nil {
				Logger.Fatalf("Error changing directory to %s: %v", element, err)
			}
			content, readErr := os.ReadFile(tempDir + "/resources.yaml")
			if readErr != nil {
				Logger.Fatalf("Error reading resources.yaml file: %v", readErr)
			}
			if err := WriteToFile("resources.yaml", string(content)); err != nil {
				Logger.Fatalf("Error writing intermediate kustomize input to resources.yaml: %v", err)
			}
			command := exec.Command("kustomize", "build", ".")
			output, cmdErr := command.CombinedOutput()
			if cmdErr != nil {
				fmt.Println("Error executing kustomize command:", cmdErr)
				os.Exit(1)
			}

			if err := WriteToFile(tempDir+"/resources.yaml", string(output)); err != nil {
				Logger.Fatalf("Error writing kustomize output to %s: %v", tempDir+"/resources.yaml", err)
			}
			if err := os.Remove("resources.yaml"); err != nil {
				Logger.Warnf("Error removing temporary resources.yaml in kustomize directory %s: %v", element, err)
			}
			if err := os.Chdir(rootDir); err != nil {
				Logger.Fatalf("Error changing directory to %s: %v", rootDir, err)
			}
		}
	}

	resourcesContent, finalReadErr := os.ReadFile(tempDir + "/resources.yaml")
	if finalReadErr != nil {
		Logger.Fatalf("Error reading resource.yaml file: %v", finalReadErr)
	}
	for _, element := range config.RawManifest {
		dirEntries, readDirErr := os.ReadDir(element)
		if readDirErr != nil {
			Logger.Fatalf("Error read raw manifest directory: %v", readDirErr)
		}
		for _, entry := range dirEntries {
			if !entry.IsDir() {
				manifestFileContent, manifestReadErr := os.ReadFile(element + "/" + entry.Name())
				if manifestReadErr != nil {
					Logger.Fatalf("Error read raw manifest file: %v", manifestReadErr)
				}
				resourcesContent = append(resourcesContent, []byte("\n---\n")...)
				resourcesContent = append(resourcesContent, manifestFileContent...)
			}
		}
	}

	result := string(resourcesContent)

	if err := WriteToFile(tempDir+"/resources.yaml", result); err != nil {
		Logger.Fatalf("Error writing final result to %s: %v", tempDir+"/resources.yaml", err)
	}
	return tempDir + "/resources.yaml"
}
