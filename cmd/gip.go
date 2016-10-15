package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nhlfr/gip/pkg/cli"
)

func newGipCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "gip",
		Short: "gip is a virtual environment and package manager for Go",
		Long:  "gip is a virtual environment and package manager for Go, very similar to Python virtualenvwrapper and pip",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		},
	}
}

func main() {
	cmd := newGipCommand()
	cmd.AddCommand(cli.NewEnvCommand())
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
