# Source Tree Analysis

## Project Structure

This project is structured as a **Monolith** focused on Infrastructure and DevOps automation.

### Directory Tree

```
turbo-flow-claude/
├── devpods/                     # Critical: Development environment configurations and setup scripts
│   ├── setup.sh                 # Primary setup script for the environment
│   ├── *.sh                     # Various boot and utility scripts
│   └── *.json                   # Configuration files for DevContainer and settings
├── bmad-custom-modules-src/     # Custom BMad modules source (currently empty)
├── docs/                        # Project documentation and scan reports
├── .github/                     # GitHub configuration (Funding)
└── [Root Files]                 # Extensive documentation guides (GUIAS_*, READMEs)
```

### Critical Directories

| Directory | Purpose | Key Files |
|-----------|---------|-----------|
| `devpods/` | Contains all logic for provisioning and configuring development environments (DevPods). This is the core functional part of the repo. | `setup.sh`, `turbo-flow-wizard.sh`, `rackspace-devcontainer.json` |
| `docs/` | Stores generated documentation and project analysis reports. | `project-scan-report.json` |
| `[Root]` | Contains high-level guides and aliases configuration. | `GUIAS_CONFIGURACION_TURBO_FLOW.md`, `claude-flow-aliases-guide.md` |

### Entry Points

- **Environment Setup**: `devpods/setup.sh` - The main script to initialize the development environment.
- **Wizard**: `devpods/turbo-flow-wizard.sh` - Interactive setup wizard.
- **Aliases**: `claude-flow-aliases-guide.md` - Documentation for the aliases provided by the system.

### Integration Points

As an infrastructure project, integration points are primarily:

- **External Tools**: npm, uv, cargo, claude-code (installed via scripts)
- **Cloud Providers**: Scripts for Rackspace, Google Cloud Shell, Codespaces (`devpods/boot_*.sh`)
