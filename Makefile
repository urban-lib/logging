.PHONY: all fmt lint test test-cover build tidy clean help

# Default target
all: fmt lint test build ## Run fmt, lint, test, build

# ── Formatting ──────────────────────────────────────────────

fmt: ## Format code with gofmt and goimports
	@echo "==> Formatting..."
	@gofmt -s -w .
	@command -v goimports >/dev/null 2>&1 && goimports -w . || true

# ── Linting ─────────────────────────────────────────────────

lint: ## Run golangci-lint
	@echo "==> Linting..."
	@golangci-lint run ./...

# ── Testing ─────────────────────────────────────────────────

test: ## Run all tests
	@echo "==> Running tests..."
	@go test ./... -count=1 -race

test-cover: ## Run tests with coverage report
	@echo "==> Running tests with coverage..."
	@go test ./... -count=1 -race -coverprofile=coverage.out -covermode=atomic
	@go tool cover -func=coverage.out
	@echo "==> HTML report: go tool cover -html=coverage.out"

# ── Building ────────────────────────────────────────────────

build: ## Build the project
	@echo "==> Building..."
	@go build ./...

# ── Dependencies ────────────────────────────────────────────

tidy: ## Run go mod tidy
	@echo "==> Tidying modules..."
	@go mod tidy

# ── Cleanup ─────────────────────────────────────────────────

clean: ## Remove build artifacts and coverage files
	@rm -f coverage.out
	@rm -rf logs/

# ── Help ────────────────────────────────────────────────────

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2}'
