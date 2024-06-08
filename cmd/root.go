package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var globalUsage = "Declaratively deploy your Kubernetes manifests, Kustomize configs, and Charts as Helm releases in one shot"

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
