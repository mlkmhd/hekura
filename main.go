package main

import (
	"fmt"
	"os"

	"github.com/mlkmhd/hekura/cmd"
)

func main() {

	rootCmd := cmd.NewRootCmd()

	t := cmd.NewTemplateCmd()
	
	rootCmd.AddCommand(t)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
