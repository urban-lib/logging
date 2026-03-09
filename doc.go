// Package logging provides a structured logging library built on top of
// [Uber Zap] with automatic log file rotation via [lumberjack].
//
// The package exposes a [Logger] interface that can be created in three ways:
//
//   - [New] — functional options on top of [DefaultConfig]
//   - [NewWithConfig] — explicit [Config] struct
//   - [NewFromEnv] — configured from environment variables
//
// For convenience, package-level functions ([Infof], [Errorf], [WithFields], etc.)
// delegate to a lazily-initialized global logger. The global logger reads
// configuration from environment variables on first use and can be replaced
// via [SetDefault].
//
// # Environment Variables
//
// The following environment variables configure the global logger:
//
//   - LOG_LEVEL_CONSOLE — console log level (debug, info, warn, error). Default: debug.
//   - LOG_FILE_ENABLE   — enable file logging (true/false). Default: false.
//   - LOG_LEVEL_FILE    — file log level. Default: debug.
//   - LOG_FILE_PATH     — path to the log file. Default: logs/example.log.
//   - LOG_FILE_MAX_SIZE — max file size in MB before rotation. Default: 100.
//   - LOG_FILE_MAX_BACKUPS — max number of rotated files. Default: 7.
//   - LOG_FILE_MAX_AGE  — max days to keep old files. Default: 5.
//
// # Output Formats
//
// Console output uses a human-readable development encoder:
//
//	2026-03-09T12:00:00.000+0200  INFO  myapp/main.go:15  Server started
//
// File output uses a JSON production encoder:
//
//	{"level":"info","ts":"2026-03-09T12:00:00.000+0200","caller":"myapp/main.go:15","msg":"Server started"}
//
// # Example (instance-based)
//
//	logger, err := logging.New(
//	    logging.WithConsoleLevel("info"),
//	    logging.WithFileEnabled(true),
//	    logging.WithFilePath("logs/app.log"),
//	)
//	if err != nil {
//	    panic(err)
//	}
//	defer logger.Sync()
//
//	logger.Infof("Hello, %s!", "world")
//	logger.WithFields(logging.Fields{"request_id": "abc"}).Warnf("Slow query")
//
// # Example (global)
//
//	logging.Infof("Using global logger")
//	logging.WithFields(logging.Fields{"user": "alice"}).Infof("Logged in")
//
// # Context-aware logging
//
// Use [ContextWithFields] to store fields in a context.Context, then
// retrieve them automatically via Logger.WithContext or the global [WithContext]:
//
//	ctx = logging.ContextWithFields(ctx, logging.Fields{"trace_id": "xyz"})
//	logger.WithContext(ctx).Infof("traced request")
//
// # Sugar key-value API
//
// The Logger interface exposes Debugw/Infow/Warnw/Errorw/Fatalw/Panicw methods
// that accept a message followed by alternating key-value pairs:
//
//	logger.Infow("request", "method", "GET", "path", "/api")
//
// # Typed field helpers
//
// For zero-allocation structured logging via [Logger.Log]:
//
//	logger.Log().Info("event", logging.String("svc", "pay"), logging.Err(err))
//
// # Log sampling
//
// Use [WithSampling] to enable rate limiting for high-throughput scenarios:
//
//	logger, _ := logging.New(logging.WithSampling(100, 10))
//
// [Uber Zap]: https://github.com/uber-go/zap
// [lumberjack]: https://github.com/natefinch/lumberjack
package logging
