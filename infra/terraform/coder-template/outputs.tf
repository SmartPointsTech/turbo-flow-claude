output "workspace_url" {
  value = "https://coder.example.com/@${data.coder_workspace.me.owner}/${data.coder_workspace.me.name}"
}

output "ssh_command" {
  value = "ssh coder.${data.coder_workspace.me.name}"
}
