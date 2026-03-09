package logging

import (
	"strings"

	"go.uber.org/zap/zapcore"
)

// containsAll returns true if s contains every one of the given substrings.
func containsAll(s string, subs ...string) bool {
	for _, sub := range subs {
		if !strings.Contains(s, sub) {
			return false
		}
	}
	return true
}

// zapcoreDebugLevel returns zapcore.DebugLevel.
// Keeps test files clean from importing zapcore directly where only the level is needed.
func zapcoreDebugLevel() zapcore.Level { return zapcore.DebugLevel }

// zapcoreInfoLevel returns zapcore.InfoLevel.
func zapcoreInfoLevel() zapcore.Level { return zapcore.InfoLevel }
