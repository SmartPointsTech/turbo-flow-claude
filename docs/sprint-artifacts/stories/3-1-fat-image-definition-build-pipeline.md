# Story 3.1: "Fat Image" Definition & Build Pipeline

**Status:** Ready for Review
**Estimate:** 5 Points
**Priority:** High
**Assignee:** AI Agent

## Story

As a Platform Engineer,
I want an automated pipeline that builds our standard developer image,
So that everyone uses the same tooling and updates are safe.

## Context

We need a standardized "Fat Image" that serves as the base for all developer workspaces. This image must contain the core toolchain (Go, Rust, Python, Node.js, Docker) to ensure consistency across the team.

This story covers:

- Defining the `Dockerfile` for the "coder-base" image.
- Creating a GitHub Actions workflow to build, test, and push this image.
- Integrating with a container registry (e.g., GHCR or Docker Hub).

## Acceptance Criteria

- [ ] **Dockerfile**: A `Dockerfile` exists at `infra/images/coder-base/Dockerfile`.
- [ ] **Toolchain**: The image includes:
  - Go (latest stable)
  - Rust (latest stable)
  - Python (3.11+)
  - Node.js (LTS)
  - Docker CLI (for Docker-in-Docker/socket mounting)
- [ ] **CI Pipeline**: A GitHub Actions workflow runs on push to `main`.
- [ ] **Build Process**: The workflow builds the Docker image.
- [ ] **Testing**: The workflow runs a smoke test (verifying tool versions) before pushing.
- [ ] **Artifacts**: The image is pushed to the container registry with `latest` and `sha-<commit>` tags.

## Technical Notes

- **Base Image**: Consider using Ubuntu 22.04 or 24.04 as the base.
- **Caching**: Utilize Docker layer caching in GitHub Actions to speed up builds.
- **Registry**: Ensure secrets are configured for registry authentication if using a private registry.
- **Multi-arch**: Ideally support `linux/amd64` and `linux/arm64` if possible (using `docker buildx`).
