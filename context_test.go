package logging

import (
	"context"
	"testing"
)

func TestContextWithFields_Basic(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithFields(ctx, Fields{"request_id": "abc-123"})

	got := FieldsFromContext(ctx)
	if got == nil {
		t.Fatal("expected fields, got nil")
	}
	if got["request_id"] != "abc-123" {
		t.Errorf("expected request_id=abc-123, got %v", got["request_id"])
	}
}

func TestContextWithFields_Merge(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithFields(ctx, Fields{"a": 1})
	ctx = ContextWithFields(ctx, Fields{"b": 2})

	got := FieldsFromContext(ctx)
	if len(got) != 2 {
		t.Fatalf("expected 2 fields, got %d", len(got))
	}
	if got["a"] != 1 || got["b"] != 2 {
		t.Errorf("unexpected fields: %v", got)
	}
}

func TestContextWithFields_Override(t *testing.T) {
	ctx := context.Background()
	ctx = ContextWithFields(ctx, Fields{"k": "old"})
	ctx = ContextWithFields(ctx, Fields{"k": "new"})

	got := FieldsFromContext(ctx)
	if got["k"] != "new" {
		t.Errorf("expected k=new, got %v", got["k"])
	}
}

func TestFieldsFromContext_Nil(t *testing.T) {
	//nolint:staticcheck // testing nil context behavior intentionally
	got := FieldsFromContext(nil)
	if got != nil {
		t.Errorf("expected nil, got %v", got)
	}
}

func TestFieldsFromContext_Empty(t *testing.T) {
	got := FieldsFromContext(context.Background())
	if got != nil {
		t.Errorf("expected nil for empty context, got %v", got)
	}
}

func TestZapLogger_WithContext(t *testing.T) {
	logger, buf := newTestLogger(zapcoreDebugLevel())
	ctx := ContextWithFields(context.Background(), Fields{"trace_id": "xyz"})

	logger.WithContext(ctx).Infof("traced")

	output := buf.String()
	if !containsAll(output, "traced", "trace_id", "xyz") {
		t.Errorf("expected trace_id=xyz in output, got: %s", output)
	}
}

func TestZapLogger_WithContext_NoFields(t *testing.T) {
	logger, _ := newTestLogger(zapcoreDebugLevel())
	ctx := context.Background()

	result := logger.WithContext(ctx)
	if result != logger {
		t.Error("WithContext with empty context should return the same logger")
	}
}

func TestGlobal_WithContext(t *testing.T) {
	buf := setupGlobalTestLogger(t, zapcoreDebugLevel())
	ctx := ContextWithFields(context.Background(), Fields{"user": "bob"})

	WithContext(ctx).Infof("user action")

	output := buf.String()
	if !containsAll(output, "user action", "user", "bob") {
		t.Errorf("expected context fields in output, got: %s", output)
	}
}
