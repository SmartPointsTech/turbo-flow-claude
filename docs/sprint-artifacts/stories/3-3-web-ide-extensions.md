# Story 3.3: Web IDE & Extensions

**Status:** Ready for Review
**Estimate:** 3 Points
**Priority:** Low
**Assignee:** AI Agent

## Story

As a Developer,
I want a fully configured VS Code environment in my browser,
So that I can code from any device without installing local tools.

## Context

We need to ensure the "Web IDE" (code-server) experience is excellent out of the box. This means pre-installing critical extensions for our stack (Go, Rust, Docker) and setting the default shell to `zsh` so users don't have to configure it manually every time.

## Acceptance Criteria

- [ ] **Extensions**: The following VS Code extensions are pre-installed in the image:
  - `golang.go`
  - `rust-lang.rust-analyzer`
  - `ms-azuretools.vscode-docker`
  - `ms-python.python`
- [ ] **Shell**: The default terminal in VS Code uses `/bin/zsh`.
- [ ] **Settings**: Basic settings (e.g., dark mode) are applied if possible.

## Technical Notes

- **Implementation**: This should be done in the `Dockerfile` (`infra/images/coder-base/Dockerfile`) to bake it into the image.
- **Code Server**: The `coder-template` installs `code-server` in the startup script. We might need to move this installation to the Dockerfile to pre-install extensions, OR keep it in the script and run `code-server --install-extension` at startup (slower boot).
- **Decision**: Move `code-server` installation to the Dockerfile for faster startup and immutable extension versions.
