package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func Build(config Config) {

	rootDir, _ := os.Getwd()

	for _, element := range config.Helmfile {
		os.Chdir(element)

		command := exec.Command("helmfile", "template", "-q")

		// Capture output
		output, err := command.CombinedOutput()
		if err != nil {
			fmt.Println("Error executing command:", err)
			os.Exit(1)
		}
		WriteToFile("/tmp/resources.yaml", string(output))
		os.Chdir(rootDir)
	}

	for _, element := range config.Kustomize {
		if _, err := os.Stat(element); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			os.Chdir(element)
			content, err := ioutil.ReadFile("/tmp/resources.yaml")
			err = ioutil.WriteFile("resources.yaml", content, 0644)
			command := exec.Command("kustomize", "build", ".")
			output, err := command.CombinedOutput()
			os.Remove("resources.yaml")

			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			WriteToFile("/tmp/resources-patched.yaml", string(output))
			os.Chdir(rootDir)
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
