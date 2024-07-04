package cmd

import (
	"github.com/mlkmhd/hekura/pkg"
	"github.com/spf13/cobra"
)

func NewDiffCmd() *cobra.Command {
	config := pkg.Config{}
	var configFileName string
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "Print template",
		Long:  "Prints template",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.LoadConfig(configFileName, &config)
			builtResourceFile := pkg.Build(&config)
			pkg.Diff(builtResourceFile)
		},
	}

	cmd.Flags().StringVar(&configFileName, "config", "hekura.yaml", "the config file")

	return cmd
}
