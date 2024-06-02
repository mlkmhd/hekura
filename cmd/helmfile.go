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
		helmfileDir := "helmfile"

		if _, err := os.Stat(helmfileDir); os.IsNotExist(err) {
			fmt.Println("the helmfile.yaml not found!")
		} else {
			command := exec.Command("helmfile", "template", "-f", "helmfile/helmfile.yaml")

			// Capture output
			output, err := command.CombinedOutput()
			if err != nil {
				fmt.Println("Error executing command:", err)
				os.Exit(1)
			}

			// Print output
			fmt.Println(string(output))
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
