# Contributing to logging/v3

Thank you for your interest in the project! Below you'll find the guidelines and contribution process.

## Requirements

- **Go 1.22+**
- `golangci-lint` for linting
- `git` with [Conventional Commits](https://www.conventionalcommits.org/) support

## Getting Started

```bash
# Clone
git clone https://github.com/urban-lib/logging.git
cd logging

# Install dependencies
go mod download

# Run tests
go test ./...

# Lint
golangci-lint run ./...
```

## Git Workflow

### Branches

| Branch | Purpose |
|--------|---------|
| `main` | Stable version, protected from direct pushes |
| `develop` | Main development branch |
| `feature/*` | New features |
| `fix/*` | Bug fixes |
| `docs/*` | Documentation changes |

### Process

1. Create a branch from `develop`:
   ```bash
   git checkout develop
   git pull origin develop
   git checkout -b feature/my-feature
   ```
2. Make your changes and write tests.
3. Make sure all tests pass: `go test ./...`
4. Make sure the linter reports no issues: `golangci-lint run ./...`
5. Open a Pull Request targeting `develop`.

## Conventional Commits

All commits **must** follow the [Conventional Commits](https://www.conventionalcommits.org/) format. This is required for automatic tag generation and changelog.

### Format

```
<type>(<scope>): <description>

[optional body]

[optional footer(s)]
```

### Types

| Type | Description | Version Impact |
|------|-------------|----------------|
| `feat` | New feature | **minor** (0.X.0) |
| `fix` | Bug fix | **patch** (0.0.X) |
| `docs` | Documentation changes | — |
| `refactor` | Refactoring without behavior change | — |
| `test` | Adding / modifying tests | — |
| `chore` | Dependency updates, CI, etc. | — |
| `perf` | Performance improvements | **patch** |
| `BREAKING CHANGE` | Breaking change (in footer or `!` after type) | **major** (X.0.0) |

### Examples

```bash
# New feature
git commit -m "feat(logger): add WithContext method for trace propagation"

# Bug fix
git commit -m "fix(env): correct default level comparison in getLogLevelConsole"

# Breaking change
git commit -m "feat(api)!: replace GetLogger with New constructor"

# Documentation
git commit -m "docs: update README with env variables table"
```

## Automatic Tagging

When merging into `main`, GitHub Actions will automatically:
1. Analyze commits since the last tag.
2. Determine the version bump type (major/minor/patch) based on Conventional Commits.
3. Create a new git tag (`vX.Y.Z`).
4. Generate a GitHub Release with a changelog.

**Do not create tags manually** — CI handles this automatically.

## Code Style

- Follow `gofmt` / `goimports` formatting.
- Public functions, types, and methods **must** have GoDoc comments.
- Variable and function names should be in English.
- Avoid global state where possible.
- Handle errors explicitly, never ignore them.

## Tests

- Every new feature or fix **must** be accompanied by tests.
- Use table-driven tests where appropriate.
- Minimum coverage: **80%**.
- Test files: `*_test.go` alongside the code being tested.

```bash
# Run tests with coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Pull Requests

### Checklist before opening a PR

- [ ] Code compiles without errors (`go build ./...`)
- [ ] All tests pass (`go test ./...`)
- [ ] Linter reports no issues (`golangci-lint run`)
- [ ] Tests added for new functionality / fix
- [ ] Commits follow Conventional Commits
- [ ] Documentation updated (if public API changed)

### Review

- A PR requires at least **1 approval** to merge.
- CI checks must be green.
- Squash merge into `develop`, merge commit into `main`.

## Reporting Bugs

Create an Issue with:
- Go version and OS
- Minimal reproduction code
- Expected vs. actual behavior
- Logs / stack trace (if available)

## License

By contributing, you agree that your contributions will be published under the same license as the project.
