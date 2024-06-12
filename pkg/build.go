package pkg

import (
	"fmt"
	"os"
	"os/exec"
)

func Build(config *Config) string {

	rootDir, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp("", "template")

	for _, element := range config.Helmfile {
		os.Chdir(element)

		command := exec.Command("helmfile", "template", "-q")

		// Capture output
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1)
		}
		WriteToFile(tempDir+"/resources.yaml", string(output))
		os.Chdir(rootDir)
	}

	for _, element := range config.Kustomize {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			os.Chdir(element)
			content, _ := os.ReadFile(tempDir + "/resources.yaml")
			WriteToFile("resources.yaml", string(content))
			command := exec.Command("kustomize", "build", ".")
			output, err := command.CombinedOutput()

			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			WriteToFile(tempDir+"/resources.yaml", string(output))
			os.Remove("resources.yaml")
			os.Chdir(rootDir)
		}
	}

	resourcesContent, _ := os.ReadFile(tempDir + "/resources.yaml")
	for _, element := range config.RawManifest {
		dirEntries, _ := os.ReadDir(element)
		for _, entry := range dirEntries {
			if !entry.IsDir() {
				manifestFileContent, _ := os.ReadFile(element + "/" + entry.Name())
				resourcesContent = append(resourcesContent, []byte("\n---\n")...)
				resourcesContent = append(resourcesContent, manifestFileContent...)
			}
		}
	}

	result := string(resourcesContent)

	WriteToFile(tempDir+"/resources.yaml", result)
	return tempDir + "/resources.yaml"
}
