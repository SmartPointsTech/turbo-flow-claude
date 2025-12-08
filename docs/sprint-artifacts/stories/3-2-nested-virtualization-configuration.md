# Story 3.2: Nested Virtualization Configuration

**Status:** Ready for Review
**Estimate:** 5 Points
**Priority:** Medium
**Assignee:** AI Agent

## Story

As a Developer,
I want to run KVM-based workloads (Firecracker) inside my workspace,
So that I can test microVMs without needing a bare-metal laptop.

## Context

Our developers need to test virtualization-based workloads, specifically Firecracker. This requires "Nested Virtualization" to be enabled on the underlying host and passed through to the workspace container.
This story focuses on the Terraform configuration and checking potential host-level requirements.

## Acceptance Criteria

- [ ] **KVM Enabled**: `kvm-ok` inside the workspace returns "KVM acceleration can be used".
- [ ] **Permissions**: The `coder` user can access `/dev/kvm` (usually group `kvm` or `render`).
- [ ] **Docker Support**: Running a standard Docker container works inside the workspace.
- [ ] **Firecracker**: Capable of launching a minimal Firecracker microVM (hello-world).

## Technical Notes

- **Terraform**: The `docker_container` resource likely needs `privileged = true` or explicit device mapping for `/dev/kvm`.
- **Host Requirement**: The physical node running Coder must have VT-x/AMD-V enabled and Nested Virt active.
- **Security Check**: Verify if `privileged` mode is acceptable for our security posture, or if `cap_add=["SYS_ADMIN"]` + device mapping is sufficient.
