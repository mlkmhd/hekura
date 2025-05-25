package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mlkmhd/hekura/pkg"
)

// NewTemplateCmd creates and returns the 'template' command.
// The 'template' command generates Kubernetes manifests based on the Hekura configuration
// and prints the resulting YAML to standard output.
// It takes a '--config' or '-c' flag to specify the configuration file (defaults to 'hekura.yaml').
// It also takes a '--loglevel' or '-l' flag to set the logging level.
func NewTemplateCmd() *cobra.Command {
	config := pkg.Config{}
	var configFileName string
	var logLevel string
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "Print template",
		Long:  "Prints template",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.SetLogLevel(logLevel)
			pkg.Logger.Debug("loading config files")
			if err := pkg.LoadConfig(configFileName, &config); err != nil {
				fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
				os.Exit(1)
			}

			builtResourceFile, err := os.ReadFile(pkg.Build(&config))
			if err != nil {
				pkg.Logger.Fatalf("Error reading config file: %v; %v", configFileName, err)
			}

			fmt.Println(string(builtResourceFile))
		},
	}

	cmd.Flags().StringVarP(&configFileName, "config", "c", "hekura.yaml", "the config file")
	cmd.Flags().StringVarP(&logLevel, "loglevel", "l", "info", "set log level (debug, info, warn, error, fatal, panic)")

	return cmd
}
