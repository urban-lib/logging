# logging

[![CI](https://github.com/urban-lib/logging/actions/workflows/ci.yml/badge.svg)](https://github.com/urban-lib/logging/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/urban-lib/logging/v3.svg)](https://pkg.go.dev/github.com/urban-lib/logging/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/urban-lib/logging/v3)](https://goreportcard.com/report/github.com/urban-lib/logging/v3)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Structured logging library for Go built on top of [Uber Zap](https://github.com/uber-go/zap) with automatic log file rotation via [lumberjack](https://github.com/natefinch/lumberjack.v2).

## Features

- **Interface-based** — `Logger` interface for easy mocking and testing
- **Multiple configuration styles** — functional options, struct config, or environment variables
- **Dual output** — console (human-readable) + file (JSON) with independent log levels
- **Automatic rotation** — file size limits, max backups, max age, gzip compression
- **Structured fields** — `WithFields()` returns a new `Logger` for request-scoped metadata
- **Context-aware** — `WithContext()` propagates trace/request IDs via `context.Context`
- **Sugar key-value API** — `Infow`, `Errorw`, etc. for ergonomic structured logging
- **Typed field helpers** — `String()`, `Int()`, `Err()`, `Duration()` for zero-allocation fields
- **Log sampling** — configurable rate limiting for high-throughput services
- **Global convenience layer** — package-level functions (`logging.Infof(...)`) for quick usage
- **CallerSkip-aware** — correct source location in both direct and global usage

## Installation

```bash
go get github.com/urban-lib/logging/v3
```

Requires **Go 1.22+**.

## Quick Start

### 1. Global logger (auto-initialized from env)

```go
package main

import "github.com/urban-lib/logging/v3"

func main() {
    // Uses environment variables — zero config needed
    logging.Infof("Server started on port %d", 8080)

    logging.WithFields(logging.Fields{
        "user_id":    42,
        "request_id": "abc-123",
    }).Infof("Request processed")
}
```

### 2. Instance-based with functional options

```go
logger, err := logging.New(
    logging.WithConsoleLevel("info"),
    logging.WithFileEnabled(true),
    logging.WithFilePath("/var/log/myapp/app.log"),
    logging.WithFileLevel("error"),
    logging.WithFileMaxSize(50),
)
if err != nil {
    panic(err)
}
defer logger.Sync()

logger.Infof("Application started")
logger.WithFields(logging.Fields{"component": "db"}).Errorf("Query failed")
```

### 3. Instance-based with config struct

```go
logger, err := logging.NewWithConfig(logging.Config{
    ConsoleLevel:   "debug",
    FileEnabled:    true,
    FileLevel:      "info",
    FilePath:       "logs/app.log",
    FileMaxSize:    100,
    FileMaxBackups: 7,
    FileMaxAge:     5,
    FileCompress:   true,
    CallerSkip:     1,
})
if err != nil {
    panic(err)
}
defer logger.Sync()
```

### 4. Instance-based from environment

```go
logger, err := logging.NewFromEnv()
if err != nil {
    panic(err)
}
defer logger.Sync()
```

### 5. Replace global default

```go
logger, _ := logging.New(logging.WithConsoleLevel("warn"))
logging.SetDefault(logger)

// Now package-level functions use your custom logger
logging.Warnf("This goes through the custom logger")
```

### 6. Context-aware logging

```go
ctx := logging.ContextWithFields(ctx, logging.Fields{
    "trace_id":   "abc-xyz-123",
    "request_id": "req-456",
})

// Instance
logger.WithContext(ctx).Infof("Processing order")

// Global
logging.WithContext(ctx).Warnf("Slow downstream call")
```

### 7. Sugar key-value API

```go
logger.Infow("request handled",
    "method", "POST",
    "path", "/api/users",
    "latency_ms", 42,
)
```

### 8. Typed field helpers

```go
logger.Log().Info("structured fields",
    logging.String("service", "payment"),
    logging.Int("attempt", 3),
    logging.Duration("latency", 150*time.Millisecond),
    logging.Err(fmt.Errorf("timeout")),
)
```

### 9. Sampling (high-throughput)

```go
logger, _ := logging.New(
    logging.WithConsoleLevel("info"),
    logging.WithSampling(100, 10), // 100/s initial, then every 10th
)
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `LOG_LEVEL_CONSOLE` | Console log level (`debug`, `info`, `warn`, `error`) | `debug` |
| `LOG_FILE_ENABLE` | Enable file logging (`true`/`false`) | `false` |
| `LOG_LEVEL_FILE` | File log level | `debug` |
| `LOG_FILE_PATH` | Path to log file | `logs/example.log` |
| `LOG_FILE_MAX_SIZE` | Max file size in MB before rotation | `100` |
| `LOG_FILE_MAX_BACKUPS` | Max number of rotated files to keep | `7` |
| `LOG_FILE_MAX_AGE` | Max days to retain old files | `5` |

## Log Levels

| Level | Description |
|-------|-------------|
| `debug` | Verbose diagnostic information |
| `info` | General operational messages |
| `warn` | Non-critical warnings |
| `error` | Errors requiring attention |
| `fatal` | Critical errors → `os.Exit(1)` |
| `panic` | Critical errors → `panic()` |

## Output Formats

**Console** — human-readable (development encoder):
```
2026-03-09T12:00:00.000+0200  INFO  myapp/main.go:15  Server started on port 8080
```

**File** — JSON (production encoder):
```json
{"level":"info","ts":"2026-03-09T12:00:00.000+0200","caller":"myapp/main.go:15","msg":"Server started on port 8080"}
```

## API Overview

### Constructors

| Function | Description |
|----------|-------------|
| `New(opts ...Option) (Logger, error)` | Create logger with functional options |
| `NewWithConfig(cfg Config) (Logger, error)` | Create logger from config struct |
| `NewFromEnv() (Logger, error)` | Create logger from environment variables |

### Global Functions

| Function | Description |
|----------|-------------|
| `SetDefault(l Logger)` | Replace the global default logger |
| `WithFields(fields Fields) Logger` | Add structured fields (returns new Logger) |
| `WithContext(ctx) Logger` | Add fields from context (returns new Logger) |
| `Debugf/Infof/Warnf/Errorf/Fatalf/Panicf` | Formatted log at given level |
| `Debug/Info/Warn/Error/Fatal/Panic` | Log args at given level |
| `Debugw/Infow/Warnw/Errorw/Fatalw/Panicw` | Log with key-value pairs |
| `ContextWithFields(ctx, Fields) context.Context` | Store fields in context |
| `FieldsFromContext(ctx) Fields` | Extract fields from context |
| `CheckEnvironments()` | Print current env var values |

### Logger Interface

```go
type Logger interface {
    Debugf(format string, args ...any)
    Infof(format string, args ...any)
    Warnf(format string, args ...any)
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
    Panicf(format string, args ...any)

    Debug(args ...any)
    Info(args ...any)
    Warn(args ...any)
    Error(args ...any)
    Panic(args ...any)
    Fatal(args ...any)

    Debugw(msg string, keysAndValues ...any)
    Infow(msg string, keysAndValues ...any)
    Warnw(msg string, keysAndValues ...any)
    Errorw(msg string, keysAndValues ...any)
    Fatalw(msg string, keysAndValues ...any)
    Panicw(msg string, keysAndValues ...any)

    WithFields(fields Fields) Logger
    WithContext(ctx context.Context) Logger
    Log() *zap.Logger
    Sync() error
}
```

### Functional Options

| Option | Description |
|--------|-------------|
| `WithConsoleLevel(level)` | Set console log level |
| `WithFileEnabled(bool)` | Enable/disable file logging |
| `WithFileLevel(level)` | Set file log level |
| `WithFilePath(path)` | Set log file path |
| `WithFileMaxSize(mb)` | Max file size before rotation |
| `WithFileMaxBackups(n)` | Max rotated files to keep |
| `WithFileMaxAge(days)` | Max days to retain files |
| `WithFileCompress(bool)` | Enable gzip compression |
| `WithCallerSkip(n)` | Adjust caller skip depth |
| `WithSampling(initial, thereafter)` | Enable log sampling |

## Development

```bash
make fmt          # Format code
make lint         # Run golangci-lint
make test         # Run tests with race detector
make test-cover   # Tests + coverage report
make build        # Build all packages
make tidy         # go mod tidy
make all          # fmt + lint + test + build
```

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT — see [LICENSE](LICENSE).