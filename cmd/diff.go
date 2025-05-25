package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/mlkmhd/hekura/pkg"
)

// NewDiffCmd creates and returns the 'diff' command.
// The 'diff' command generates Kubernetes manifests based on the Hekura configuration
// and then (placeholder) prints a message including the path to the generated resources.
//
// TODO: Implement actual diff functionality. This should ideally show differences
// against a live cluster or a previous state.
//
// It takes a '--config' flag to specify the configuration file (defaults to 'hekura.yaml').
func NewDiffCmd() *cobra.Command {
	config := pkg.Config{}
	var configFileName string
	var cmd = &cobra.Command{
		Use:   "diff",
		Short: "Show differences between generated manifests and current state (placeholder)",
		Long:  "Generates manifests and shows differences. Currently a placeholder.",
		Run: func(cmd *cobra.Command, args []string) {
			if err := pkg.LoadConfig(configFileName, &config); err != nil {
				fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
				os.Exit(1)
			}
			builtResourceFile := pkg.Build(&config)
			_ = pkg.Diff(builtResourceFile) // Assign to _ to indicate the return value is intentionally unused
		},
	}

	cmd.Flags().StringVar(&configFileName, "config", "hekura.yaml", "the config file")

	return cmd
}
