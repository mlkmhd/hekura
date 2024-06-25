package pkg

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

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
		os.Chdir(element)

		cmd := exec.Command("helmfile", "template", "--debug")

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr

		err := cmd.Run()
		outStr, errStr := stdout.String(), stderr.String()
		if err != nil {
			fmt.Printf("Error executing helmfile command: %s\n", errStr)
			os.Exit(1)
		}

		WriteToFile(tempDir+"/resources.yaml", outStr)
		os.Chdir(rootDir)
	}

	for _, element := range config.Kustomize {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			os.Chdir(element)
			content, err := os.ReadFile(tempDir + "/resources.yaml")
			if err != nil {
				Logger.Fatalf("Error reading resources.yaml file: %v", err)
			}
			WriteToFile("resources.yaml", string(content))
			command := exec.Command("kustomize", "build", ".")
			output, err := command.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing kustomize command:", command, err)
				os.Exit(1)
			}

			WriteToFile(tempDir+"/resources.yaml", string(output))
			os.Remove("resources.yaml")
			os.Chdir(rootDir)
		}
	}

	resourcesContent, err := os.ReadFile(tempDir + "/resources.yaml")
	if err != nil {
		Logger.Fatalf("Error reading resource.yaml file: %v", err)
	}
	for _, element := range config.RawManifest {
		dirEntries, err := os.ReadDir(element)
		if err != nil {
			Logger.Fatalf("Error read raw manifest directory: %v", err)
		}
		for _, entry := range dirEntries {
			if !entry.IsDir() {
				manifestFileContent, err := os.ReadFile(element + "/" + entry.Name())
				if err != nil {
					Logger.Fatalf("Error read raw manifest file: %v", err)
				}
				resourcesContent = append(resourcesContent, []byte("\n---\n")...)
				resourcesContent = append(resourcesContent, manifestFileContent...)
			}
		}
	}

	result := string(resourcesContent)

	WriteToFile(tempDir+"/resources.yaml", result)
	return tempDir + "/resources.yaml"
}
