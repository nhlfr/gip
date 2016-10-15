package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/nhlfr/gip/pkg/env"
)

func NewEnvCommand() *cobra.Command {
	activateCmd := &cobra.Command{
		Use: "activate",
		Run: runActivate,
	}

	deleteCmd := &cobra.Command{
		Use: "delete",
		Run: runDelete,
	}

	listCmd := &cobra.Command{
		Use: "list",
		Run: runList,
	}

	envCmd := &cobra.Command{
		Use: "env [arg...]",
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Usage(); err != nil {
				fmt.Fprintln(os.Stderr, err.Error())
			}
		},
	}

	envCmd.AddCommand(activateCmd)
	envCmd.AddCommand(deleteCmd)
	envCmd.AddCommand(listCmd)

	return envCmd
}

func runActivate(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(1)
	}

	if err := env.ActivateOrCreateEnv(args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runDelete(cmd *cobra.Command, args []string) {
	if len(args) != 1 {
		cmd.Usage()
		os.Exit(1)
	}

	if err := env.DeleteEnv(args[0]); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func runList(cmd *cobra.Command, args []string) {
	envs, err := env.ListEnvs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	for _, e := range envs {
		fmt.Println(e)
	}
}
