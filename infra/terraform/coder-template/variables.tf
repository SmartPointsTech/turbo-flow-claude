variable "repo_url" {
  description = "Git repository URL to clone into the workspace"
  type        = string
  default     = ""
}

variable "docker_image" {
  description = "Docker image to use for the workspace"
  type        = string
  default     = "codercom/enterprise-base:ubuntu"
}

variable "arch" {
  description = "Architecture of the workspace (amd64 or arm64)"
  type        = string
  default     = "amd64"
}

variable "os" {
  description = "Operating system of the workspace (linux, darwin, windows)"
  type        = string
  default     = "linux"
}
