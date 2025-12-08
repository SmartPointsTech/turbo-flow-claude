---
stepsCompleted: [1, 2, 3, 4]
status: 'complete'
completedAt: '2025-12-07'
---

# turbo-flow-claude - Epic Breakdown

## Overview

This document provides the complete epic and story breakdown for turbo-flow-claude, decomposing the requirements from the PRD, UX Design if it exists, and Architecture requirements into implementable stories.

## Requirements Inventory

### Functional Requirements

FR1: Users can provision a new workspace using the `spt-flow` Terraform template.
FR2: Users can specify a target Git repository URL at provisioning time.
FR3: The system must clone the specified repository into the workspace automatically on startup.
FR4: The system must handle authentication challenges for private repositories during the validation phase.
FR5: The workspace must support Nested Virtualization (KVM access) for running Firecracker/Kuasar.
FR6: The workspace must include a pre-installed "Fat Image" toolchain (Go, Rust, Python, Node.js, Docker).
FR7: The workspace must persist user data in the `/home/coder` directory across restarts.
FR8: The workspace must provide a web-based VS Code instance accessible via browser.
FR9: Users can initialize the local client configuration via `sptCoder init`.
FR10: Users can create or connect to a workspace for a specific repo via a single command (`sptCoder up [repo]`).
FR11: Users can SSH into a running workspace via `sptCoder ssh`.
FR12: The CLI must automatically check for version updates against the server.
FR13: The CLI must support installation on Linux, macOS, and Windows (PowerShell & WSL2).
FR14: Users can mount a personal "dotfiles" repository to customize their shell environment.
FR15: The system must forward the user's local SSH keys to the workspace.
FR16: The system must forward the user's local Git credentials (via `gh` CLI or credential helper) to the workspace.
FR17: The system must automatically stop workspaces after a configurable period of inactivity (default: 2 hours).
FR18: CI/CD pipelines can request ephemeral workspaces with a strict Time-To-Live (TTL).
FR19: Users can manually restart a workspace to apply template updates.
FR20: Users can enable a "Keep-Alive" override to prevent auto-stop for long-running tasks (up to 24h).

### NonFunctional Requirements

NFR1: (Startup Time) A new workspace must be "Interactive" (VS Code accessible) within 2 minutes.
NFR2: (CLI Latency) CLI commands (`sptCoder up`) must respond to user input within 200ms.
NFR3: (Image Caching) The "Fat Image" must be pre-pulled/cached on all Coder nodes.
NFR4: (Privilege Management) Passwordless `sudo` access with audit logging.
NFR5: (Isolation) Each workspace must be isolated in its own Kubernetes Pod/Docker container.
NFR6: (Secret Management) SSH keys and Git credentials must be injected via memory-only agents.
NFR7: (Availability) 99.5% availability during Business Hours.
NFR8: (Data Persistence) User data in `/home/coder` must survive container restarts.
NFR9: (Backup Strategy) Daily snapshots with 7-day retention.

### Additional Requirements

- **Starter Template:** Scaffold from Standard Patterns (Monorepo: `cli/`, `infra/`).
- **State Strategy:** Stateless CLI (Query-on-demand).
- **Auth Strategy:** SSH Agent Forwarding (No disk persistence).
- **Observability:** Distributed Tracing via Correlation IDs.
- **Naming Patterns:** Kebab-case (CLI), Snake_case (Terraform).
- **Structure:** `cli/cmd`, `cli/internal`, `infra/terraform`, `infra/images`.
- **Testing:** `test/e2e/` for integration testing.
- **Windows Risk:** Support Windows Named Pipes for SSH Agent.

### FR Coverage Map

FR1: Epic 2 - Provision workspace
FR2: Epic 2 - Specify repo URL
FR3: Epic 4 - Clone repo (Personalization/Setup)
FR4: Epic 4 - Auth for private repos
FR5: Epic 3 - Nested Virt
FR6: Epic 3 - Fat Image
FR7: Epic 4 - Data persistence
FR8: Epic 3 - Web VS Code
FR9: Epic 1 - CLI Init
FR10: Epic 2 - CLI Up
FR11: Epic 2 - CLI SSH
FR12: Epic 1 - CLI Updates
FR13: Epic 1 - Cross-platform
FR14: Epic 4 - Dotfiles
FR15: Epic 4 - SSH Keys
FR16: Epic 4 - Git Creds
FR17: Epic 2 - Auto-stop
FR18: Epic 2 - TTL
FR19: Epic 2 - Restart
FR20: Epic 2 - Keep-Alive

