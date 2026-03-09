package logging

import (
	"bytes"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// newTestLogger creates a logger that writes to a buffer for assertions.
func newTestLogger(level zapcore.Level) (Logger, *bytes.Buffer) {
	buf := &bytes.Buffer{}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(buf), level)
	l := zap.New(core)
	return &zapLogger{sugar: l.Sugar(), logger: l}, buf
}

// --- DefaultConfig ---

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg.ConsoleLevel != "debug" {
		t.Errorf("ConsoleLevel = %q, want %q", cfg.ConsoleLevel, "debug")
	}
	if cfg.FileEnabled {
		t.Error("FileEnabled should be false by default")
	}
	if cfg.FileLevel != "debug" {
		t.Errorf("FileLevel = %q, want %q", cfg.FileLevel, "debug")
	}
	if cfg.FilePath != "logs/example.log" {
		t.Errorf("FilePath = %q, want %q", cfg.FilePath, "logs/example.log")
	}
	if cfg.FileMaxSize != 100 {
		t.Errorf("FileMaxSize = %d, want 100", cfg.FileMaxSize)
	}
	if cfg.FileMaxBackups != 7 {
		t.Errorf("FileMaxBackups = %d, want 7", cfg.FileMaxBackups)
	}
	if cfg.FileMaxAge != 5 {
		t.Errorf("FileMaxAge = %d, want 5", cfg.FileMaxAge)
	}
	if !cfg.FileCompress {
		t.Error("FileCompress should be true by default")
	}
	if cfg.CallerSkip != 1 {
		t.Errorf("CallerSkip = %d, want 1", cfg.CallerSkip)
	}
}

// --- New ---

func TestNew_Default(t *testing.T) {
	logger, err := New()
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if logger == nil {
		t.Fatal("New() returned nil logger")
	}
	_ = logger.Sync()
}

func TestNew_WithOptions(t *testing.T) {
	logger, err := New(
		WithConsoleLevel("warn"),
		WithFileEnabled(false),
	)
	if err != nil {
		t.Fatalf("New() error: %v", err)
	}
	if logger == nil {
		t.Fatal("New() returned nil logger")
	}
	_ = logger.Sync()
}

// --- NewWithConfig ---

func TestNewWithConfig_ConsoleOnly(t *testing.T) {
	cfg := Config{
		ConsoleLevel: "info",
		FileEnabled:  false,
		CallerSkip:   1,
	}
	logger, err := NewWithConfig(cfg)
	if err != nil {
		t.Fatalf("NewWithConfig() error: %v", err)
	}
	if logger == nil {
		t.Fatal("NewWithConfig() returned nil logger")
	}
	_ = logger.Sync()
}

func TestNewWithConfig_FileEnabledNoPath(t *testing.T) {
	cfg := Config{
		ConsoleLevel: "info",
		FileEnabled:  true,
		FilePath:     "",
		CallerSkip:   1,
	}
	_, err := NewWithConfig(cfg)
	if err == nil {
		t.Fatal("NewWithConfig() should return error when FileEnabled=true and FilePath is empty")
	}
	if !strings.Contains(err.Error(), "FilePath is empty") {
		t.Errorf("error message = %q, should mention FilePath", err.Error())
	}
}

func TestNewWithConfig_WithFile(t *testing.T) {
	tmpFile := t.TempDir() + "/test.log"
	cfg := Config{
		ConsoleLevel:   "debug",
		FileEnabled:    true,
		FileLevel:      "info",
		FilePath:       tmpFile,
		FileMaxSize:    10,
		FileMaxBackups: 1,
		FileMaxAge:     1,
		FileCompress:   false,
		CallerSkip:     1,
	}
	logger, err := NewWithConfig(cfg)
	if err != nil {
		t.Fatalf("NewWithConfig() error: %v", err)
	}
	logger.Infof("test file message")
	_ = logger.Sync()

	data, err := os.ReadFile(tmpFile)
	if err != nil {
		t.Fatalf("cannot read log file: %v", err)
	}
	if !strings.Contains(string(data), "test file message") {
		t.Errorf("log file should contain 'test file message', got: %s", string(data))
	}
}

// --- NewFromEnv ---

