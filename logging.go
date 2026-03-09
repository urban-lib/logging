package logging

import (
	"go.uber.org/zap"
)

// Logger is the interface that wraps structured logging methods.
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

	// WithFields returns a new Logger with the given fields attached.
	WithFields(fields Fields) Logger

	// Log returns the underlying *zap.Logger for advanced usage.
	Log() *zap.Logger

	// Sync flushes any buffered log entries. Should be called before application exit.
	Sync() error
}

// Fields is a map of key-value pairs for structured logging.
type Fields map[string]any

// zapLogger implements the Logger interface using Uber Zap.
type zapLogger struct {
	sugar  *zap.SugaredLogger
	logger *zap.Logger
}

// Debugf logs a formatted debug message.
func (z *zapLogger) Debugf(format string, args ...any) { z.sugar.Debugf(format, args...) }

// Infof logs a formatted info message.
func (z *zapLogger) Infof(format string, args ...any) { z.sugar.Infof(format, args...) }

// Warnf logs a formatted warning message.
func (z *zapLogger) Warnf(format string, args ...any) { z.sugar.Warnf(format, args...) }

// Errorf logs a formatted error message.
func (z *zapLogger) Errorf(format string, args ...any) { z.sugar.Errorf(format, args...) }

// Fatalf logs a formatted fatal message and calls os.Exit(1).
func (z *zapLogger) Fatalf(format string, args ...any) { z.sugar.Fatalf(format, args...) }

// Panicf logs a formatted panic message and panics.
func (z *zapLogger) Panicf(format string, args ...any) { z.sugar.Panicf(format, args...) }

// Debug logs args at debug level.
func (z *zapLogger) Debug(args ...any) { z.sugar.Debug(args...) }

// Info logs args at info level.
func (z *zapLogger) Info(args ...any) { z.sugar.Info(args...) }

// Warn logs args at warn level.
func (z *zapLogger) Warn(args ...any) { z.sugar.Warn(args...) }

// Error logs args at error level.
func (z *zapLogger) Error(args ...any) { z.sugar.Error(args...) }

// Panic logs args at panic level and panics.
func (z *zapLogger) Panic(args ...any) { z.sugar.Panic(args...) }

// Fatal logs args at fatal level and calls os.Exit(1).
func (z *zapLogger) Fatal(args ...any) { z.sugar.Fatal(args...) }

// WithFields returns a new Logger with the given fields attached.
func (z *zapLogger) WithFields(fields Fields) Logger {
	f := make([]any, 0, len(fields)*2)
	for k, v := range fields {
		f = append(f, k, v)
	}
	return &zapLogger{
		sugar:  z.sugar.With(f...),
		logger: z.logger,
	}
}

// Log returns the underlying *zap.Logger.
func (z *zapLogger) Log() *zap.Logger { return z.logger }

// Sync flushes any buffered log entries.
func (z *zapLogger) Sync() error { return z.logger.Sync() }
