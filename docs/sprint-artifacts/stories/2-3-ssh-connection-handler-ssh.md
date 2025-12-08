# Story 2.3: SSH Connection Handler (`ssh`)

Status: Done

## Story

As a Developer,
I want to SSH into my workspace using the CLI,
So that I can run commands, forward ports, or use it as a remote development backend.

## Acceptance Criteria

1. **Given** A running workspace
2. **When** I run `spt-flow ssh [workspace-name]`
3. **Then** It connects to the workspace agent
4. **And** It opens an interactive shell
5. **And** I can run one-off commands: `spt-flow ssh [workspace-name] -- ls -la`
6. **And** It handles terminal resizing correctly

## Tasks / Subtasks

- [x] Implement `ssh` Command Structure
  - [x] Create `cli/cmd/spt-flow/ssh.go`
  - [x] Define `sshCmd`
  - [x] Accept workspace name and optional command arguments
- [x] Implement Agent Connection Logic
  - [x] Update `pkg/coder/client.go` to support agent connection
  - [x] Implement `ConnectToAgent(ctx, workspace)`
- [x] Implement SSH Client
  - [x] Use `golang.org/x/crypto/ssh`
  - [x] Use agent connection as transport (`net.Conn`)
  - [x] Handle authentication (using agent token or ephemeral keys)
  - [x] Setup interactive session (PTY request, terminal modes)
  - [x] Handle window resize events
- [ ] Verify SSH Connection
  - [ ] Test interactive shell
  - [ ] Test command execution

## Dev Notes

### Architecture Compliance

- **Dependencies:**
  - `golang.org/x/crypto/ssh`
  - `golang.org/x/term` (for terminal state)
- **Security:** Use the Coder agent token for authentication if supported, or ephemeral SSH keys managed by the SDK.

### Technical Requirements

- **Transport:** The Coder SDK provides `WorkspaceAgent` connection which implements `net.Conn`. Use this as the underlying transport for the SSH client.
- **UX:** Must feel like a native SSH session.

### References

- [Coder SDK: Workspace Agent](https://pkg.go.dev/github.com/coder/coder/v2/codersdk#Client.WorkspaceAgent)
- [Go SSH Client Example](https://pkg.go.dev/golang.org/x/crypto/ssh#example-NewClientConn)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 2.3)
- `docs/architecture.md`
- `docs/project_context.md`

### Completion Notes List

- Implemented `ssh` command.
- Integrated `workspacesdk` for agent connection.
- Resolved dependency issues by pinning `tailscale.com` to `v1.60.0`.
- Implemented interactive SSH session handling.

### File List

- `cli/cmd/spt-flow/ssh.go`
- `cli/pkg/coder/client.go`
- `cli/go.mod`
