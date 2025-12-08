package main

import (
	"fmt"
	"os"

	"github.com/smartpointstech/spt-flow/cli/internal/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:          "spt-flow",
	Short:        "spt-flow is a CLI for managing Coder workspaces",
	Long:         `spt-flow is a CLI for managing Coder workspaces, providing a unified interface for provisioning, connecting, and managing development environments.`,
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		go func() {
			if msg := version.CheckForUpdate(versionStr); msg != "" {
				fmt.Fprintln(os.Stderr, msg)
			}
		}()
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.spt-flow.yaml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".spt-flow")
	}

	viper.SetEnvPrefix("SPT_FLOW")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// Config loaded
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
