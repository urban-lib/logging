package logging

import (
	"go.uber.org/zap"
	"sync"
)

var (
	logger *internalLogger
	once   sync.Once
)

type (
	Logger interface {
		Debugf(format string, args ...interface{})
		Infof(format string, args ...interface{})
		Warnf(format string, args ...interface{})
		Errorf(format string, args ...interface{})
		Fatalf(format string, args ...interface{})
		Panicf(format string, args ...interface{})
		WithFields(fields Fields) Logger
		Log() *zap.Logger
		Debug(args ...interface{})
		Info(args ...interface{})
		Warn(args ...interface{})
		Error(args ...interface{})
		Panic(args ...interface{})
		Fatal(args ...interface{})
	}
	Fields         map[string]interface{}
	internalLogger struct {
		log   *zap.Logger
		sugar *zap.SugaredLogger
	}
)

func (l *internalLogger) Debugf(format string, args ...interface{}) { l.sugar.Debugf(format, args...) }

func (l *internalLogger) Infof(format string, args ...interface{}) { l.sugar.Infof(format, args...) }

func (l *internalLogger) Warnf(format string, args ...interface{}) { l.sugar.Warnf(format, args...) }

func (l *internalLogger) Errorf(format string, args ...interface{}) { l.sugar.Errorf(format, args...) }

func (l *internalLogger) Fatalf(format string, args ...interface{}) { l.sugar.Fatalf(format, args...) }

func (l *internalLogger) Panicf(format string, args ...interface{}) { l.sugar.Panicf(format, args...) }

func (l *internalLogger) Log() *zap.Logger { return l.log }

func (l *internalLogger) Debug(args ...interface{}) { l.sugar.Debug(args...) }

func (l *internalLogger) Info(args ...interface{}) { l.sugar.Info(args...) }

func (l *internalLogger) Warn(args ...interface{}) { l.sugar.Warn(args...) }

func (l *internalLogger) Error(args ...interface{}) { l.sugar.Error(args...) }

func (l *internalLogger) Panic(args ...interface{}) { l.sugar.Panic(args...) }

func (l *internalLogger) Fatal(args ...interface{}) { l.sugar.Fatal(args...) }

func (l *internalLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugar.With(f...)
	return &internalLogger{
		log:   l.log,
		sugar: newLogger,
	}
}