## Epic List

### Epic 1: CLI Foundation & Configuration

- **Goal:** Users can install the CLI, initialize configuration, and manage their local setup.

- **FRs covered:** FR9, FR12, FR13

### Epic 2: Core Workspace Operations

- **Goal:** Users can create, start, stop, and connect to workspaces using the standard template.

- **FRs covered:** FR1, FR2, FR10, FR11, FR17, FR18, FR19, FR20

### Epic 3: Developer Experience & Image Pipeline

- **Goal:** Users have a fully equipped environment (Nested Virt, Tools) backed by an automated build pipeline.

- **FRs covered:** FR5, FR6, FR8

### Epic 4: Personalization & Security

- **Goal:** Users can bring their own identity (keys, dotfiles) and persist data securely.

- **FRs covered:** FR3, FR4, FR7, FR14, FR15, FR16

## Epic 1: CLI Foundation & Configuration

Users can install the CLI, initialize configuration, and manage their local setup.

### Story 1.1: CLI Skeleton & Cross-Platform Build

As a Developer,
I want a CLI binary that runs on my specific Operating System,
So that I can manage my workspaces regardless of my machine type.

**Acceptance Criteria:**

**Given** The source code is compiled via `goreleaser`
**When** I run the binary on Linux (amd64), macOS (arm64), and Windows (amd64)
**Then** The command `sptCoder --version` prints the correct version and commit hash
**And** The command `sptCoder --help` prints the root help text
**And** The binary has no external runtime dependencies (static linking)

### Story 1.2: Configuration Initialization (`init`)

As a User,
I want to initialize my local configuration interactively,
So that I don't have to provide the Coder URL and token for every command.

**Acceptance Criteria:**

**Given** I have a fresh installation
**When** I run `sptCoder init`
**Then** It prompts me for the Coder Deployment URL and Session Token
**And** It validates that the URL is reachable
**And** It saves the configuration to `~/.spt-flow.yaml` (or XDG config path)
**And** Subsequent commands read these values automatically

### Story 1.3: Automated Version Check

As a Platform Engineer,
I want the CLI to check for updates automatically,
So that users are aware when their tool is out of sync with the platform.

**Acceptance Criteria:**

**Given** An outdated local CLI version
**When** I run any command (e.g., `sptCoder up`)
**Then** The CLI queries the release source (GitHub Releases) in the background
**And** It prints a warning to stderr: "Update available: vX.Y.Z -> vA.B.C"
**And** It respects a `SPT_FLOW_SKIP_UPDATE_CHECK=true` env var to disable this behavior

## Epic 2: Core Workspace Operations

Users can create, start, stop, and connect to workspaces using the standard template.

### Story 2.1: Terraform Template Definition

As a Platform Engineer,
I want a standard Coder Terraform template,
So that all workspaces are consistent and managed as code.

**Acceptance Criteria:**

**Given** A new Coder deployment
**When** I upload the template from `infra/terraform/coder-template`
**Then** It validates successfully
**And** It exposes a `repo_url` parameter
**And** It defines a `coder_agent` resource that runs as the `coder` user
**And** It defines a `docker_container` resource using the "Fat Image"

### Story 2.2: Workspace Provisioning (`up`)

As a Developer,
I want to create or start a workspace for a specific repo with one command,
So that I can get to work immediately without UI clicks.

**Acceptance Criteria:**

**Given** I am authenticated via `sptCoder`
**When** I run `sptCoder up https://github.com/my/repo`
**Then** It checks if a workspace named `repo-name` exists
**And** If missing, it creates it using the `spt-flow` template and `repo_url` param
**And** If stopped, it starts it
**And** It waits for the agent to be ready (streaming logs)
**And** It prints the VS Code Web URL to stdout (or opens browser if `--open` flag is set)

### Story 2.3: SSH Connection Handler (`ssh`)

As a Developer,
I want to SSH into my workspace securely,
So that I can use my local terminal tools and forward my keys.

**Acceptance Criteria:**

