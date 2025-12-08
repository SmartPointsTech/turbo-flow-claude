package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/smartpointstech/spt-flow/cli/pkg/coder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start [workspace-name]",
	Short: "Start a workspace",
	Long:  `Start a stopped workspace. If no name is provided, the current directory name is used.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get configuration
		url := viper.GetString("url")
		token := viper.GetString("token")
		if url == "" || token == "" {
			return fmt.Errorf("not logged in. Run 'spt-flow init' first")
		}

		// Determine workspace name
		var workspaceName string
		if len(args) > 0 {
			workspaceName = args[0]
		} else {
			// Default to current directory name
			wd, err := os.Getwd()
			if err != nil {
				return fmt.Errorf("failed to get working directory: %w", err)
			}
			workspaceName = filepath.Base(wd)
		}

		// Initialize client
		client, err := coder.NewClient(url, token)
		if err != nil {
			return fmt.Errorf("failed to create client: %w", err)
		}

		fmt.Printf("Starting workspace '%s'...\n", workspaceName)

		// Get workspace
		ws, err := client.GetWorkspace(cmd.Context(), workspaceName)
		if err != nil {
			return fmt.Errorf("failed to get workspace: %w", err)
		}

		// Start workspace
		_, err = client.StartWorkspace(cmd.Context(), *ws)
		if err != nil {
			return fmt.Errorf("failed to start workspace: %w", err)
		}

		fmt.Println("Workspace started.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
