resource "coder_agent" "main" {
  arch           = var.arch
  os             = var.os
  startup_script = <<EOT
    #!/bin/bash
    set -euo pipefail

    # Clone repository if provided
    if [ -n "${var.repo_url}" ]; then
      REPO_DIR="$HOME/$(basename "${var.repo_url}" .git)"
      if [ ! -d "$REPO_DIR" ]; then
        echo "Cloning ${var.repo_url}..."
        git clone "${var.repo_url}" "$REPO_DIR"
      fi
    fi

    # Install code-server
    curl -fsSL https://code-server.dev/install.sh | sh
    code-server --auth none --port 13337 >/dev/null 2>&1 &
  EOT
}

resource "docker_volume" "home_volume" {
  name = "coder-${data.coder_workspace.me.id}-home"
}

resource "docker_container" "workspace" {
  count = data.coder_workspace.me.start_count
  image = var.docker_image
  name  = "coder-${data.coder_workspace.me.owner}-${data.coder_workspace.me.name}"
  # Uses 127.0.0.1:13337 for code-server
  # Uses 127.0.0.1:22 for SSH
  
  # Hostname makes the shell more user friendly: coder@my-workspace:~$
  hostname = data.coder_workspace.me.name
  
  # Use the docker gateway if the access URL is 127.0.0.1
  entrypoint = ["sh", "-c", coder_agent.main.init_script]
  env        = ["CODER_AGENT_TOKEN=${coder_agent.main.token}"]
  
  host {
    host = "host.docker.internal"
    ip   = "host-gateway"
  }

  volumes {
    container_path = "/home/coder"
    volume_name    = docker_volume.home_volume.name
    read_only      = false
  }
}

data "coder_workspace" "me" {}
