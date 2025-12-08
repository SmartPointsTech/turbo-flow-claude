# Architecture Documentation

## Executive Summary

**Turbo Flow Claude** is an advanced agentic development environment designed to streamline AI-assisted coding. It leverages **DevPod** for reproducible cloud environments and **Claude Flow** for orchestrating AI agents. The architecture is primarily script-based, focusing on automation, environment provisioning, and tool integration.

## Technology Stack

| Category | Technology | Description |
|----------|------------|-------------|
| **Core** | Bash | Primary scripting language for setup and automation. |
| **Environment** | DevPod | Infrastructure-as-Code tool for creating reproducible dev environments. |
| **Containerization** | Docker | Used for containerizing the development environment. |
| **AI Orchestration** | Claude Flow | Framework for managing and coordinating AI agents. |
| **Package Managers** | npm, uv, cargo | Used to install various tools and dependencies. |

## Architecture Pattern

The project follows a **Script-based Automation** pattern. It does not have a traditional application architecture (MVC, Microservices) but rather a collection of utility scripts and configuration files that work together to provision a specific state (the development environment).

### Key Components

1. **Bootstrapping Layer**: `devpods/boot_*.sh` scripts that handle the initial OS-specific setup (macOS, Linux, Cloud Shell).
2. **Provisioning Layer**: `devpods/setup.sh` which installs core tools, configures the shell, and sets up the AI environment.
3. **Configuration Layer**: JSON and YAML files (`devcontainer.json`, `mcp.json`) that define tool settings and environment variables.
4. **Agent Layer**: A library of AI agents (`agents/` directory, installed dynamically) that provide specific capabilities within the environment.

## Data Architecture

As an infrastructure project, it does not manage a traditional database. However, it manages:

- **State**: `project-scan-report.json` and other status files track the state of workflows.
- **Configuration**: User preferences and tool configs are stored in dotfiles and JSON configs.

## Deployment Architecture

Deployment is handled via **DevPod**, which abstracts the underlying infrastructure provider (AWS, DigitalOcean, Local Docker). The `devcontainer.json` serves as the blueprint for the deployed environment.
