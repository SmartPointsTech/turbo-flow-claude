package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	versionStr = "dev"
	commit     = "none"
	date       = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of spt-flow",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("spt-flow version %s, commit %s, built at %s, %s/%s\n", versionStr, commit, date, runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
