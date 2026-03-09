# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- `WithContext(ctx context.Context) Logger` — context-aware logging for trace/request ID propagation.
- `ContextWithFields(ctx, Fields) context.Context` — store fields in context.
- `FieldsFromContext(ctx) Fields` — extract fields from context.
- Sugar key-value methods: `Debugw`, `Infow`, `Warnw`, `Errorw`, `Fatalw`, `Panicw` on `Logger` interface and global functions.
- Typed field helpers: `String()`, `Int()`, `Int64()`, `Float64()`, `Bool()`, `Err()`, `NamedErr()`, `Duration()`, `Time()`, `Any()`, `Stringer()`.
- Log sampling via `WithSampling(initial, thereafter)` option and `Config.SamplingInitial`/`SamplingThereafter` fields.
- New files: `context.go`, `fields.go`.

## [3.0.0] — 2026-03-09

### Breaking Changes
- Module path changed from `v2` to `github.com/urban-lib/logging/v3`.
- `GetLogger()` now returns `(Logger, error)` instead of `(*zap.SugaredLogger, *zap.Logger, error)`.
- `WithFields()` returns `Logger` interface instead of `*zap.SugaredLogger`.
- Removed the old `Configuration` struct; replaced with `Config`.
- Replaced `interface{}` with `any` throughout the public API.

### Added
- `Logger` interface — allows mocking and custom implementations.
- `New(opts ...Option) (Logger, error)` — functional-options constructor.
- `NewWithConfig(cfg Config) (Logger, error)` — struct-based constructor.
- `NewFromEnv() (Logger, error)` — env-based constructor.
- `Config` struct with documented fields and `DefaultConfig()`.
- 9 functional options: `WithConsoleLevel`, `WithFileEnabled`, `WithFileLevel`, `WithFilePath`, `WithFileMaxSize`, `WithFileMaxBackups`, `WithFileMaxAge`, `WithFileCompress`, `WithCallerSkip`.
- `SetDefault(l Logger)` — replace the global default logger.
- `Logger.Sync() error` — flush buffered log entries.
- `Logger.Log() *zap.Logger` — access the underlying zap logger.
- Non-formatted log methods: `Debug`, `Info`, `Warn`, `Error`, `Panic`, `Fatal`.
- Package-level `doc.go` with comprehensive GoDoc.
- Unit tests for env parsing, constructors, global functions, WithFields, integration (file logging).
- `.golangci.yml` linter configuration.
- `Makefile` with targets: `fmt`, `lint`, `test`, `test-cover`, `build`, `tidy`, `clean`.
- GitHub Actions CI workflow (lint → test → build on Go 1.22/1.23).
- GitHub Actions Release workflow (auto-tag + GitHub Release from Conventional Commits).
- `CONTRIBUTING.md`, `docs/README.md`, `CHANGELOG.md`, `LICENSE`.

### Fixed
- **Nil pointer panic**: global functions now lazy-init the logger on first use.
- **`getLogLevelConsole()` bug**: `len(def) > 1` → `def != ""` (single-char levels like `"warn"` now work).
- **`CheckEnvironments()` formatting**: added missing `\t` for `LOG_FILE_MAX_BACKUPS`.
- **Error handling in `sync.Once`**: init errors are now stored and returned from `GetLogger()`.
- **Import path typo**: `natefinsh` → `natefinch` in lumberjack import.

### Changed
- Upgraded to Go 1.22 (from 1.18).
- Upgraded `go.uber.org/zap` to v1.27.0.
- Upgraded `gopkg.in/natefinch/lumberjack.v2` to v2.2.1.
- Removed deprecated `go.uber.org/atomic` dependency.
- Caller encoder changed from `FullCallerEncoder` to `ShortCallerEncoder`.
- `GetLogger()` marked as deprecated in favor of instance-based constructors.
- Global functions use correct CallerSkip for accurate source locations.

## [2.0.0] — Pre-refactor

Legacy version with global state, `Configuration` struct, and direct `*zap.SugaredLogger` exposure. Superseded by v3.