func TestNewFromEnv_Defaults(t *testing.T) {
	clearEnv(t)
	logger, err := NewFromEnv()
	if err != nil {
		t.Fatalf("NewFromEnv() error: %v", err)
	}
	if logger == nil {
		t.Fatal("NewFromEnv() returned nil logger")
	}
	_ = logger.Sync()
}

func TestNewFromEnv_WithEnvVars(t *testing.T) {
	clearEnv(t)
	os.Setenv(LogLevelConsole, "warn")
	os.Setenv(LogFileEnable, "false")

	logger, err := NewFromEnv()
	if err != nil {
		t.Fatalf("NewFromEnv() error: %v", err)
	}
	if logger == nil {
		t.Fatal("NewFromEnv() returned nil logger")
	}
	_ = logger.Sync()
}

// --- GetLogger (singleton) ---

func TestGetLogger_Singleton(t *testing.T) {
	clearEnv(t)
	resetGlobalState()

	l1, err1 := GetLogger()
	if err1 != nil {
		t.Fatalf("GetLogger() error: %v", err1)
	}
	l2, err2 := GetLogger()
	if err2 != nil {
		t.Fatalf("GetLogger() second call error: %v", err2)
	}
	if l1 != l2 {
		t.Error("GetLogger() should return the same instance on subsequent calls")
	}
	resetGlobalState()
}

// --- Log level parsing ---

func TestParseLevel(t *testing.T) {
	tests := []struct {
		input    string
		fallback zapcore.Level
		want     zapcore.Level
	}{
		{"debug", zapcore.InfoLevel, zapcore.DebugLevel},
		{"info", zapcore.DebugLevel, zapcore.InfoLevel},
		{"warn", zapcore.InfoLevel, zapcore.WarnLevel},
		{"error", zapcore.InfoLevel, zapcore.ErrorLevel},
		{"", zapcore.InfoLevel, zapcore.InfoLevel},
		{"invalid", zapcore.WarnLevel, zapcore.WarnLevel},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseLevel(tt.input, tt.fallback)
			if got != tt.want {
				t.Errorf("parseLevel(%q, %v) = %v, want %v", tt.input, tt.fallback, got, tt.want)
			}
		})
	}
}

// --- zapLogger methods ---

func TestZapLogger_LogLevels(t *testing.T) {
	logger, buf := newTestLogger(zapcore.DebugLevel)

	logger.Debug("debug msg")
	logger.Info("info msg")
	logger.Warn("warn msg")
	logger.Error("error msg")

	output := buf.String()
	for _, msg := range []string{"debug msg", "info msg", "warn msg", "error msg"} {
		if !strings.Contains(output, msg) {
			t.Errorf("output should contain %q, got:\n%s", msg, output)
		}
	}
}

func TestZapLogger_FormattedLogLevels(t *testing.T) {
	logger, buf := newTestLogger(zapcore.DebugLevel)

	logger.Debugf("debug %s", "formatted")
	logger.Infof("info %d", 42)
	logger.Warnf("warn %v", true)
	logger.Errorf("error %s", "formatted")

	output := buf.String()
	for _, msg := range []string{"debug formatted", "info 42", "warn true", "error formatted"} {
		if !strings.Contains(output, msg) {
			t.Errorf("output should contain %q, got:\n%s", msg, output)
		}
	}
}

func TestZapLogger_LevelFiltering(t *testing.T) {
	logger, buf := newTestLogger(zapcore.WarnLevel)

	logger.Debug("should not appear")
	logger.Info("should not appear")
	logger.Warn("warn visible")
	logger.Error("error visible")

	output := buf.String()
	if strings.Contains(output, "should not appear") {
		t.Errorf("debug/info messages should be filtered at warn level, got:\n%s", output)
	}
	if !strings.Contains(output, "warn visible") {
		t.Errorf("warn message should appear, got:\n%s", output)
	}
	if !strings.Contains(output, "error visible") {
		t.Errorf("error message should appear, got:\n%s", output)
	}
}

// --- WithFields ---

