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
    
    # Dotfiles personalization
    if [ -n "${var.dotfiles_url}" ]; then
      echo "Installing dotfiles from ${var.dotfiles_url}..."
      DOTFILES_DIR="$HOME/dotfiles"
      if [ ! -d "$DOTFILES_DIR" ]; then
        git clone "${var.dotfiles_url}" "$DOTFILES_DIR"
      fi

      if [ -d "$DOTFILES_DIR" ]; then
        # Look for install scripts
        for script in install.sh setup.sh bootstrap.sh; do
          if [ -f "$DOTFILES_DIR/$script" ]; then
            echo "Executing $script..."
            chmod +x "$DOTFILES_DIR/$script"
            # execute in subshell, ignore failure
            (cd "$DOTFILES_DIR" && ./$script) || echo "Dotfiles script failed, continuing..."
            break
          fi
        done
      fi
    fi

    # Install code-server (pre-installed in image, just launch)
    code-server --auth none --port 13337 >/dev/null 2>&1 &
  EOT
}

resource "docker_volume" "home_volume" {
  name = "coder-${data.coder_workspace.me.id}-home"
  # Lifecycle: This volume perists across workspace stops/starts and rebuilds.
  # Data in /home/coder is safe.
  lifecycle {
    prevent_destroy = false # Allow manual destruction if needed, but defaults preserve it.
  }
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
  
  # Nested Virtualization Support
  privileged = var.enable_nested_virt

  dynamic "devices" {
    for_each = var.enable_nested_virt ? [1] : []
    content {
      host_path      = "/dev/kvm"
      container_path = "/dev/kvm"
    }
  }

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
