---
stepsCompleted: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
inputDocuments:
  - /home/prichard/repos/turbo-flow-claude/docs/analysis/brainstorming-session-2025-12-07.md
  - /home/prichard/repos/turbo-flow-claude/docs/index.md
  - /home/prichard/repos/turbo-flow-claude/docs/project-overview.md
  - /home/prichard/repos/turbo-flow-claude/docs/architecture.md
documentCounts:
  briefs: 0
  research: 0
  brainstorming: 1
  projectDocs: 3
workflowType: 'prd'
lastStep: 0
project_name: 'turbo-flow-claude'
user_name: 'Philippe'
date: '2025-12-07'
---

# Product Requirements Document - turbo-flow-claude

**Author:** Philippe
**Date:** 2025-12-07

## Executive Summary

**spt-flow** (formerly turbo-flow-claude) represents a strategic evolution of our development platform. We are transitioning from a client-side, DevPod-centric model to a **centralized, server-first architecture using Coder**.

This shift establishes Coder as the **default target environment** for all team members and CI/CD pipelines, ensuring consistent, reproducible, and powerful workspaces. Crucially, we maintain strict backward compatibility with DevPod and Codespaces, allowing for ad-hoc experimentation and edge-case support without fragmenting the codebase.

### What Makes This Special

**Unified Platform Architecture:**
Unlike typical "rip and replace" migrations, spt-flow implements a **"Single Source, Multi-Target"** strategy. It leverages a shared configuration to deliver the enterprise-grade benefits of centralized management (standardization, persistence, nested virtualization) while preserving the developer's freedom to spin up ephemeral local environments. It unifies the **"Team Standard"** (Coder) with the **"Hacker's Lab"** (DevPod) in a single, cohesive platform.

## Project Classification

**Technical Type:** Infrastructure / DevOps
**Domain:** Developer Tools / DevOps
**Complexity:** High
**Project Context:** Brownfield - extending existing system

**Classification Notes:**
This project is classified as **High Complexity** due to the requirement for **Nested Virtualization** (enabling Firecracker/WASM testing within Coder) and the architectural challenge of maintaining **Multi-Target Support** from a single configuration source.

## Success Criteria

### User Success

- **Speed:** "Time-to-Interactive" < 2 minutes; "Time-to-Full-Build" < 5 minutes.
- **Capability:** "Firecracker-Ready" - Users can run nested virtualization workloads (WASM, microVMs) immediately upon boot.
- **Integration:** "Pipeline-Native" - Environment mirrors CI/CD context, ensuring "works on my machine" means "works in prod".

### Business Success

- **CI/CD Velocity:** Zero configuration drift between dev and CI environments.
- **Standardization:** 100% adoption of the shared `spt-flow` template for team development.

### Technical Success

- **Nested Virtualization:** `kvm-ok` returns success; `/dev/kvm` is accessible to user.
- **Secure Context:** SSH keys and Git credentials are securely forwarded to the workspace (no manual key copying).
- **Repo Injection:** Target repository is cloned with correct permissions and hooks.

### Measurable Outcomes

- **Drift:** Zero environment-related failures in CI/CD pipelines for projects using Coder.
- **Validation:** 100% pass rate for `spt-flow` validation scripts (including nested virt checks).

## Product Scope

### MVP - Minimum Viable Product

- **Core Template:** Terraform template for `spt-flow` with Coder provider.
- **Infrastructure:** Ubuntu 22.04 base with KVM/Nested Virtualization enabled.
- **Security:** Agent injection for SSH/GPG forwarding.
- **Provisioning:** Startup script to clone the user-specified repository.
- **Toolchain:** Docker (in-docker), git, gh CLI, and basic WASM tools.

### Growth Features (Post-MVP)

- **Full Parity:** Complete port of all `devpods/setup.sh` tools and aliases.
- **Performance:** Pre-warmed Docker images to reduce "Time-to-Interactive" to < 30 seconds.
- **Multi-Repo:** Support for mounting multiple repositories (microservices pattern).

### Vision (Future)

- **Unified Ops:** Coder serves as the single pane of glass for all development, testing, and CI/CD infrastructure management.

## User Journeys

### Journey 1: Alex (The Developer) - "Zero to Hero"

Alex needs to test a Firecracker microVM. Instead of spending hours on local config, Alex logs into Coder.
**The Magic:** Alex selects the "spt-flow" template. No sizing options, no config flags—just "Start." The defaults (4 vCPU, 8GB RAM, Nested Virt) are pre-tuned by the platform team.
Within 2 minutes, the IDE opens. `kvm-ok` returns success. Alex runs the test, it passes.
*Crucially:* When Alex is done, they don't even have to shut it down. The workspace has an "Auto-Stop" policy of 2 hours, saving costs automatically.

