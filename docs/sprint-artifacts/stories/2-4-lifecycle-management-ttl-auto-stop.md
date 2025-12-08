# Story 2.4: Lifecycle Management (TTL & Auto-Stop)

**Status:** Ready for Review
**Estimate:** 3 Points
**Priority:** Medium
**Assignee:** AI Agent

## Story

As a Developer,
I want to explicitly start, stop, and restart my workspaces,
So that I can manage resources effectively and recover from issues.

## Acceptance Criteria

1. **Given** A running workspace
2. **When** I run `spt-flow stop [workspace-name]`
3. **Then** The workspace transitions to "stopped" state
4. **Given** A stopped workspace
5. **When** I run `spt-flow start [workspace-name]`
6. **Then** The workspace transitions to "running" state
7. **Given** A running workspace
8. **When** I run `spt-flow restart [workspace-name]`
9. **Then** The workspace stops and then starts again

## Tasks / Subtasks

- [x] Implement `stop` Command
  - [x] Create `cli/cmd/spt-flow/stop.go`
  - [x] Add `StopWorkspace` method to `pkg/coder/client.go`
- [x] Implement `start` Command
  - [x] Create `cli/cmd/spt-flow/start.go`
  - [x] Add `StartWorkspace` method to `pkg/coder/client.go` (or reuse existing logic)
- [x] Implement `restart` Command
  - [x] Create `cli/cmd/spt-flow/restart.go`
  - [x] Orchestrate stop then start

## Dev Notes

### Coder SDK

- Use `CreateWorkspaceBuild` with `Transition: WorkspaceTransitionStop` for stopping.
- Use `CreateWorkspaceBuild` with `Transition: WorkspaceTransitionStart` for starting.

### UX

- Commands should wait for the transition to complete (or at least provide a flag `--wait`, default true?).
- Provide feedback on current status.

## Dev Agent Record

### Context Reference

- `docs/epics.md` (Story 2.4)
- `docs/architecture.md`

### Agent Model Used

- Claude 3.5 Sonnet (via Antigravity)
