package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		dirName, _ := cmd.Flags().GetString("dir")
		fmt.Println("checking the " + dirName + " directory")

		if _, err := os.Stat(dirName + "/helmfile.yaml"); os.IsNotExist(err) {
			fmt.Println("the helmfile.yaml not found!")
		} else {
			currentDir, _ := os.Getwd()
			os.Chdir(dirName)

			command := exec.Command("helmfile", "template", "-q")

			// Capture output
			output, err := command.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}
			writeToFile("resources.yaml", string(output))
			os.Chdir(currentDir)
		}

		if _, err := os.Stat(dirName + "/kustomization.yaml"); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			currentDir, _ := os.Getwd()
			os.Chdir(dirName)
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

		if _, err := os.Stat(dirName + "/raw-manifests"); os.IsNotExist(err) {
			fmt.Println("the kustomize patch files could not be found")
		} else {
			fmt.Println("There're some raw manifests")
		}
	},
}

func writeToFile(fileName string, content string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func Execute() {
	rootCmd.Flags().String("dir", "sample-manifests", "the directory for manifests")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
