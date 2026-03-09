package main

import (
	"fmt"
	"os"

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

	fmt.Println("All examples completed successfully.")
}
