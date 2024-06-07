package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Execute(config Config) {

	for _, element := range config.Helmfile {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the helmfile.yaml not found!")
		} else {
			currentDir, _ := os.Getwd()
			os.Chdir(filepath.Dir(element))

			command := exec.Command("helmfile", "template", "-q")

			// Capture output
			output, err := command.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}
			WriteToFile("resources.yaml", string(output))
			os.Chdir(currentDir)
		}
	}

	for _, element := range config.Kustomize {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			currentDir, _ := os.Getwd()
			os.Chdir(filepath.Dir(element))
			command := exec.Command("kustomize", "build", ".")
			output, err := command.CombinedOutput()

			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			// Print output
			fmt.Println(string(output))
			os.Chdir(currentDir)
		}
	}

	for _, element := range config.RawManifest {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			fmt.Println("There're some raw manifests")
		}
	}
}
