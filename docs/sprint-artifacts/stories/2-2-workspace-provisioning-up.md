# Story 2.2: Workspace Provisioning (`up`)

Status: Done

## Story

As a Developer,
I want to provision and start my workspace with a single command,
So that I can get to work quickly without managing infrastructure details.

## Acceptance Criteria

1. **Given** A configured CLI (URL/Token)
2. **When** I run `spt-flow up [workspace-name]`
3. **Then** It checks if the workspace exists
4. **And** If it doesn't exist, it creates it using the standard template (Story 2.1)
5. **And** If it exists but is stopped, it starts the workspace
6. **And** It waits until the workspace is in `running` state
7. **And** It prints the VS Code Web URL and SSH command

## Tasks / Subtasks

- [x] Implement `up` Command Structure
  - [x] Create `cli/cmd/spt-flow/up.go`
  - [x] Define `upCmd`
  - [x] Accept optional workspace name argument (default to current directory name or user name)
- [x] Implement Coder Client Integration
  - [x] Create `cli/pkg/coder/client.go` (wrapper around Coder SDK)
  - [x] Implement `EnsureWorkspace` function
  - [x] Handle template versioning (use "active" version)
- [x] Implement Provisioning Logic
  - [x] Check existence: `client.GetWorkspace`
  - [x] Create: `client.CreateWorkspace`
  - [x] Start: `client.StartWorkspace`
  - [x] Wait: Poll status until `running`
- [x] Implement Output
  - [x] Print "Workspace is ready!"
  - [x] Print VS Code Web URL
  - [x] Print `ssh <workspace>` command

## Dev Notes

### Architecture Compliance

- **Package:** `pkg/coder` for Coder API interactions.
- **SDK:** Use `github.com/coder/coder/v2/coderd/codersdk`.
- **Concurrency:** Use `time.Ticker` for polling status.

### Technical Requirements

- **Template:** Hardcode the template name (e.g., "coder-template") for now, or allow configuration. Let's assume "standard-template" or similar.
- **Parameters:** Pass `repo_url` if creating a new workspace (flag `--repo`).

### References

- [Architecture: Workspace Lifecycle](docs/architecture.md#workspace-lifecycle)
- [Coder SDK Docs](https://pkg.go.dev/github.com/coder/coder/v2/coderd/codersdk)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 2.2)
- `docs/architecture.md`
- `docs/project_context.md`

### Completion Notes List

- Implemented `up` command with workspace name argument.
- Implemented Coder client wrapper using `codersdk`.
- Handled workspace creation (using `CreateUserWorkspace`) and starting.
- Implemented polling for workspace status.
- Verified build success.

### File List

- `cli/cmd/spt-flow/up.go`
- `cli/pkg/coder/client.go`
