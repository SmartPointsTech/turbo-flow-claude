# Story 2.1: Terraform Template Definition

Status: Done

## Story

As a Platform Engineer,
I want a standard Coder Terraform template,
So that all workspaces are consistent and managed as code.

## Acceptance Criteria

1. **Given** A new Coder deployment
2. **When** I upload the template from `infra/terraform/coder-template`
3. **Then** It validates successfully
4. **And** It exposes a `repo_url` parameter
5. **And** It defines a `coder_agent` resource that runs as the `coder` user
6. **And** It defines a `docker_container` resource using the "Fat Image"

## Tasks / Subtasks

- [x] Create Template Directory Structure
  - [x] Create `infra/terraform/coder-template/`
  - [x] Create `main.tf`, `variables.tf`, `outputs.tf`, `versions.tf`
- [x] Implement `variables.tf`
  - [x] Define `repo_url` (string, description: "Git repository URL to clone")
  - [x] Define `docker_image` (string, default: "codercom/enterprise-base:ubuntu")
- [x] Implement `versions.tf`
  - [x] Define `coder` provider
  - [x] Define `docker` provider
- [x] Implement `main.tf` - Coder Agent
  - [x] Resource `coder_agent` "main"
  - [x] Set `arch` and `os`
  - [x] Configure startup script
- [x] Implement `main.tf` - Docker Container
  - [x] Resource `docker_container` "workspace"
  - [x] Use `docker_image` variable
  - [x] Mount home volume (persistent)
  - [x] Connect agent via `env`
- [x] Verify Template
  - [x] Run `terraform init`
  - [x] Run `terraform validate`

## Dev Notes

### Architecture Compliance

- **IaC:** Terraform 1.5+
- **Providers:**
  - `coder/coder`
  - `kreuzwerker/docker`
- **Structure:** Standard Module Layout
- **Naming:** Snake_case for resources and variables.

### Technical Requirements

- **Fat Image:** For now, use a standard Coder image or a placeholder until Epic 3 builds the custom one.
- **Persistence:** Ensure `/home/coder` is mounted to a docker volume.
- **Agent:** Must run as the user inside the container (usually `coder` or `1000`).

### References

- [Architecture: Technology Stack](docs/architecture.md#1-technology-stack--versions)
- [Project Context: Naming Patterns](docs/project_context.md#3-naming--organization-patterns)

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 2.1)
- `docs/architecture.md`
- `docs/project_context.md`

### Completion Notes List

- Created Terraform template structure.
- Implemented `versions.tf` with Coder and Docker providers.
- Implemented `variables.tf` with `repo_url` and `docker_image`.
- Implemented `main.tf` with `coder_agent` and `docker_container`.
- Verified template with `terraform init` and `terraform validate`.

### File List

- `infra/terraform/coder-template/versions.tf`
- `infra/terraform/coder-template/variables.tf`
- `infra/terraform/coder-template/main.tf`
- `infra/terraform/coder-template/outputs.tf`
