package logging

import "go.uber.org/zap"

// Debugf formatting string
func Debugf(format string, args ...interface{}) { sugaredLogger.Debugf(format, args...) }

// Infof formatting string
func Infof(format string, args ...interface{}) { sugaredLogger.Infof(format, args...) }

// Warnf formatting string
func Warnf(format string, args ...interface{}) { sugaredLogger.Warnf(format, args...) }

// Errorf formatting string
func Errorf(format string, args ...interface{}) { sugaredLogger.Errorf(format, args...) }

// Fatalf formatting string
func Fatalf(format string, args ...interface{}) { sugaredLogger.Fatalf(format, args...) }

// Panicf formatting string
func Panicf(format string, args ...interface{}) { sugaredLogger.Panicf(format, args...) }

// Debug ...
func Debug(args ...interface{}) { sugaredLogger.Debug(args...) }

// Info ...
func Info(args ...interface{}) { sugaredLogger.Info(args...) }

// Warn ...
func Warn(args ...interface{}) { sugaredLogger.Warn(args...) }

// Error ...
func Error(args ...interface{}) { sugaredLogger.Error(args...) }

// Panic ...
func Panic(args ...interface{}) { sugaredLogger.Panic(args...) }

// Fatal ...
func Fatal(args ...interface{}) { sugaredLogger.Fatal(args...) }

// WithFields formatting string
func WithFields(fields Fields) *zap.SugaredLogger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := sugaredLogger.With(f...)
	return newLogger
}
