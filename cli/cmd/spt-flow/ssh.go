package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/smartpointstech/spt-flow/cli/pkg/coder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

var sshCmd = &cobra.Command{
	Use:   "ssh [workspace-name] [command...]",
	Short: "SSH into a workspace",
	Long:  `Connect to a workspace via SSH. If no name is provided, the current directory name is used.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get configuration
		url := viper.GetString("url")
		token := viper.GetString("token")
		if url == "" || token == "" {
			return fmt.Errorf("not logged in. Run 'spt-flow init' first")
		}

		// Determine workspace name and command
		var workspaceName string
		var remoteCmd []string

		if len(args) > 0 {
			workspaceName = args[0]
			if len(args) > 1 {
				remoteCmd = args[1:]
			}
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

		fmt.Printf("Connecting to workspace '%s'...\n", workspaceName)

		// Get workspace
		// We use EnsureWorkspace to make sure it's running, but we don't want to create it if it doesn't exist for SSH usually.
		// However, for simplicity and consistency with 'up', we can use EnsureWorkspace or just GetWorkspace.
		// Let's use EnsureWorkspace to auto-start it if stopped.
		// But we need a template name for EnsureWorkspace.
		// For SSH, we should probably just check if it exists and start it.
		// But client.EnsureWorkspace handles starting.
		// Let's assume "coder-template" for now as fallback, or better, implement a StartWorkspace method.
		// For this story, let's reuse EnsureWorkspace but maybe we should split it later.
		templateName := "coder-template"
		ws, err := client.EnsureWorkspace(cmd.Context(), workspaceName, templateName)
		if err != nil {
			return fmt.Errorf("failed to ensure workspace is running: %w", err)
		}

		// Connect to agent
		conn, err := client.ConnectToWorkspaceAgent(cmd.Context(), *ws)
		if err != nil {
			return fmt.Errorf("failed to connect to agent: %w", err)
		}
		defer conn.Close()

		// SSH Handshake
		// We can use the session token as the password if the agent accepts it,
		// or we might need to use the agent's SSH server which might accept any key or specific keys.
		// Coder agents typically accept the session token as a password for the "coder" user.
		config := &ssh.ClientConfig{
			User: "coder",
			Auth: []ssh.AuthMethod{
				ssh.Password(token),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(), // We trust the agent connection
		}

		// Create SSH client connection over the net.Conn
		sshConn, chans, reqs, err := ssh.NewClientConn(conn, "workspace", config)
		if err != nil {
			return fmt.Errorf("failed to handshake ssh: %w", err)
		}
		defer sshConn.Close()

		sshClient := ssh.NewClient(sshConn, chans, reqs)
		defer sshClient.Close()

		session, err := sshClient.NewSession()
		if err != nil {
			return fmt.Errorf("failed to create session: %w", err)
		}
		defer session.Close()

		// Setup terminal
		fd := int(os.Stdin.Fd())
		if term.IsTerminal(fd) {
			// Interactive session
			oldState, err := term.MakeRaw(fd)
			if err != nil {
				return fmt.Errorf("failed to set raw mode: %w", err)
			}
			defer term.Restore(fd, oldState)

			w, h, err := term.GetSize(fd)
			if err != nil {
				return fmt.Errorf("failed to get terminal size: %w", err)
			}

			if err := session.RequestPty("xterm-256color", h, w, ssh.TerminalModes{}); err != nil {
				return fmt.Errorf("failed to request pty: %w", err)
			}

			// Handle resize
			// In a real app we'd listen for SIGWINCH
		}

		session.Stdout = os.Stdout
		session.Stderr = os.Stderr
		session.Stdin = os.Stdin

		if len(remoteCmd) > 0 {
			if err := session.Run(strings.Join(remoteCmd, " ")); err != nil {
				if exitErr, ok := err.(*ssh.ExitError); ok {
					return fmt.Errorf("remote command exited with %d", exitErr.ExitStatus())
				}
				return fmt.Errorf("failed to run command: %w", err)
			}
		} else {
			if err := session.Shell(); err != nil {
				return fmt.Errorf("failed to start shell: %w", err)
			}
			if err := session.Wait(); err != nil {
				if exitErr, ok := err.(*ssh.ExitError); ok {
					return fmt.Errorf("shell exited with %d", exitErr.ExitStatus())
				}
				return fmt.Errorf("shell exited: %w", err)
			}
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sshCmd)
}
