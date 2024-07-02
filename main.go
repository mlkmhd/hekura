package main

import (
	"fmt"
	"os"

	"github.com/mlkmhd/hekura/cmd"
	"github.com/mlkmhd/hekura/pkg"
)

func main() {
	pkg.Logger.Info("starting app")

	rootCmd := cmd.NewRootCmd()

	rootCmd.AddCommand(cmd.NewTemplateCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
