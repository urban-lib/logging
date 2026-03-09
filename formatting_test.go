package logging

import (
	"bytes"
	"strings"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// setupGlobalTestLogger sets the global default logger to write to a buffer and returns it.
func setupGlobalTestLogger(t *testing.T, level zapcore.Level) *bytes.Buffer {
	t.Helper()
	resetGlobalState()

	buf := &bytes.Buffer{}
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, zapcore.AddSync(buf), level)
	l := zap.New(core)
	// Use CallerSkip(2) to match the global wrapper depth
	l = l.WithOptions(zap.AddCallerSkip(2))
	defaultLogger = &zapLogger{sugar: l.Sugar(), logger: l}

	t.Cleanup(func() {
		resetGlobalState()
	})
	return buf
}

func TestGlobal_Debugf(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Debugf("hello %s", "world")
	if !strings.Contains(buf.String(), "hello world") {
		t.Errorf("Debugf: expected 'hello world', got: %s", buf.String())
	}
}

func TestGlobal_Infof(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Infof("info %d", 42)
	if !strings.Contains(buf.String(), "info 42") {
		t.Errorf("Infof: expected 'info 42', got: %s", buf.String())
	}
}

func TestGlobal_Warnf(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Warnf("warn %v", true)
	if !strings.Contains(buf.String(), "warn true") {
		t.Errorf("Warnf: expected 'warn true', got: %s", buf.String())
	}
}

func TestGlobal_Errorf(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Errorf("error %s", "msg")
	if !strings.Contains(buf.String(), "error msg") {
		t.Errorf("Errorf: expected 'error msg', got: %s", buf.String())
	}
}

func TestGlobal_Debug(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Debug("debug plain")
	if !strings.Contains(buf.String(), "debug plain") {
		t.Errorf("Debug: expected 'debug plain', got: %s", buf.String())
	}
}

func TestGlobal_Info(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Info("info plain")
	if !strings.Contains(buf.String(), "info plain") {
		t.Errorf("Info: expected 'info plain', got: %s", buf.String())
	}
}

func TestGlobal_Warn(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Warn("warn plain")
	if !strings.Contains(buf.String(), "warn plain") {
		t.Errorf("Warn: expected 'warn plain', got: %s", buf.String())
	}
}

func TestGlobal_Error(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	Error("error plain")
	if !strings.Contains(buf.String(), "error plain") {
		t.Errorf("Error: expected 'error plain', got: %s", buf.String())
	}
}

func TestGlobal_WithFields(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcore.DebugLevel)
	WithFields(Fields{"key": "value"}).Infof("with fields msg")
	output := buf.String()
	if !strings.Contains(output, "with fields msg") {
		t.Errorf("WithFields: expected 'with fields msg', got: %s", output)
	}
	if !strings.Contains(output, "key") || !strings.Contains(output, "value") {
		t.Errorf("WithFields: expected key/value in output, got: %s", output)
	}
}

func TestGlobal_LazyInit(t *testing.T) {
	clearEnv(t)
	resetGlobalState()
	defer resetGlobalState()

	// Calling a global function without prior GetLogger should not panic
	// (lazy init kicks in)
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("global function panicked on lazy init: %v", r)
			}
		}()
		Debug("lazy init test")
	}()
}
