# logging/v3 — Documentation

A Go structured logging library built on [Uber Zap](https://github.com/uber-go/zap) with log file rotation via [lumberjack](https://github.com/natefinch/lumberjack.v2).

## Installation

```bash
go get github.com/urban-lib/logging/v3
```

## Constructors

### `New(opts ...Option) (Logger, error)`

Creates a logger with functional options on top of `DefaultConfig()`:

```go
logger, err := logging.New(
    logging.WithConsoleLevel("info"),
    logging.WithFileEnabled(true),
    logging.WithFilePath("logs/app.log"),
    logging.WithFileLevel("error"),
)
```

### `NewWithConfig(cfg Config) (Logger, error)`

Creates a logger directly from a `Config` struct:

```go
logger, err := logging.NewWithConfig(logging.Config{
    ConsoleLevel: "debug",
    FileEnabled:  true,
    FileLevel:    "info",
    FilePath:     "logs/app.log",
    FileMaxSize:  100,
    CallerSkip:   1,
})
```

### `NewFromEnv() (Logger, error)`

Creates a logger from environment variables (overrides `DefaultConfig()`):

```go
logger, err := logging.NewFromEnv()
```

## Logger Interface

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

- `WithFields` — returns a **new** `Logger` with added fields (the parent is unchanged).
- `WithContext` — returns a **new** `Logger` with fields from the context (see `ContextWithFields`).
- `Debugw/Infow/...` — sugar key-value API: `logger.Infow("msg", "key", "val", "num", 42)`.
- `Log()` — returns `*zap.Logger` for advanced usage.
- `Sync()` — flushes buffers. Always call `defer logger.Sync()` before exiting the program.

## Config

```go
type Config struct {
    ConsoleLevel       string  // "debug", "info", "warn", "error" — default: "debug"
    FileEnabled        bool    // default: false
    FileLevel          string  // default: "debug"
    FilePath           string  // default: "logs/example.log"
    FileMaxSize        int     // MB — default: 100
    FileMaxBackups     int     // default: 7
    FileMaxAge         int     // days — default: 5
    FileCompress       bool    // default: true
    CallerSkip         int     // default: 1
    SamplingInitial    int     // messages/sec before sampling — default: 0 (disabled)
    SamplingThereafter int     // keep every Nth after initial — default: 0
}
```

`DefaultConfig()` returns a configuration with all default values.

## Functional Options

| Option | Description |
|--------|-------------|
| `WithConsoleLevel(level)` | Console log level |
| `WithFileEnabled(bool)` | Enable/disable file logging |
| `WithFileLevel(level)` | File log level |
| `WithFilePath(path)` | Path to the log file |
| `WithFileMaxSize(mb)` | Maximum file size |
| `WithFileMaxBackups(n)` | Number of rotated files |
| `WithFileMaxAge(days)` | File retention period |
| `WithFileCompress(bool)` | Gzip compression of rotated files |
| `WithCallerSkip(n)` | Caller skip depth |
| `WithSampling(initial, thereafter)` | Log sampling (rate limiting) |

## Context-Aware Logging

For propagating trace/request IDs through `context.Context`:

```go
// Store fields in context
ctx = logging.ContextWithFields(ctx, logging.Fields{
    "trace_id":   "abc-xyz",
    "request_id": "req-123",
})

// Multiple calls accumulate fields:
ctx = logging.ContextWithFields(ctx, logging.Fields{"user_id": 42})

// Use with a logger
logger.WithContext(ctx).Infof("Processing request")

// Or globally
logging.WithContext(ctx).Infof("Global with context")

// Retrieve fields from context
fields := logging.FieldsFromContext(ctx)
```

## Sugar Key-Value API

The `Debugw/Infow/Warnw/Errorw/Fatalw/Panicw` methods accept a message and key-value pairs:

```go
logger.Infow("request handled",
    "method", "POST",
    "path", "/api/users",
    "latency_ms", 42,
)
```

Also available at the package level: `logging.Infow(...)`, `logging.Errorw(...)`, etc.

## Typed Fields (Field Helpers)

For zero-allocation fields when working with `*zap.Logger` via `Log()`:

```go
logger.Log().Info("structured",
    logging.String("service", "payment"),
    logging.Int("attempt", 3),
    logging.Duration("latency", 150*time.Millisecond),
    logging.Bool("cached", false),
    logging.Err(err),
    logging.NamedErr("cause", rootErr),
    logging.Float64("score", 0.95),
    logging.Time("started_at", startTime),
    logging.Any("payload", myStruct),
    logging.Stringer("addr", netAddr),
)
```

| Helper | Description |
|--------|-------------|
| `String(key, val)` | String |
| `Int(key, val)` | Integer |
| `Int64(key, val)` | Integer (int64) |
| `Float64(key, val)` | Float |
| `Bool(key, val)` | Boolean |
| `Err(err)` | Error (key = `"error"`) |
| `NamedErr(key, err)` | Error with a custom key |
| `Duration(key, val)` | Duration |
| `Time(key, val)` | Time |
| `Any(key, val)` | Arbitrary value (reflection) |
| `Stringer(key, val)` | Value with a `String()` method |

## Log Sampling

For high-throughput services, you can limit the number of identical log lines:

```go
logger, _ := logging.New(
    logging.WithSampling(100, 10), // 100 msg/sec initially, then every 10th
)
```

When `SamplingInitial = 0` — sampling is disabled (default).

## Global Functions

The package provides a convenient global function layer that delegates to the default logger:

```go
logging.Infof("Server started on port %d", 8080)
logging.WithFields(logging.Fields{"user": "alice"}).Warnf("Slow request")
logging.Infow("event", "key", "value")
logging.WithContext(ctx).Infof("Traced request")
```

The global logger is lazily initialized from environment variables on first use.

### `SetDefault(l Logger)`

Replaces the global logger with a custom instance. CallerSkip is automatically adjusted (+1) to correctly display the call source.

```go
logger, _ := logging.New(logging.WithConsoleLevel("warn"))
logging.SetDefault(logger)
logging.Warnf("Now using custom logger")
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `LOG_LEVEL_CONSOLE` | Console log level | `debug` |
| `LOG_FILE_ENABLE` | Enable file logging | `false` |
| `LOG_LEVEL_FILE` | File log level | `debug` |
| `LOG_FILE_PATH` | Log file path | `logs/example.log` |
| `LOG_FILE_MAX_SIZE` | Maximum file size (MB) | `100` |
| `LOG_FILE_MAX_BACKUPS` | Number of rotated files | `7` |
| `LOG_FILE_MAX_AGE` | Maximum file age (days) | `5` |

## Output Format

**Console** — human-readable (development encoder):
```
2026-03-09T12:00:00.000+0200  INFO  myapp/main.go:15  Request processed  {"user_id": 42}
```

**File** — JSON (production encoder):
```json
{"level":"info","ts":"2026-03-09T12:00:00.000+0200","caller":"myapp/main.go:15","msg":"Request processed","user_id":42}
```

## File Rotation

When `LOG_FILE_ENABLE=true`, log files are rotated automatically:
- When the file reaches `LOG_FILE_MAX_SIZE` MB
- Old files are compressed (gzip)
- No more than `LOG_FILE_MAX_BACKUPS` archives are kept
- Files older than `LOG_FILE_MAX_AGE` days are deleted

## Log Levels

| Level | Purpose |
|-------|---------|
| `debug` | Detailed diagnostic information |
| `info` | General operational information |
| `warn` | Warnings that do not block operation |
| `error` | Errors that require attention |
| `fatal` | Critical errors → `os.Exit(1)` |
| `panic` | Critical errors → `panic()` |

## CallerSkip

- **Instance-based** (`New`, `NewWithConfig`) — `CallerSkip: 1` (default), the caller points to the code that called `logger.Infof(...)`.
- **Global functions** (`logging.Infof(...)`) — +1 automatically (= 2), to skip the wrapper in `formatting.go`.
- **`SetDefault`** — automatically adds +1 to CallerSkip.
- **`WithFields` (global)** — adjusts CallerSkip(-1) so the returned `Logger` correctly displays the caller when used directly.

## Diagnostics

```go
logging.CheckEnvironments()
```

Prints the current values of all environment variables via `log.Println`.

## Contributing

See [CONTRIBUTING.md](../CONTRIBUTING.md).
