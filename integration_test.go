package logging

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestIntegration_FileLogging(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "integration.log")

	logger, err := NewWithConfig(Config{
		ConsoleLevel:   "debug",
		FileEnabled:    true,
		FileLevel:      "info",
		FilePath:       logFile,
		FileMaxSize:    1, // 1 MB
		FileMaxBackups: 1,
		FileMaxAge:     1,
		FileCompress:   false,
		CallerSkip:     0,
	})
	if err != nil {
		t.Fatalf("NewWithConfig error: %v", err)
	}

	// Write messages at different levels
	logger.Debug("debug only console") // below file level, should NOT appear in file
	logger.Info("info file message")
	logger.Warn("warn file message")
	logger.Error("error file message")
	logger.WithFields(Fields{"key": "val"}).Info("structured message")
	_ = logger.Sync()

	// Read and verify file content
	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("cannot read log file: %v", err)
	}
	content := string(data)

	// Debug should NOT be in file (level filter)
	if strings.Contains(content, "debug only console") {
		t.Error("debug message should not appear in file at info level")
	}

	// Info, Warn, Error should be present
	for _, msg := range []string{"info file message", "warn file message", "error file message"} {
		if !strings.Contains(content, msg) {
			t.Errorf("file should contain %q", msg)
		}
	}

	// Verify JSON format
	lines := strings.Split(strings.TrimSpace(content), "\n")
	if len(lines) < 4 {
		t.Fatalf("expected at least 4 lines, got %d", len(lines))
	}

	// Verify structured message has the field
	lastLine := lines[len(lines)-1]
	var entry map[string]any
	if err := json.Unmarshal([]byte(lastLine), &entry); err != nil {
		t.Fatalf("failed to parse JSON log entry: %v\nline: %s", err, lastLine)
	}
	if entry["key"] != "val" {
		t.Errorf("structured field 'key' = %v, want 'val'", entry["key"])
	}
	if entry["msg"] != "structured message" {
		t.Errorf("msg = %v, want 'structured message'", entry["msg"])
	}
}

func TestIntegration_FileLogging_MultipleMessages(t *testing.T) {
	tmpDir := t.TempDir()
	logFile := filepath.Join(tmpDir, "multi.log")

	logger, err := NewWithConfig(Config{
		ConsoleLevel:   "error", // suppress console noise in tests
		FileEnabled:    true,
		FileLevel:      "debug",
		FilePath:       logFile,
		FileMaxSize:    1,
		FileMaxBackups: 1,
		FileMaxAge:     1,
		FileCompress:   false,
		CallerSkip:     0,
	})
	if err != nil {
		t.Fatalf("NewWithConfig error: %v", err)
	}

	for i := 0; i < 100; i++ {
		logger.Infof("message %d", i)
	}
	_ = logger.Sync()

	data, err := os.ReadFile(logFile)
	if err != nil {
		t.Fatalf("cannot read log file: %v", err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	if len(lines) != 100 {
		t.Errorf("expected 100 log lines, got %d", len(lines))
	}
}
