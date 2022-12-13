package logging

import "go.uber.org/zap"

// Debugf formatting string
func Debugf(format string, args ...interface{}) { logger.Debugf(format, args...) }

// Infof formatting string
func Infof(format string, args ...interface{}) { logger.Infof(format, args...) }

// Warnf formatting string
func Warnf(format string, args ...interface{}) { logger.Warnf(format, args...) }

// Errorf formatting string
func Errorf(format string, args ...interface{}) { logger.Errorf(format, args...) }

// Fatalf formatting string
func Fatalf(format string, args ...interface{}) { logger.Fatalf(format, args...) }

// Panicf formatting string
func Panicf(format string, args ...interface{}) { logger.Panicf(format, args...) }

// Debug ...
func Debug(args ...interface{}) { logger.Debug(args...) }

// Info ...
func Info(args ...interface{}) { logger.Info(args...) }

// Warn ...
func Warn(args ...interface{}) { logger.Warn(args...) }

// Error ...
func Error(args ...interface{}) { logger.Error(args...) }

// Panic ...
func Panic(args ...interface{}) { logger.Panic(args...) }

// Fatal ...
func Fatal(args ...interface{}) { logger.Fatal(args...) }

// WithFields formatting string
func WithFields(fields Fields) *zap.SugaredLogger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := logger.With(f...)
	return newLogger
}
