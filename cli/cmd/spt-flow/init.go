package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/term"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize local configuration",
	Long:  `Initialize local configuration by prompting for Coder URL and Session Token.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		// Prompt for URL
		fmt.Print("Enter Coder Deployment URL: ")
		url, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("failed to read URL: %w", err)
		}
		url = strings.TrimSpace(url)

		// Validate URL
		if err := validateURL(url); err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}

		// Prompt for Token (Masked)
		fmt.Print("Enter Session Token: ")
		byteToken, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return fmt.Errorf("failed to read token: %w", err)
		}
		fmt.Println() // Print newline after masked input
		token := strings.TrimSpace(string(byteToken))

		// Save Config
		viper.Set("url", url)
		viper.Set("token", token)

		var configPath string
		if cfgFile != "" {
			configPath = cfgFile
		} else {
			configDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get user home dir: %w", err)
			}
			configPath = filepath.Join(configDir, ".spt-flow.yaml")
		}

		viper.SetConfigFile(configPath)
		if err := viper.WriteConfig(); err != nil {
			// Try creating if it doesn't exist
			if err := viper.SafeWriteConfig(); err != nil {
				return fmt.Errorf("failed to write config: %w", err)
			}
		}

		fmt.Printf("Configuration saved to %s\n", configPath)
		return nil
	},
}

func validateURL(url string) error {
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode >= 400 {
		return fmt.Errorf("URL returned status %d", resp.StatusCode)
	}
	return nil
}

func init() {
	rootCmd.AddCommand(initCmd)
}
