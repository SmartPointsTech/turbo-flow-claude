# Story 4.1: Dotfiles & Personalization Hook

**Status:** Ready for Review
**Estimate:** 5 Points
**Priority:** Medium
**Assignee:** AI Agent

## Story

As a Developer,
I want my workspace to load my personal shell configuration (dotfiles),
So that I have my aliases, prompt, and tools ready immediately.

## Context

Personalization is key for developer productivity. We need a way to clone a user's dotfiles repository into the workspace startup and run their installation script. This should be driven by a parameter passed to the `spt-flow start` or `up` command, which is then passed to the Coder template.

## Acceptance Criteria

- [ ] **Parameter**: The template accepts a `dotfiles_url` variable.
- [ ] **Cloning**: The workspace clones the repo to `~/dotfiles` (or `~/.dotfiles`) if the parameter is provided.
- [ ] **Execution**: It detects and runs one of `install.sh`, `setup.sh`, or `bootstrap.sh`.
- [ ] **Resilience**: Failures in the dotfiles script do NOT prevent the workspace from starting (warnings logged).
- [ ] **CLI Support**: `spt-flow up --dotfiles <url>` passes this parameter.

## Technical Notes

- **Terraform**: Add `variable "dotfiles_url"` to `infra/terraform/coder-template/variables.tf`.
- **Startup Script**: Update `infra/terraform/coder-template/main.tf` startup script to handle the cloning and execution.
- **CLI**: Update `cmd/spt-flow/up.go` to accept the flag and pass it to the parameter list when creating the workspace.
