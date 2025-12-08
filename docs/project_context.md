---
project_name: 'turbo-flow-claude'
user_name: 'Philippe'
date: '2025-12-07'
sections_completed:
  ['technology_stack', 'language_rules', 'framework_rules', 'testing_rules', 'quality_rules', 'workflow_rules', 'critical_rules']
status: 'complete'
rule_count: 42
optimized_for_llm: true
existing_patterns_found: 5
---

# Project Context: turbo-flow-claude

> **AI AGENT INSTRUCTION:** Read this file FIRST before implementing any code. It contains critical rules, patterns, and conventions that must be followed to ensure consistency.

## 1. Technology Stack & Versions

| Category | Technology | Version | Critical Note |
|----------|------------|---------|---------------|
| **CLI Language** | Go | 1.21+ | Use `spf13/cobra` for commands, `spf13/viper` for config |
| **IaC** | Terraform | 1.5+ | Use Standard Module Layout |
| **Containerization** | Docker | Latest | "Fat Image" strategy for agents |
| **API** | Coder API | v2 | Source of Truth for state |

## 2. Critical Implementation Rules

> 🛑 **STOP & READ:** Violating these rules will cause system failure.

- **Stateless CLI:** NEVER persist workspace state locally. Always query the Coder API. Local state is for user prefs only.
- **Credential Safety:** NEVER write SSH keys or tokens to disk (except session token). Use SSH Agent Forwarding.
- **Observability:** ALWAYS generate and propagate `X-Request-ID` (Correlation ID) for all API calls.
- **No Local DB:** Do NOT use sqlite or local databases.

## 3. Naming & Organization Patterns

- **CLI Commands:** Kebab-case (e.g., `spt-flow up`)
- **Go Packages:** Short, lowercase, singular (e.g., `internal/ssh`)
- **Terraform Vars:** Snake_case with prefix (e.g., `agent_image_url`)
- **Directory Structure:**
  - `cli/`: Go code only
  - `infra/`: Terraform & Docker only
  - `internal/`: Private logic (no `pkg/`)

## 4. Testing Standards

- **Integration:** `test/e2e/` for CLI -> Coder API tests.
- **Unit:** Co-located `_test.go` files.

### Language-Specific Rules (Go)

- **Error Handling:** ALWAYS use `fmt.Errorf("context: %w", err)` to wrap errors. Do NOT use `pkg/errors`.
- **Context Usage:** Pass `context.Context` as the first argument to all long-running or I/O bound functions.
- **Interfaces:** Define interfaces where they are *used* (consumer-side), not where they are implemented.
- **Zero Values:** Leverage zero values (e.g., `var mu sync.Mutex`) instead of explicit initialization where possible.
- **Constructors:** Use `New...` functions (e.g., `NewClient`) only when initialization logic is required.
- **Naked Returns:** Avoid named returns (naked returns) in functions longer than 5 lines.
- **No Panic:** Do NOT use `panic` in `internal/` packages. Return errors. Only `cmd/` can panic.
- **Testing:** Use Table Driven Tests for all logic with multiple edge cases.

### Framework-Specific Rules (Cobra/Viper)

- **Command Execution:** Use `RunE` instead of `Run` to ensure errors are returned and handled by the root command.
- **Config Binding:** Bind flags to Viper keys in `init()`. Do NOT use flags directly in `RunE`.
- **Env Vars:** Enable `viper.AutomaticEnv()` with prefix `SPT_FLOW` (e.g., `SPT_FLOW_DEBUG`).
- **Strict Dependency Injection:** `viper` imports are **BANNED** in `internal/`. Unmarshal config into a struct in `cmd/` and pass it down.
- **Precedence:** Ensure Flags > Env Vars > Config File > Defaults.
- **Silence Usage:** Set `SilenceUsage: true` on the root command so usage isn't printed on runtime errors.
- **Secret Safety:** Never log the full config struct without sanitizing secrets (tokens, keys).

### Testing Rules

- **Unit Tests:** Co-located `_test.go` files. Use `testify/assert` for readable assertions (avoid `testify/suite`).
- **Integration Tests:** Place in `test/e2e/`. Use `//go:build integration` build tag.
- **Real vs Mock:** `test/e2e/` MUST run against a real Coder instance (or full simulator). Do NOT mock the API client in E2E.
- **Mocking:** Use `moq` to generate mocks from interfaces for *unit tests* only.
- **Golden Files:** Use golden files (`.golden`) for complex CLI output assertions to ensure exact text matching.
- **Test Main:** Use `TestMain` only for global setup/teardown (e.g., starting a mock server).

### Code Quality & Style Rules

- **Linting:** Zero tolerance. Enable `errcheck`, `staticcheck`, and `revive` in `golangci-lint`.
- **Formatting:** `gofmt` is the law.
- **Comments:** Explain **WHY** (intent/context), not **WHAT** (implementation).
- **Complexity:** Use **Guard Clauses** (Early Returns) to avoid deep nesting. Keep indentation minimal.
- **Imports:** Group imports: 1. Stdlib, 2. Third-party, 3. Internal.
- **Function Length:** Aim for < 50 lines. Break down large functions into helpers.

### Development Workflow Rules

- **Branch Naming:** Use prefixes: `feat/`, `fix/`, `chore/`, `docs/` (e.g., `feat/add-login`).
- **Commit Messages:** Follow **Conventional Commits** with **Scopes** (`feat(cli): add login`, `fix(infra): update image`).
- **Pull Requests:** Squash and Merge strategy to keep main history clean. Link to issues.
- **Releases:** Automated via `goreleaser`. Test locally first with `goreleaser --snapshot --skip-publish`.
- **Pre-Push:** Run `make lint` locally. Do not rely on CI for basic checks.

### Critical Don't-Miss Rules

- **Anti-Pattern:** Do NOT use `init()` for logic. Only use it for Cobra flag binding.
- **Global State:** Global variables for Clients, Config, or Loggers are **BANNED**. Pass dependencies explicitly.
- **Cross-Platform:** Do NOT use hardcoded paths. Use `os.TempDir()` and `filepath.Join`.
- **Windows SSH:** Support Windows Named Pipes (`\\.\pipe\openssh-ssh-agent`) for agent connection. Do not rely solely on `SSH_AUTH_SOCK` on Windows.
- **Shell Safety:** Do NOT use `os/exec` with raw strings or `sh -c`. Use `exec.Command("cmd", args...)`.
- **Timeouts:** All API calls MUST have a timeout (via Context). Never hang indefinitely.

---

## Usage Guidelines

**For AI Agents:**

- Read this file before implementing any code
- Follow ALL rules exactly as documented
- When in doubt, prefer the more restrictive option
- Update this file if new patterns emerge

**For Humans:**

- Keep this file lean and focused on agent needs
- Update when technology stack changes
- Review quarterly for outdated rules
- Remove rules that become obvious over time

Last Updated: 2025-12-07
