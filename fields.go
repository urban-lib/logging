package logging

import (
	"fmt"
	"time"

	"go.uber.org/zap"
)

// Field represents a single structured log field.
// Use the typed constructors below to create fields efficiently.
type Field = zap.Field

// String creates a string field.
func String(key string, val string) Field { return zap.String(key, val) }

// Int creates an int field.
func Int(key string, val int) Field { return zap.Int(key, val) }

// Int64 creates an int64 field.
func Int64(key string, val int64) Field { return zap.Int64(key, val) }

// Float64 creates a float64 field.
func Float64(key string, val float64) Field { return zap.Float64(key, val) }

// Bool creates a bool field.
func Bool(key string, val bool) Field { return zap.Bool(key, val) }

// Err creates an "error" field. The key is always "error".
func Err(err error) Field { return zap.Error(err) }

// NamedErr creates an error field with a custom key.
func NamedErr(key string, err error) Field { return zap.NamedError(key, err) }

// Duration creates a time.Duration field.
func Duration(key string, val time.Duration) Field { return zap.Duration(key, val) }

// Time creates a time.Time field.
func Time(key string, val time.Time) Field { return zap.Time(key, val) }

// Any creates a field with an arbitrary value. The value is serialized using
// reflection if no specialized encoder is available.
func Any(key string, val any) Field { return zap.Any(key, val) }

// Stringer creates a field from a value that implements fmt.Stringer.
func Stringer(key string, val fmt.Stringer) Field { return zap.Stringer(key, val) }
