# Development Guide

## Prerequisites

- **Git**: Required for cloning the repository.
- **DevPod** (Optional but recommended): For cloud development environments.
- **Node.js**: Required for local setup (LTS version).
- **Docker**: Required for local DevPod or containerized setup.

## Environment Setup

### Quick Start (DevPod)

1. **Install DevPod**:
    - macOS: `brew install loft-sh/devpod/devpod`
    - Windows: `choco install devpod`
    - Linux: `curl -L -o devpod "https://github.com/loft-sh/devpod/releases/latest/download/devpod-linux-amd64" && sudo install devpod /usr/local/bin`

2. **Launch Workspace**:

    ```bash
    devpod up https://github.com/marcuspat/turbo-flow-claude --ide vscode
    ```

### GitHub Codespaces

1. Create a new Codespace from the repository.
2. Run the boot script:

    ```bash
    touch boot.sh && chmod +x boot.sh && vi boot.sh
    # Paste content from GUIAS_CONFIGURACION_TURBO_FLOW.md
    bash boot.sh
    ```

3. Connect to tmux: `tmux attach -t workspace`

### Local Setup (macOS/Linux)

1. Clone the repository:

    ```bash
    git clone https://github.com/marcuspat/turbo-flow-claude.git
    cd turbo-flow-claude/devpods
    ```

2. Run the boot script:
    - macOS: `./boot_macosx.sh`
    - Linux: `./boot_linux.sh`
3. Attach to the workspace: `tmux attach -t workspace`

## Development Workflow

### Tmux Workspace

The environment uses `tmux` with 4 pre-configured windows:

- **0: Claude-1**: Main workspace
- **1: Claude-2**: Secondary workspace
- **2: Claude-Monitor**: Usage tracking
- **3: htop**: System monitor

### Common Commands

- **Claude Flow**: `cf "task"`
- **Swarm**: `cf-swarm "task"`
- **Hive Mind**: `cf-hive "task"`
- **Optimization**: `af-optimize --agent coder --task "task"`

## Testing

- **Playwright**: Installed for UI testing.
- **Agentic-Flow**: Includes a `tester` agent (`af-tester`).
