package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:          "spt-flow",
	Short:        "spt-flow is a CLI for managing Coder workspaces",
	Long:         `spt-flow is a CLI for managing Coder workspaces, providing a unified interface for provisioning, connecting, and managing development environments.`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
