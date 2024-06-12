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
	var cmd = &cobra.Command{
		Use:   "template",
		Short: "Print template",
		Long:  "Prints template",
		Run: func(cmd *cobra.Command, args []string) {
			pkg.LoadConfig(configFileName, &config)
			builtResourceFile, _ := os.ReadFile(pkg.Build(&config))
			fmt.Println(string(builtResourceFile))
		},
	}

	cmd.Flags().StringVar(&configFileName, "config", "hekura.yaml", "the config file")

	return cmd
}
