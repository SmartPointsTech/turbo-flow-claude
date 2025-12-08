package coder

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/coder/coder/v2/codersdk"
	"github.com/coder/coder/v2/codersdk/workspacesdk"
	"github.com/google/uuid"
)

type Client struct {
	sdk *codersdk.Client
}

func NewClient(rawURL, token string) (*Client, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %w", err)
	}

	client := codersdk.New(u)
	client.SetSessionToken(token)

	return &Client{sdk: client}, nil
}

func (c *Client) GetWorkspace(ctx context.Context, name string) (*codersdk.Workspace, error) {
	ws, err := c.sdk.WorkspaceByOwnerAndName(ctx, codersdk.Me, name, codersdk.WorkspaceOptions{})
	if err != nil {
		return nil, err
	}
	return &ws, nil
}

type EnsureWorkspaceOptions struct {
	Name         string
	TemplateName string
	Parameters   map[string]string
}

func (c *Client) EnsureWorkspace(ctx context.Context, opts EnsureWorkspaceOptions) (*codersdk.Workspace, error) {
	// Check if workspace exists
	ws, err := c.GetWorkspace(ctx, opts.Name)
	if err == nil {
		// Workspace exists, ensure it's running
		return c.ensureRunning(ctx, *ws)
	}

	// Create workspace
	// First, find the template
	template, err := c.findTemplateByName(ctx, opts.TemplateName)
	if err != nil {
		return nil, fmt.Errorf("failed to find template %q: %w", opts.TemplateName, err)
	}

	// Build rich parameters
	var richParams []codersdk.WorkspaceBuildParameter
	for k, v := range opts.Parameters {
		richParams = append(richParams, codersdk.WorkspaceBuildParameter{
			Name:  k,
			Value: v,
		})
	}

	// Create the workspace
	wsVal, err := c.sdk.CreateUserWorkspace(ctx, codersdk.Me, codersdk.CreateWorkspaceRequest{
		TemplateID:          template.ID,
		Name:                opts.Name,
		AutostartSchedule:   ptr("CRON_TZ=UTC 30 9 * * 1-5"),   // Default 9:30 AM Mon-Fri
		TTLMillis:           ptr(8 * time.Hour.Milliseconds()), // Default 8 hours
		RichParameterValues: richParams,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create workspace: %w", err)
	}
	ws = &wsVal

	// Wait for build to complete
	if err := c.waitForBuild(ctx, ws.LatestBuild.ID); err != nil {
		return nil, fmt.Errorf("failed to wait for build: %w", err)
	}

	// Refresh workspace
	wsVal, err = c.sdk.Workspace(ctx, ws.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh workspace: %w", err)
	}
	ws = &wsVal

	return ws, nil
}

func (c *Client) findTemplateByName(ctx context.Context, name string) (codersdk.Template, error) {
	// List templates to find the one with the matching name
	templates, err := c.sdk.Templates(ctx, codersdk.TemplateFilter{})
	if err != nil {
		return codersdk.Template{}, err
	}

	for _, t := range templates {
		if t.Name == name {
			return t, nil
		}
	}

	return codersdk.Template{}, fmt.Errorf("template not found")
}

func (c *Client) StartWorkspace(ctx context.Context, ws codersdk.Workspace) (*codersdk.Workspace, error) {
	if ws.LatestBuild.Job.Status == codersdk.ProvisionerJobRunning {
		// Already running or starting, wait for it
		if err := c.waitForBuild(ctx, ws.LatestBuild.ID); err != nil {
			return nil, err
		}
		// Refresh
		updatedWs, err := c.sdk.Workspace(ctx, ws.ID)
		if err != nil {
			return nil, err
		}
		return &updatedWs, nil
	}

	if ws.LatestBuild.Transition == codersdk.WorkspaceTransitionStart && ws.LatestBuild.Job.Status == codersdk.ProvisionerJobSucceeded {
		// Already started and running
		return &ws, nil
	}

	// Start the workspace
	build, err := c.sdk.CreateWorkspaceBuild(ctx, ws.ID, codersdk.CreateWorkspaceBuildRequest{
		Transition: codersdk.WorkspaceTransitionStart,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to start workspace: %w", err)
	}

	if err := c.waitForBuild(ctx, build.ID); err != nil {
		return nil, fmt.Errorf("failed to wait for start: %w", err)
	}

	updatedWs, err := c.sdk.Workspace(ctx, ws.ID)
	if err != nil {
		return nil, err
	}
	return &updatedWs, nil
}

func (c *Client) StopWorkspace(ctx context.Context, ws codersdk.Workspace) (*codersdk.Workspace, error) {
	if ws.LatestBuild.Job.Status == codersdk.ProvisionerJobRunning {
		// Wait for current build to finish before stopping
		if err := c.waitForBuild(ctx, ws.LatestBuild.ID); err != nil {
			return nil, err
		}
		// Refresh
		updatedWs, err := c.sdk.Workspace(ctx, ws.ID)
		if err != nil {
			return nil, err
		}
		ws = updatedWs
	}

	if ws.LatestBuild.Transition == codersdk.WorkspaceTransitionStop && ws.LatestBuild.Job.Status == codersdk.ProvisionerJobSucceeded {
		// Already stopped
		return &ws, nil
	}

	// Stop the workspace
	build, err := c.sdk.CreateWorkspaceBuild(ctx, ws.ID, codersdk.CreateWorkspaceBuildRequest{
		Transition: codersdk.WorkspaceTransitionStop,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to stop workspace: %w", err)
	}

	if err := c.waitForBuild(ctx, build.ID); err != nil {
		return nil, fmt.Errorf("failed to wait for stop: %w", err)
	}

	updatedWs, err := c.sdk.Workspace(ctx, ws.ID)
	if err != nil {
		return nil, err
	}
	return &updatedWs, nil
}

func (c *Client) ensureRunning(ctx context.Context, ws codersdk.Workspace) (*codersdk.Workspace, error) {
	return c.StartWorkspace(ctx, ws)
}

func (c *Client) waitForBuild(ctx context.Context, buildID uuid.UUID) error {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-ticker.C:
			build, err := c.sdk.WorkspaceBuild(ctx, buildID)
			if err != nil {
				return err
			}

			if build.Job.Status == codersdk.ProvisionerJobSucceeded {
				return nil
			}
			if build.Job.Status == codersdk.ProvisionerJobFailed {
				return fmt.Errorf("build failed: %s", build.Job.Error)
			}
			if build.Job.Status == codersdk.ProvisionerJobCanceled {
				return fmt.Errorf("build canceled")
			}
		}
	}
}

func (c *Client) ConnectToWorkspaceAgent(ctx context.Context, ws codersdk.Workspace) (net.Conn, error) {
	// Find the first running resource with an agent
	for _, res := range ws.LatestBuild.Resources {
		for _, agent := range res.Agents {
			if agent.Status == codersdk.WorkspaceAgentConnected && agent.LifecycleState == codersdk.WorkspaceAgentLifecycleReady {
				// Connect to the agent
				client := workspacesdk.New(c.sdk)
				agentConn, err := client.DialAgent(ctx, agent.ID, nil)
				if err != nil {
					return nil, fmt.Errorf("failed to dial agent: %w", err)
				}

				// Get SSH connection
				conn, err := agentConn.SSH(ctx)
				if err != nil {
					agentConn.Close()
					return nil, fmt.Errorf("failed to establish SSH connection: %w", err)
				}

				return conn, nil
			}
		}
	}
	return nil, fmt.Errorf("no connected agent found")
}

func ptr[T any](v T) *T {
	return &v
}
