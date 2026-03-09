package logging

import (
	"testing"
)

func TestZapLogger_Debugw(t *testing.T) {
	logger, buf := newTestLogger(zapcoreDebugLevel())
	logger.Debugw("event", "key", "value", "num", 42)

	output := buf.String()
	if !containsAll(output, "event", "key", "value", "num") {
		t.Errorf("expected key-value pairs in output, got: %s", output)
	}
}

func TestZapLogger_Infow(t *testing.T) {
	logger, buf := newTestLogger(zapcoreDebugLevel())
	logger.Infow("request", "method", "GET", "path", "/api")

	output := buf.String()
	if !containsAll(output, "request", "method", "GET", "path", "/api") {
		t.Errorf("expected key-value pairs in output, got: %s", output)
	}
}

func TestZapLogger_Warnw(t *testing.T) {
	logger, buf := newTestLogger(zapcoreDebugLevel())
	logger.Warnw("slow", "latency_ms", 500)

	output := buf.String()
	if !containsAll(output, "slow", "latency_ms") {
		t.Errorf("expected key-value pairs in output, got: %s", output)
	}
}

func TestZapLogger_Errorw(t *testing.T) {
	logger, buf := newTestLogger(zapcoreDebugLevel())
	logger.Errorw("failed", "err", "timeout")

	output := buf.String()
	if !containsAll(output, "failed", "err", "timeout") {
		t.Errorf("expected key-value pairs in output, got: %s", output)
	}
}

func TestGlobal_Debugw(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcoreDebugLevel())
	Debugw("test", "k", "v")

	if !containsAll(buf.String(), "test", "k", "v") {
		t.Errorf("expected global Debugw output, got: %s", buf.String())
	}
}

func TestGlobal_Infow(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcoreDebugLevel())
	Infow("test", "k", "v")

	if !containsAll(buf.String(), "test", "k", "v") {
		t.Errorf("expected global Infow output, got: %s", buf.String())
	}
}

func TestGlobal_Warnw(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcoreDebugLevel())
	Warnw("test", "k", "v")

	if !containsAll(buf.String(), "test", "k", "v") {
		t.Errorf("expected global Warnw output, got: %s", buf.String())
	}
}

func TestGlobal_Errorw(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcoreDebugLevel())
	Errorw("test", "k", "v")

	if !containsAll(buf.String(), "test", "k", "v") {
		t.Errorf("expected global Errorw output, got: %s", buf.String())
	}
}
