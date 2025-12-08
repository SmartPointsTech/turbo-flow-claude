# Deployment Guide

## Infrastructure as Code (DevPod)

This project uses **DevPod** to define and deploy development environments across various cloud providers.

### Supported Providers

| Provider | Configuration Command | Key Options |
|----------|----------------------|-------------|
| **DigitalOcean** | `devpod provider add digitalocean` | `DROPLET_SIZE`, `DIGITALOCEAN_ACCESS_TOKEN` |
| **AWS** | `devpod provider add aws` | `AWS_INSTANCE_TYPE`, `AWS_REGION` |
| **Azure** | `devpod provider add azure` | `AZURE_VM_SIZE`, `AZURE_LOCATION` |
| **Google Cloud** | `devpod provider add gcp` | `GOOGLE_PROJECT_ID`, `GOOGLE_MACHINE_TYPE` |

### Deployment Process

1. **Configure Provider**:

    ```bash
    devpod provider use [provider_name]
    ```

2. **Deploy Workspace**:

    ```bash
    devpod up https://github.com/marcuspat/turbo-flow-claude --ide vscode
    ```

3. **Stop/Resume**:

    ```bash
    devpod stop turbo-flow-claude
    devpod up turbo-flow-claude
    ```

## CI/CD Pipeline

- **GitHub Actions**: The `.github/workflows` directory is currently minimal (Funding only), but the project includes scripts for CI/CD integration.
- **QA Pipeline**: `devpods/qa_pipeline_boot.sh` is available for bootstrapping QA environments.

## Docker Configuration

- **DevContainer**: `.devcontainer/devcontainer.json` (or `devpods/rackspace-devcontainer.json`) defines the container image and extensions.
- **Dockerfile**: Used by DevPod to build the environment.
