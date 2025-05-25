// Package cmd implements the command-line interface for Hekura.
// It uses the cobra library to define commands and flags.
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var globalUsage = "Declaratively deploy your Kubernetes manifests, Kustomize configs, and Charts as Helm releases in one shot"

// NewRootCmd creates and returns the root command for Hekura.
// The root command itself doesn't do much but provides the base for subcommands.
// It displays a general usage message.
func NewRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "hekura",
		Short: globalUsage,
		Long:  globalUsage,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(globalUsage)
		},
	}

	return cmd
}
