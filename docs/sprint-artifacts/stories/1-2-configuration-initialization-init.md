# Story 1.2: Configuration Initialization (`init`)

Status: Done

## Story

As a User,
I want to initialize my local configuration interactively,
So that I don't have to provide the Coder URL and token for every command.

## Acceptance Criteria

1. **Given** I have a fresh installation
2. **When** I run `spt-flow init`
3. **Then** It prompts me for the Coder Deployment URL and Session Token
4. **And** It validates that the URL is reachable
5. **And** It saves the configuration to `~/.spt-flow.yaml` (or XDG config path)
6. **And** Subsequent commands read these values automatically

## Tasks / Subtasks

- [x] Implement `init` Command
  - [x] Create `cli/cmd/spt-flow/init.go`
  - [x] Define `initCmd` using Cobra
  - [x] Implement interactive prompts (using `manifoldco/promptui` or similar, or simple `fmt.Scan`)
- [x] Implement Configuration Management
  - [x] Configure `viper` to read/write config file
  - [x] Define config structure (URL, Token)
  - [x] Implement `SaveConfig` function
- [x] Implement URL Validation
  - [x] Add check to ensure URL is valid and reachable (simple HTTP HEAD/GET)
- [x] Verify Configuration Persistence
  - [x] Ensure config is saved to the correct location (`$HOME/.spt-flow.yaml` or XDG)
  - [x] Verify subsequent commands can read the config

## Dev Notes

### Architecture Compliance

- **Command Name:** `init` (per Architecture Naming Patterns)
- **Configuration:** Use `spf13/viper` for managing configuration.
- **File Path:** Default to `$HOME/.spt-flow.yaml`.
- **Validation:** Ensure URL is a valid HTTP/HTTPS URL.

### Technical Requirements

- **Interactive Prompts:** Use standard input/output for prompts.
- **Security:** Do not log the session token.
- **Error Handling:** Handle invalid URLs and file permission errors gracefully.

### References

- [Architecture: Naming Patterns](docs/architecture.md#naming-patterns)
- [Project Context: Critical Implementation Rules](docs/project_context.md#2-critical-implementation-rules)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 1.2)
- `docs/architecture.md`
- `docs/project_context.md`

### Agent Model Used

### Completion Notes List

- Implemented `init` command with interactive prompts.
- Configured `viper` to manage configuration in `$HOME/.spt-flow.yaml`.
- Implemented URL validation using `http.Head`.
- Verified configuration persistence and loading.
- Fixed security issue: Masked session token input.
- Fixed logic issue: Respected `--config` flag.

### File List

- `cli/cmd/spt-flow/init.go`
- `cli/cmd/spt-flow/root.go`
- `cli/go.mod`
- `cli/go.sum`
