// Package-level convenience functions that delegate to the global default logger.
// The default logger is automatically initialized from environment variables on first use.
package logging

import (
	"context"

	"go.uber.org/zap"
)

// Debugf logs a formatted debug message using the default logger.
func Debugf(format string, args ...any) { getDefault().Debugf(format, args...) }

// Infof logs a formatted info message using the default logger.
func Infof(format string, args ...any) { getDefault().Infof(format, args...) }

// Warnf logs a formatted warning message using the default logger.
func Warnf(format string, args ...any) { getDefault().Warnf(format, args...) }

// Errorf logs a formatted error message using the default logger.
func Errorf(format string, args ...any) { getDefault().Errorf(format, args...) }

// Fatalf logs a formatted fatal message and calls os.Exit(1) using the default logger.
func Fatalf(format string, args ...any) { getDefault().Fatalf(format, args...) }

// Panicf logs a formatted panic message and panics using the default logger.
func Panicf(format string, args ...any) { getDefault().Panicf(format, args...) }

// Debug logs args at debug level using the default logger.
func Debug(args ...any) { getDefault().Debug(args...) }

// Info logs args at info level using the default logger.
func Info(args ...any) { getDefault().Info(args...) }

// Warn logs args at warn level using the default logger.
func Warn(args ...any) { getDefault().Warn(args...) }

// Error logs args at error level using the default logger.
func Error(args ...any) { getDefault().Error(args...) }

// Panic logs args at panic level and panics using the default logger.
func Panic(args ...any) { getDefault().Panic(args...) }

// Fatal logs args at fatal level and calls os.Exit(1) using the default logger.
func Fatal(args ...any) { getDefault().Fatal(args...) }

// Debugw logs a message with key-value pairs at debug level using the default logger.
func Debugw(msg string, keysAndValues ...any) { getDefault().Debugw(msg, keysAndValues...) }

// Infow logs a message with key-value pairs at info level using the default logger.
func Infow(msg string, keysAndValues ...any) { getDefault().Infow(msg, keysAndValues...) }

// Warnw logs a message with key-value pairs at warn level using the default logger.
func Warnw(msg string, keysAndValues ...any) { getDefault().Warnw(msg, keysAndValues...) }

// Errorw logs a message with key-value pairs at error level using the default logger.
func Errorw(msg string, keysAndValues ...any) { getDefault().Errorw(msg, keysAndValues...) }

// Fatalw logs a message with key-value pairs at fatal level using the default logger.
func Fatalw(msg string, keysAndValues ...any) { getDefault().Fatalw(msg, keysAndValues...) }

// Panicw logs a message with key-value pairs at panic level using the default logger.
func Panicw(msg string, keysAndValues ...any) { getDefault().Panicw(msg, keysAndValues...) }

// WithFields returns a new Logger with the given fields attached, using the default logger.
// The returned Logger has CallerSkip adjusted for direct use (not through global wrapper).
func WithFields(fields Fields) Logger {
	dl := getDefault()
	if zl, ok := dl.(*zapLogger); ok {
		// Build fields for sugar
		f := make([]any, 0, len(fields)*2)
		for k, v := range fields {
			f = append(f, k, v)
		}
		// Create a new logger with CallerSkip decreased by 1 (for direct use)
		adjusted := zl.logger.WithOptions(zap.AddCallerSkip(-1))
		return &zapLogger{
			sugar:  adjusted.Sugar().With(f...),
			logger: adjusted,
		}
	}
	return dl.WithFields(fields)
}

// WithContext returns a new Logger enriched with fields from the context,
// using the default logger. CallerSkip is adjusted for direct use.
// See [ContextWithFields] for storing fields in a context.
func WithContext(ctx context.Context) Logger {
	dl := getDefault()
	if zl, ok := dl.(*zapLogger); ok {
		adjusted := zl.logger.WithOptions(zap.AddCallerSkip(-1))
		l := &zapLogger{
			sugar:  adjusted.Sugar(),
			logger: adjusted,
		}
		return l.WithContext(ctx)
	}
	return dl.WithContext(ctx)
}
