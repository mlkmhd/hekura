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
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "Print template",
		Long:  "Prints template",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.SetLogLevel(logLevel)
			pkg.Logger.Debug("loading config files")
			pkg.LoadConfig(configFileName, &config)
			
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