### Journey 2: Sam (The Platform Engineer) - "Safe Evolution"

Sam needs to roll out a new version of the `gh` CLI to the team. Sam updates the Terraform template and pushes to git.
**The Magic:** Coder detects the template change. It *doesn't* force-restart Alex's active workspace (preventing data loss). Instead, it shows a "Update Available" badge.
When Alex chooses to update (or on the next rebuild), the new tool is there. Sam can deprecate old templates without breaking active work.

### Journey 3: CI-Bot (The Pipeline) - "Ephemeral & Efficient"

The CI pipeline triggers. It calls the Coder API to provision a workspace with a `ttl=30m` flag.
**The Magic:** The workspace spins up, runs the heavy integration tests (which require KVM), and reports the status.
*Safety Net:* Even if the pipeline crashes and fails to send the "destroy" command, the 30-minute TTL ensures the expensive instance is killed automatically. No zombie cloud bills.

### Journey Requirements Summary

**Capabilities Revealed:**

- **Smart Defaults:** "One-click" provisioning requires opinionated, pre-set resource sizing.
- **Lifecycle Management:** Auto-stop (for devs) and Auto-TTL (for CI) to control costs.
- **Safe Updates:** Template versioning that prompts users rather than forcing restarts.
- **Nested Virt Support:** Validated as a core requirement for all personas.

## Developer Tool Specific Requirements

### Project-Type Overview

`spt-flow` functions as a hybrid **Infrastructure Platform** and **Developer CLI**. It provides a standardized remote development environment (via Coder) accessed and managed through a custom desktop CLI tool (`sptCoder`).

### Technical Architecture Considerations

**1. IDE & Access Support**

- **Primary:** VS Code Server (Browser & Desktop).
- **Secondary:** SSH (Terminal access).
- **Excluded (MVP):** JetBrains Gateway / Projector.

**2. Language & Tool Stack Strategy**

- **Strategy:** "Optimized Fat Image" with Layer Caching.
- **Decision:** Bake-in Core Runtimes to ensure "Time-to-Interactive" < 2 mins.
- **Base Image Contents:**
  - Ubuntu 22.04 LTS
  - Go (Latest), Rust (Latest), Python (Latest Stable), Node.js (LTS)
  - WASM Toolchain (wasm-pack, etc.)
  - Docker (for Nested Virt)
- **Performance Requirement:** Base image must be pre-pulled/cached on Coder nodes to mitigate download time.

**3. Client-Side Tooling: `sptCoder` CLI**

- **Architecture:** Standalone Go binary (not a bash wrapper).
- **Philosophy:** "Workflow Orchestrator" - handles `spt-flow` specifics, delegates raw connection to `coder` CLI.
- **Key Commands:**
  - `sptCoder init`: Interactive setup (checks dependencies, configures defaults).
  - `sptCoder up [repo]`: Idempotent "create or connect" command.
  - `sptCoder ssh`: Smart SSH wrapper.
- **Distribution:** Cross-platform support for **Linux, macOS, and Windows**.
  - Linux/macOS: Homebrew Tap (`brew install spt-flow/tap/sptcoder`).
  - Windows: Scoop bucket or direct binary (`scoop install sptcoder`).

### Implementation Considerations

- **Image Build Pipeline:** Automated weekly builds to keep runtimes fresh.
- **CLI Versioning:** CLI must check for updates on run to prevent "drift" between client tools and server templates.

## Project Scoping & Phased Development

### MVP Strategy & Philosophy

**MVP Approach:** Platform MVP
**Philosophy:** Build the robust "Golden Path" foundation first. Focus on reliability and core developer experience (CLI + Image).
**Resource Requirements:** Small Core Team (1-2 Engineers) focusing on Terraform, Go CLI, and Docker Image automation.

### MVP Feature Set (Phase 1)

**Core User Journeys Supported:**

- **Alex (Developer):** "Zero to Hero" provisioning of a standard workspace.
- **Sam (Platform Eng):** Updating the "Fat Image" and CLI.
- **CI-Bot:** Ephemeral workspace creation via CLI.

**Must-Have Capabilities:**

