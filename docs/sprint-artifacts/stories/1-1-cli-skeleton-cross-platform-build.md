# Story 1.1: CLI Skeleton & Cross-Platform Build

Status: Done

## Story

As a Developer,
I want a CLI binary that runs on my specific Operating System,
so that I can manage my workspaces regardless of my machine type.

## Acceptance Criteria

1. **Given** The source code is compiled via `goreleaser`
2. **When** I run the binary on Linux (amd64), macOS (arm64), and Windows (amd64)
3. **Then** The command `spt-flow --version` prints the correct version and commit hash
4. **And** The command `spt-flow --help` prints the root help text
5. **And** The binary has no external runtime dependencies (static linking)

## Tasks / Subtasks

- [x] Initialize Go Module and Directory Structure
  - [x] Create `cli/` directory
  - [x] Run `go mod init github.com/smartpointstech/spt-flow/cli`
  - [x] Create `cli/cmd/spt-flow/main.go`
  - [x] Create `cli/internal/` directory
- [x] Implement Root Command (Cobra)
  - [x] Install `spf13/cobra`
  - [x] Define root command in `cli/cmd/spt-flow/root.go` (or similar)
  - [x] Implement `Execute()` function
  - [x] Set `SilenceUsage: true`
- [x] Implement Version Command
  - [x] Create `version` command
  - [x] Wire up build-time variables (`version`, `commit`, `date`) using `ldflags`
- [x] Configure Goreleaser
  - [x] Create `.goreleaser.yaml` in `cli/`
  - [x] Configure builds for `linux/amd64`, `darwin/arm64`, `windows/amd64`
  - [x] Configure `CGO_ENABLED=0` for static linking
- [x] Verify Cross-Platform Builds
  - [x] Run `goreleaser build --snapshot --clean` locally
  - [x] Verify binaries exist in `dist/`

## Dev Notes

### Architecture Compliance

- **Binary Name:** `spt-flow` (per Architecture Naming Patterns)
- **Directory Structure:**
  - `cli/cmd/spt-flow/` for entry point.
  - `cli/internal/` for logic.
- **Libraries:**
  - `github.com/spf13/cobra` (Commands)
  - `github.com/spf13/viper` (Config - though not strictly needed for *this* story, good to setup)
- **Error Handling:** Use `fmt.Errorf` wrapping. Handle errors in `main.go` by printing to stderr and exiting non-zero.

### Technical Requirements

- **Go Version:** 1.21+
- **Static Linking:** Ensure `CGO_ENABLED=0`.
- **Windows Support:** Ensure paths use `filepath.Join`.

### References

- [Architecture: Naming Patterns](docs/architecture.md#naming-patterns)
- [Architecture: Project Structure](docs/architecture.md#complete-project-directory-structure)
- [Project Context: Critical Implementation Rules](docs/project_context.md#2-critical-implementation-rules)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 1.1)
- `docs/architecture.md`
- `docs/project_context.md`

### Agent Model Used

### Completion Notes List

- Initialized Go module and directory structure.
- Implemented root command with Cobra.
- Implemented version command with build-time variable injection.
- Configured Goreleaser for cross-platform builds (Linux, Windows, macOS).
- Verified builds locally (fixed: used `--clean` instead of deprecated `--rm-dist`).

### File List

- `cli/go.mod`
- `cli/go.sum`
- `cli/cmd/spt-flow/main.go`
- `cli/cmd/spt-flow/root.go`
- `cli/cmd/spt-flow/version.go`
- `cli/.goreleaser.yaml`