func TestZapLogger_WithFields(t *testing.T) {
	logger, buf := newTestLogger(zapcore.DebugLevel)

	child := logger.WithFields(Fields{
		"user_id":    42,
		"request_id": "abc-123",
	})
	child.Info("request processed")

	output := buf.String()
	if !strings.Contains(output, "request processed") {
		t.Errorf("output should contain 'request processed', got:\n%s", output)
	}

	// Parse JSON to verify fields
	lines := strings.Split(strings.TrimSpace(output), "\n")
	lastLine := lines[len(lines)-1]
	var logEntry map[string]any
	if err := json.Unmarshal([]byte(lastLine), &logEntry); err != nil {
		t.Fatalf("failed to parse log entry JSON: %v", err)
	}
	if logEntry["user_id"] != float64(42) {
		t.Errorf("user_id = %v, want 42", logEntry["user_id"])
	}
	if logEntry["request_id"] != "abc-123" {
		t.Errorf("request_id = %v, want abc-123", logEntry["request_id"])
	}
}

func TestZapLogger_WithFields_DoesNotMutateParent(t *testing.T) {
	logger, buf := newTestLogger(zapcore.DebugLevel)

	_ = logger.WithFields(Fields{"child_field": "value"})
	logger.Info("parent message")

	output := buf.String()
	if strings.Contains(output, "child_field") {
		t.Errorf("parent logger should not contain child fields, got:\n%s", output)
	}
}

func TestZapLogger_WithFields_ReturnsLogger(t *testing.T) {
	logger, _ := newTestLogger(zapcore.DebugLevel)
	child := logger.WithFields(Fields{"key": "val"})

	// Verify it implements Logger interface
	var _ Logger = child
}

// --- Log() ---

func TestZapLogger_Log(t *testing.T) {
	logger, _ := newTestLogger(zapcore.DebugLevel)
	raw := logger.Log()
	if raw == nil {
		t.Fatal("Log() returned nil")
	}
}

// --- Sync ---

func TestZapLogger_Sync(t *testing.T) {
	logger, _ := newTestLogger(zapcore.DebugLevel)
	if err := logger.Sync(); err != nil {
		t.Errorf("Sync() error: %v", err)
	}
}

// --- SetDefault ---

func TestSetDefault(t *testing.T) {
	clearEnv(t)
	resetGlobalState()

	custom, _ := New(WithConsoleLevel("error"))
	SetDefault(custom)

	dl := getDefault()
	if dl == nil {
		t.Fatal("getDefault() returned nil after SetDefault")
	}
	resetGlobalState()
}

// --- Functional options ---

func TestFunctionalOptions(t *testing.T) {
	cfg := DefaultConfig()

	WithConsoleLevel("error")(&cfg)
	if cfg.ConsoleLevel != "error" {
		t.Errorf("WithConsoleLevel: got %q", cfg.ConsoleLevel)
	}

	WithFileEnabled(true)(&cfg)
	if !cfg.FileEnabled {
		t.Error("WithFileEnabled should set true")
	}

	WithFileLevel("warn")(&cfg)
	if cfg.FileLevel != "warn" {
		t.Errorf("WithFileLevel: got %q", cfg.FileLevel)
	}

	WithFilePath("/tmp/test.log")(&cfg)
	if cfg.FilePath != "/tmp/test.log" {
		t.Errorf("WithFilePath: got %q", cfg.FilePath)
	}

	WithFileMaxSize(50)(&cfg)
	if cfg.FileMaxSize != 50 {
		t.Errorf("WithFileMaxSize: got %d", cfg.FileMaxSize)
	}

	WithFileMaxBackups(3)(&cfg)
	if cfg.FileMaxBackups != 3 {
		t.Errorf("WithFileMaxBackups: got %d", cfg.FileMaxBackups)
	}

	WithFileMaxAge(10)(&cfg)
	if cfg.FileMaxAge != 10 {
		t.Errorf("WithFileMaxAge: got %d", cfg.FileMaxAge)
	}

	WithFileCompress(false)(&cfg)
	if cfg.FileCompress {
		t.Error("WithFileCompress should set false")
	}

	WithCallerSkip(2)(&cfg)
	if cfg.CallerSkip != 2 {
		t.Errorf("WithCallerSkip: got %d", cfg.CallerSkip)
	}
}
