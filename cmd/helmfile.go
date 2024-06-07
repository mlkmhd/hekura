package cmd

import (
	"fmt"
	"os"

	"github.com/mlkmhd/hekura/pkg"
	"github.com/spf13/cobra"
)

var globalUsage = "Declaratively deploy your Kubernetes manifests, Kustomize configs, and Charts as Helm releases in one shot"

func Execute() {
	var rootCmd = &cobra.Command{
		Use:   "hekura",
		Short: globalUsage,
		Long:  globalUsage,
		Run: func(cmd *cobra.Command, args []string) {
			configFileName, _ := cmd.Flags().GetString("config-file")
			pkg.Build(pkg.LoadConfig(configFileName))
		},
	}

	var templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Print template",
		Long:  "Prints template",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Printing template")
		},
	}
	rootCmd.AddCommand(templateCmd)

	rootCmd.Flags().String("config-file", "hekura.yaml", "the config file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
