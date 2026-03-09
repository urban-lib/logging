// Example usage of the logging/v3 package.
//
// The examples below demonstrate all supported configuration styles:
// global/env, struct-based, functional options, SetDefault, context-aware
// logging, sugar key-value methods, typed field helpers, and sampling.
//
// In production, set environment variables externally (shell, .env file, k8s ConfigMap)
// instead of calling os.Setenv. For local development you may use a .env loader
// such as github.com/joho/godotenv.
package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/urban-lib/logging/v3"
)

func main() {
	// ── Example 1: Global logger from env variables ──
	// Set env vars (in production, set these externally)
	os.Setenv(logging.LogLevelConsole, "DEBUG")
	os.Setenv(logging.LogFileEnable, "true")
	os.Setenv(logging.LogLevelFile, "DEBUG")
	os.Setenv(logging.LogFilePath, "logs/test.log")
	os.Setenv(logging.LogFileMaxSize, "50")
	os.Setenv(logging.LogFileMaxBackups, "2")
	os.Setenv(logging.LogFileMaxAge, "1")

	logging.CheckEnvironments()

	// Package-level functions auto-initialize from env
	logging.Debug("Debug message (global)")
	logging.Infof("Info message (global)")
	logging.Warnf("Warning message (global)")

	logging.WithFields(logging.Fields{"user": "alice"}).Infof("User logged in (global)")

	// ── Example 2: Instance-based logger with Config struct ──
	logger, err := logging.NewWithConfig(logging.Config{
		ConsoleLevel: "info",
		FileEnabled:  false,
		CallerSkip:   1,
	})
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	logger.Infof("Instance logger: info message")
	logger.WithFields(logging.Fields{"request_id": "abc-123"}).Warnf("Instance logger: with fields")

	// ── Example 3: Instance-based logger with functional options ──
	logger2, err := logging.New(
		logging.WithConsoleLevel("debug"),
		logging.WithFileEnabled(true),
		logging.WithFilePath("logs/options.log"),
		logging.WithFileLevel("error"),
	)
	if err != nil {
		panic(err)
	}
	defer logger2.Sync()

	logger2.Debugf("Options logger: debug message")
	logger2.Errorf("Options logger: error message")

	// ── Example 4: Replace default global logger with custom instance ──
	logging.SetDefault(logger)
	logging.Infof("Now using custom instance as global default")

	// ── Example 5: Sugar key-value API ──
	logger.Infow("request handled",
		"method", "POST",
		"path", "/api/users",
		"latency_ms", 42,
	)

	// ── Example 6: Context-aware logging ──
	ctx := context.Background()
	ctx = logging.ContextWithFields(ctx, logging.Fields{
		"trace_id":   "abc-xyz-123",
		"request_id": "req-456",
	})
	logger.WithContext(ctx).Infof("Processing order")
	// Global variant
	logging.WithContext(ctx).Warnf("Slow downstream call")

	// ── Example 7: Typed field helpers ──
	logger.Log().Info("structured fields",
		logging.String("service", "payment"),
		logging.Int("attempt", 3),
		logging.Duration("latency", 150*time.Millisecond),
		logging.Bool("success", true),
		logging.Err(fmt.Errorf("timeout")),
	)

	// ── Example 8: Sampling (high-throughput) ──
	sampledLogger, err := logging.New(
		logging.WithConsoleLevel("info"),
		logging.WithSampling(5, 3), // 5 messages/sec initial, then every 3rd
	)
	if err != nil {
		panic(err)
	}
	defer sampledLogger.Sync()

	for i := 0; i < 20; i++ {
		sampledLogger.Infof("High-throughput message #%d", i)
	}

	fmt.Println("All examples completed successfully.")
}
