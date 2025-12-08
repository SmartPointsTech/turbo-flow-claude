# Story 4.2: Secret Injection (SSH & Git)

**Status:** Ready for Review
**Estimate:** 5 Points
**Priority:** High
**Assignee:** AI Agent

## Story

As a Developer,
I want to use my local SSH keys and Git credentials inside the workspace,
So that I can push code and access private servers without copying keys manually.

## Context

Security is paramount. We should not copy private keys to the workspace disk. Instead, we should leverage SSH Agent Forwarding. The Coder workspace needs to be configured to accept forwarded agents, and the CLI needs to ensure the user's local agent is forwarded when connecting.

## Acceptance Criteria

- [ ] **SSH Agent**: `ssh-add -l` inside the workspace shows the user's local keys.
- [ ] **Git Auth**: `git push` works out-of-the-box for repositories authenticated via SSH.
- [ ] **No Persistence**: Private keys are NOT stored on the container filesystem.
- [ ] **CLI Support**: `spt-flow ssh` enables forwarding automatically (`-A` behavior).

## Technical Notes

- **Terraform**: The `coder_agent` resource usually handles the listening part.
- **CLI**: The `ssh` command in `cmd/spt-flow/ssh.go` is already setting up an SSH client. We need to ensure it requests agent forwarding.
- **Windows**: On Windows, we might need a specific library (e.g., `github.com/Microsoft/go-winio`) to talk to the named pipe for the agent, but `golang.org/x/crypto/ssh/agent` is the standard abstraction.
