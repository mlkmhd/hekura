package main

import (
	"fmt"
	"os"

	"github.com/mlkmhd/hekura/cmd"
)

func main() {

	rootCmd := cmd.NewRootCmd()

	rootCmd.AddCommand(cmd.NewTemplateCmd())
	rootCmd.AddCommand(cmd.NewDiffCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
