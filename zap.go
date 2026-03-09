package logging

import (
	"fmt"
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	defaultLogger Logger
	once          sync.Once
	initErr       error
)

func parseLevel(level string, fallback zapcore.Level) zapcore.Level {
	if l, err := zapcore.ParseLevel(level); err != nil {
		return fallback
	} else {
		return l
	}
}

func textEncoder() zapcore.Encoder {
	devConfig := zap.NewDevelopmentEncoderConfig()
	devConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(devConfig)
}

func jsonEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(cfg)
}

// NewWithConfig creates a new Logger instance from the given Config.
func NewWithConfig(cfg Config) (Logger, error) {
	cores := make([]zapcore.Core, 0, 2)

	consoleLevel := parseLevel(cfg.ConsoleLevel, zapcore.InfoLevel)
	core := zapcore.NewCore(textEncoder(), zapcore.Lock(os.Stdout), consoleLevel)
	cores = append(cores, core)

	if cfg.FileEnabled {
		if cfg.FilePath == "" {
			return nil, fmt.Errorf("logging: file logging enabled but FilePath is empty")
		}
		fileLevel := parseLevel(cfg.FileLevel, zapcore.InfoLevel)
		writer := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.FilePath,
			MaxSize:    cfg.FileMaxSize,
			MaxAge:     cfg.FileMaxAge,
			MaxBackups: cfg.FileMaxBackups,
			Compress:   cfg.FileCompress,
		})
		cores = append(cores, zapcore.NewCore(jsonEncoder(), writer, fileLevel))
	}

	combinedCore := zapcore.NewTee(cores...)
	logger := zap.New(
		combinedCore,
		zap.AddCallerSkip(cfg.CallerSkip),
		zap.AddCaller(),
	)

	return &zapLogger{
		sugar:  logger.Sugar(),
		logger: logger,
	}, nil
}

// New creates a new Logger instance with functional options applied on top of DefaultConfig.
func New(opts ...Option) (Logger, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}
	return NewWithConfig(cfg)
}

// NewFromEnv creates a new Logger instance configured from environment variables.
// Environment variables override DefaultConfig values.
func NewFromEnv() (Logger, error) {
	cfg := DefaultConfig()

	if v := getLogLevelConsole(""); v != "" {
		cfg.ConsoleLevel = v
	}
	if getLogFileEnable() {
		cfg.FileEnabled = true
	}
	if v := getLogLevelFile(""); v != "" {
		cfg.FileLevel = v
	}
	if v := getLogFilePath(""); v != "" {
		cfg.FilePath = v
	}
	if v := getLogFileMaxSize(0); v != 0 {
		cfg.FileMaxSize = v
	}
	if v := getLogFileMaxBackups(0); v != 0 {
		cfg.FileMaxBackups = v
	}
	if v := getLogFileMaxAge(0); v != 0 {
		cfg.FileMaxAge = v
	}

	return NewWithConfig(cfg)
}

// GetLogger initializes the global default logger (singleton, thread-safe) from environment variables.
// The global logger uses CallerSkip(2) to correctly report the caller through package-level functions.
// Deprecated: Use New(), NewWithConfig(), or NewFromEnv() for instance-based logging.
// Kept for backward compatibility.
func GetLogger() (Logger, error) {
	once.Do(func() {
		defaultLogger, initErr = NewFromEnv()
		// Rebuild with CallerSkip(2) for correct caller attribution through global functions
		if initErr == nil {
			zl := defaultLogger.(*zapLogger)
			rebuild := zl.logger.WithOptions(zap.AddCallerSkip(1)) // +1 on top of existing 1 = 2
			defaultLogger = &zapLogger{
				sugar:  rebuild.Sugar(),
				logger: rebuild,
			}
		}
	})
	return defaultLogger, initErr
}

// SetDefault sets the global default logger instance used by package-level functions.
// The logger's CallerSkip is automatically adjusted (+1) for correct caller attribution
// through the global formatting.go wrapper.
func SetDefault(l Logger) {
	if zl, ok := l.(*zapLogger); ok {
		rebuild := zl.logger.WithOptions(zap.AddCallerSkip(1))
		defaultLogger = &zapLogger{
			sugar:  rebuild.Sugar(),
			logger: rebuild,
		}
	} else {
		defaultLogger = l
	}
}

// getDefault returns the global default logger, initializing it if needed.
func getDefault() Logger {
	if defaultLogger == nil {
		if _, err := GetLogger(); err != nil {
			panic(fmt.Sprintf("logging: failed to initialize default logger: %v", err))
		}
	}
	return defaultLogger
}
