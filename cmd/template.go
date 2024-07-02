package cmd

import (
	"fmt"
	"os"

	"github.com/mlkmhd/hekura/pkg"
	"github.com/spf13/cobra"
)

func NewTemplateCmd() *cobra.Command {
	config := pkg.Config{}
	var configFileName string
	var logLevel string
	var output string
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "Prints generated maniefsts",
		Long:  "Prints generated manifests",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.SetLogLevel(logLevel)
			pkg.Logger.Debug("loading config files")
			pkg.LoadConfig(configFileName, &config)

			builtResourceFile, err := os.ReadFile(pkg.Build(&config))
			if err != nil {
				pkg.Logger.Fatalf("Error reading config file: %v; %v", configFileName, err)
			}

			pkg.WriteToFile(output, string(builtResourceFile))
			fmt.Printf("Successfully generated manifests to %v file\n", output)
		},
	}

	cmd.Flags().StringVarP(&configFileName, "config", "c", "hekura.yaml", "the config file")
	cmd.Flags().StringVarP(&logLevel, "loglevel", "l", "info", "set log level (debug, info, warn, error, fatal, panic)")
	cmd.Flags().StringVarP(&output, "output", "o", "all.yaml", "set the output file for generated manifests")

	return cmd
}
