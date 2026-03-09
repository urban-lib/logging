package logging

import (
	"errors"
	"testing"
	"time"
)

func TestFieldHelpers_String(t *testing.T) {
	f := String("name", "alice")
	if f.Key != "name" {
		t.Errorf("expected key=name, got %s", f.Key)
	}
}

func TestFieldHelpers_Int(t *testing.T) {
	f := Int("count", 42)
	if f.Key != "count" {
		t.Errorf("expected key=count, got %s", f.Key)
	}
	if f.Integer != 42 {
		t.Errorf("expected value=42, got %d", f.Integer)
	}
}

func TestFieldHelpers_Int64(t *testing.T) {
	f := Int64("big", 9999999999)
	if f.Key != "big" {
		t.Errorf("expected key=big, got %s", f.Key)
	}
}

func TestFieldHelpers_Float64(t *testing.T) {
	f := Float64("pi", 3.14)
	if f.Key != "pi" {
		t.Errorf("expected key=pi, got %s", f.Key)
	}
}

func TestFieldHelpers_Bool(t *testing.T) {
	f := Bool("active", true)
	if f.Key != "active" {
		t.Errorf("expected key=active, got %s", f.Key)
	}
}

func TestFieldHelpers_Err(t *testing.T) {
	f := Err(errors.New("boom"))
	if f.Key != "error" {
		t.Errorf("expected key=error, got %s", f.Key)
	}
}

func TestFieldHelpers_NamedErr(t *testing.T) {
	f := NamedErr("cause", errors.New("oops"))
	if f.Key != "cause" {
		t.Errorf("expected key=cause, got %s", f.Key)
	}
}

func TestFieldHelpers_Duration(t *testing.T) {
	f := Duration("latency", 150*time.Millisecond)
	if f.Key != "latency" {
		t.Errorf("expected key=latency, got %s", f.Key)
	}
}

func TestFieldHelpers_Time(t *testing.T) {
	now := time.Now()
	f := Time("timestamp", now)
	if f.Key != "timestamp" {
		t.Errorf("expected key=timestamp, got %s", f.Key)
	}
}

func TestFieldHelpers_Any(t *testing.T) {
	f := Any("data", map[string]int{"a": 1})
	if f.Key != "data" {
		t.Errorf("expected key=data, got %s", f.Key)
	}
}

func TestFieldHelpers_Stringer(t *testing.T) {
	f := Stringer("addr", stringerAddr("127.0.0.1"))
	if f.Key != "addr" {
		t.Errorf("expected key=addr, got %s", f.Key)
	}
}

// stringerAddr implements fmt.Stringer for testing.
type stringerAddr string

func (s stringerAddr) String() string { return string(s) }
