variable "arch" {
  description = "Architecture of the workspace"
  type        = string
  default     = "amd64"
}

variable "os" {
  description = "Operating system of the workspace"
  type        = string
  default     = "linux"
}

variable "repo_url" {
  description = "URL of the repository to clone"
  type        = string
  default     = ""
}

variable "docker_image" {
  description = "Docker image to use for the workspace"
  type        = string
  default     = "coder-base:latest"
}

variable "enable_nested_virt" {
  description = "Enable Nested Virtualization (privileged mode + /dev/kvm)"
  type        = bool
  default     = true
}

variable "dotfiles_url" {
  description = "URL of the dotfiles repository to clone"
  type        = string
  default     = ""
}
