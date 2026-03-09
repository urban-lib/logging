package logging

import (
	"testing"
)

func TestWithSampling_Option(t *testing.T) {
	cfg := DefaultConfig()
	opt := WithSampling(100, 10)
	opt(&cfg)

	if cfg.SamplingInitial != 100 {
		t.Errorf("expected SamplingInitial=100, got %d", cfg.SamplingInitial)
	}
	if cfg.SamplingThereafter != 10 {
		t.Errorf("expected SamplingThereafter=10, got %d", cfg.SamplingThereafter)
	}
}

func TestNew_WithSampling(t *testing.T) {
	logger, err := New(
		WithConsoleLevel("info"),
		WithSampling(100, 10),
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if logger == nil {
		t.Fatal("expected non-nil logger")
	}
	// Verify it's functional — just log without panic
	logger.Infof("sampled message")
}

func TestNew_WithoutSampling(t *testing.T) {
	// SamplingInitial=0 means disabled (default)
	cfg := DefaultConfig()
	if cfg.SamplingInitial != 0 {
		t.Errorf("expected SamplingInitial=0 by default, got %d", cfg.SamplingInitial)
	}
	if cfg.SamplingThereafter != 0 {
		t.Errorf("expected SamplingThereafter=0 by default, got %d", cfg.SamplingThereafter)
	}
}
