package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/smartpointstech/spt-flow/cli/pkg/coder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var upCmd = &cobra.Command{
	Use:   "up [workspace-name]",
	Short: "Create or start a workspace",
	Long:  `Create a new workspace or start an existing one. If no name is provided, the current directory name is used.`,
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

		// Parse flags
		dotfilesURL, _ := cmd.Flags().GetString("dotfiles-url")

		fmt.Printf("Provisioning workspace '%s'...\n", workspaceName)

		// Ensure workspace exists and is running
		// TODO: Make template name configurable
		templateName := "coder-template"

		params := map[string]string{}
		if dotfilesURL != "" {
			params["dotfiles_url"] = dotfilesURL
		}

		ws, err := client.EnsureWorkspace(cmd.Context(), coder.EnsureWorkspaceOptions{
			Name:         workspaceName,
			TemplateName: templateName,
			Parameters:   params,
		})
		if err != nil {
			return fmt.Errorf("failed to provision workspace: %w", err)
		}

		fmt.Println("Workspace is ready!")
		fmt.Printf("VS Code Web: %s\n", fmt.Sprintf("%s/@%s/%s/apps/code-server", url, ws.OwnerName, ws.Name))
		fmt.Printf("SSH: ssh coder.%s\n", ws.Name)

		return nil
	},
}

func init() {
	upCmd.Flags().String("dotfiles-url", "", "URL of a dotfiles repository to install")
	rootCmd.AddCommand(upCmd)
}
