# Story 1.3: Automated Version Check

Status: Done

## Story

As a Platform Engineer,
I want the CLI to check for updates automatically,
So that users are aware when their tool is out of sync with the platform.

## Acceptance Criteria

1. **Given** An outdated local CLI version
2. **When** I run any command (e.g., `spt-flow up`)
3. **Then** The CLI queries the release source (GitHub Releases) in the background
4. **And** It prints a warning to stderr: "Update available: vX.Y.Z -> vA.B.C"
5. **And** It respects a `SPT_FLOW_SKIP_UPDATE_CHECK=true` env var to disable this behavior

## Tasks / Subtasks

- [x] Implement Version Check Logic
  - [x] Create `cli/internal/version/check.go`
  - [x] Implement `CheckForUpdate` function
  - [x] Use `google/go-github` or simple HTTP client to check releases
  - [x] Compare current version with latest version (using `hashicorp/go-version` or semver)
- [x] Integrate with Root Command
  - [x] Add `PersistentPreRun` hook in `root.go`
  - [x] Run check in background (goroutine)
  - [x] Print warning to stderr if update available
- [x] Implement Skip Flag/Env Var
  - [x] Check `SPT_FLOW_SKIP_UPDATE_CHECK` env var
  - [x] Skip check if set

## Dev Notes

### Architecture Compliance

- **Package:** `internal/version`
- **Concurrency:** Run check in a goroutine to avoid blocking the main command.
- **Output:** Print to `os.Stderr` to avoid polluting stdout (which might be used for piping).
- **Timeout:** Set a short timeout (e.g., 2s) for the check to avoid hanging.

### Technical Requirements

- **GitHub API:** Use public API for releases.
- **Semver:** Ensure robust version comparison.
- **State:** Optionally cache the last check time to avoid checking on every single command run (e.g., check once per hour). *Note: AC implies "any command", but caching is a good practice.*

### References

- [Architecture: Observability](docs/architecture.md#observability)
- [Project Context: Critical Implementation Rules](docs/project_context.md#2-critical-implementation-rules)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 1.3)
- `docs/architecture.md`
- `docs/project_context.md`

### Completion Notes List

- Implemented `CheckForUpdate` logic using GitHub API.
- Integrated version check into `root.go` `PersistentPreRun`.
- Implemented `SPT_FLOW_SKIP_UPDATE_CHECK` env var support.
- Fixed module import path issues.

### File List

- `cli/internal/version/check.go`
- `cli/cmd/spt-flow/root.go`
- `cli/cmd/spt-flow/version.go`
