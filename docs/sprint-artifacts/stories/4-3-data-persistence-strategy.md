# Story 4.3: Data Persistence Strategy

**Status:** Ready for Review
**Estimate:** 3 Points
**Priority:** Critical
**Assignee:** AI Agent

## Story

As a Developer,
I want my project files and installed tools to persist when the workspace is stopped and started,
So that I don't lose my work.

## Context

Currently, the Docker container is ephemeral. If it's removed and recreated (e.g., during an update or rebuild), data in `/home/coder` is lost unless we mount a volume. Coder handles persistent volumes via `coder_agent` and Docker volumes, but we need to ensure our Terraform template explicitly defines and mounts a volume for the home directory.

## Acceptance Criteria

- [ ] **Persistence**: Create a file in `~/`, stop the workspace, start it again -> file exists.
- [ ] **Volume**: A Docker volume is created and attached to `/home/coder`.
- [ ] **Independence**: The volume survives container destruction.

## Technical Notes

- **Terraform**:
  - Resource: `docker_volume` "home_volume".
  - `docker_container`: Mount `home_volume` to `/home/coder`.
  - Lifecycle: Ensure volume is not destroyed on `terraform apply` unless explicitly requested.