**Given** A running workspace
**When** I run `sptCoder ssh repo-name`
**Then** It connects via the Coder tunnel (no public IP needed)
**And** It forwards my local `ssh-agent` socket to the container
**And** It does NOT write any private keys to disk inside the container
**And** I land in a shell as the `coder` user

### Story 2.4: Lifecycle Management (TTL & Auto-Stop)

As a FinOps Manager,
I want unused workspaces to stop automatically,
So that we don't pay for idle compute.

**Acceptance Criteria:**

**Given** A running workspace
**When** It has been idle (no SSH/VS Code connection) for 2 hours
**Then** It stops automatically
**And** When I run `sptCoder up --ttl 30m`
**Then** The workspace is hard-stopped after 30 minutes (for CI jobs)
**And** I can run `sptCoder stop` to stop it immediately

## Epic 3: Developer Experience & Image Pipeline

Users have a fully equipped environment (Nested Virt, Tools) backed by an automated build pipeline.

### Story 3.1: "Fat Image" Definition & Build Pipeline

As a Platform Engineer,
I want an automated pipeline that builds our standard developer image,
So that everyone uses the same tooling and updates are safe.

**Acceptance Criteria:**

**Given** The `infra/images/coder-base/Dockerfile`
**When** I push a change to the `main` branch
**Then** GitHub Actions builds the image
**And** It runs a smoke test (e.g., `go version`, `node --version`)
**And** It pushes the image to our container registry with `latest` and `sha-tag`
**And** The image includes Go, Rust, Python, Node.js, and Docker

### Story 3.2: Nested Virtualization Configuration

As a Developer,
I want to run KVM-based workloads (Firecracker) inside my workspace,
So that I can test microVMs without needing a bare-metal laptop.

**Acceptance Criteria:**

**Given** A workspace provisioned with the `spt-flow` template
**When** I run `kvm-ok` inside the terminal
**Then** It returns "KVM acceleration can be used"
**And** I can run a simple Docker container
**And** I can run a Firecracker microVM (hello-world)
**And** The `/dev/kvm` device is accessible by the `coder` user

### Story 3.3: Web IDE & Extensions

As a Developer,
I want a fully configured VS Code environment in my browser,
So that I can code from any device without installing local tools.

**Acceptance Criteria:**

**Given** A running workspace
**When** I click the "Open VS Code" button in the Coder dashboard (or via CLI)
**Then** VS Code opens in my browser
**And** The "Go", "Rust Analyzer", and "Docker" extensions are pre-installed
**And** My terminal shell defaults to `zsh` (or preferred shell)
**And** I can access the file system and run terminals

## Epic 4: Personalization & Security

Users can bring their own identity (keys, dotfiles) and persist data securely.

### Story 4.1: Dotfiles & Personalization Hook

As a Developer,
I want my workspace to load my personal shell configuration (dotfiles),
So that I have my aliases, prompt, and tools ready immediately.

**Acceptance Criteria:**

**Given** I have a public dotfiles repository (e.g., GitHub)
**When** I provision a workspace with `sptCoder up --dotfiles https://github.com/me/dotfiles`
**Then** The system clones the repo to `~/dotfiles`
**And** It detects and runs the install script (`install.sh`, `setup.sh`, or `bootstrap.sh`)
**And** If the script fails, the workspace still starts but logs a warning
**And** My shell environment reflects the changes (e.g., `.zshrc` is sourced)

### Story 4.2: Secret Injection (SSH & Git)

As a Developer,
I want to use my local SSH keys and Git credentials inside the workspace,
So that I can push code and access private servers without copying keys manually.

**Acceptance Criteria:**

**Given** I have SSH keys loaded in my local agent (`ssh-add -l`)
**When** I SSH into the workspace (or open VS Code)
**Then** `ssh-add -l` inside the workspace shows my local keys
**And** I can run `git push` to a private repo without entering a password
**And** No private keys are stored on the workspace disk (`~/.ssh/id_rsa` does not exist)

### Story 4.3: Data Persistence Strategy

As a User,
I want my project files and home directory to survive a workspace restart,
So that I don't lose work if the container is recreated.

**Acceptance Criteria:**

**Given** I have created files in `/home/coder`
**When** I run `sptCoder stop` and then `sptCoder start`
**Then** The files are still present
**And** If the underlying node changes (pod rescheduling), the data follows
**And** The persistence is backed by a PVC (Persistent Volume Claim) in Kubernetes