- **Core Infrastructure:** Coder Terraform Template with KVM/Nested Virt enabled.
- **Universal CLI:** `sptCoder` binary for Windows, macOS, and Linux.
- **Optimized Runtime:** "Fat Image" with Go, Rust, Python, Node.js pre-installed.
- **Personalization:** Support for **User Dotfiles** (mounting a personal repo) to handle aliases/preferences cleanly.

### Post-MVP Features

**Phase 2 (Growth - "The Parity Push"):**

- **Legacy Parity:** Porting shared team aliases to the core image.
- **Advanced Multi-Repo:** Support for complex microservice layouts.
- **Performance Tuning:** Pre-warming workspaces.

**Phase 3 (Expansion):**

- **IDE Choice:** JetBrains Gateway / Projector support.
- **Unified Dashboard:** Integrating CI/CD status directly into the Coder dashboard.

### Risk Mitigation Strategy

**Technical Risks:**

- *Risk:* Nested Virtualization performance is poor on some cloud providers.
- *Mitigation:* Validate `kvm-ok` on target providers immediately. Fallback to software emulation if needed.

**Adoption Risks:**

- *Risk:* Developers reject the new environment due to missing custom workflows.
- *Mitigation:* **Dotfiles Support** (MVP) allows immediate personalization. We will also track "missing feature" requests via a CLI command (`sptCoder feedback`) to prioritize Phase 2.

## Functional Requirements

### Workspace Provisioning

- **FR1:** Users can provision a new workspace using the `spt-flow` Terraform template.
- **FR2:** Users can specify a target Git repository URL at provisioning time.
- **FR3:** The system must clone the specified repository into the workspace automatically on startup.
- **FR4:** The system must handle authentication challenges for private repositories during the validation phase.

### Environment Capabilities

- **FR5:** The workspace must support Nested Virtualization (KVM access) for running Firecracker/Kuasar.
- **FR6:** The workspace must include a pre-installed "Fat Image" toolchain (Go, Rust, Python, Node.js, Docker).
- **FR7:** The workspace must persist user data in the `/home/coder` directory across restarts.
- **FR8:** The workspace must provide a web-based VS Code instance accessible via browser.

### CLI Operations (`sptCoder`)

- **FR9:** Users can initialize the local client configuration via `sptCoder init`.
- **FR10:** Users can create or connect to a workspace for a specific repo via a single command (`sptCoder up [repo]`).
  - *Constraint:* This action performs a fresh clone of the remote HEAD (no local sync in MVP).
- **FR11:** Users can SSH into a running workspace via `sptCoder ssh`.
- **FR12:** The CLI must automatically check for version updates against the server.
- **FR13:** The CLI must support installation on Linux, macOS, and Windows (PowerShell & WSL2).

### Personalization & Identity

- **FR14:** Users can mount a personal "dotfiles" repository to customize their shell environment.
- **FR15:** The system must forward the user's local SSH keys to the workspace.
- **FR16:** The system must forward the user's local Git credentials (via `gh` CLI or credential helper) to the workspace.

### Lifecycle Management

- **FR17:** The system must automatically stop workspaces after a configurable period of inactivity (default: 2 hours).
- **FR18:** CI/CD pipelines can request ephemeral workspaces with a strict Time-To-Live (TTL).
- **FR19:** Users can manually restart a workspace to apply template updates.
- **FR20:** Users can enable a "Keep-Alive" override to prevent auto-stop for long-running tasks (up to 24h).

## Non-Functional Requirements

### Performance

- **NFR1 (Startup Time):** A new workspace must be "Interactive" (VS Code accessible) within **2 minutes** of provisioning (assuming cached image).
- **NFR2 (CLI Latency):** CLI commands (`sptCoder up`) must respond to user input within **200ms**.
- **NFR3 (Image Caching - CRITICAL):** The "Fat Image" must be pre-pulled/cached on all Coder nodes. Cache misses are considered a system failure.

### Security

- **NFR4 (Privilege Management):** The `coder` user has passwordless `sudo` access to facilitate development, but all privileged commands must be logged to the system audit trail.
- **NFR5 (Isolation):** Each workspace must be isolated in its own Kubernetes Pod or Docker container with strict resource limits.
- **NFR6 (Secret Management):** SSH keys and Git credentials must be injected via memory-only agents (ssh-agent) and never written to disk.

### Reliability

- **NFR7 (Availability):** System targets 99.5% availability during Business Hours (9am-5pm). Maintenance windows are scheduled outside these hours.
- **NFR8 (Data Persistence):** User data in `/home/coder` must survive container restarts.
- **NFR9 (Backup Strategy):** Persistent volumes must be snapshotted daily with a retention policy of 7 days to protect against corruption/deletion.
