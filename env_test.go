package logging

import (
	"os"
	"testing"
)

func clearEnv(t *testing.T) {
	t.Helper()
	envVars := []string{
		LogLevelConsole, LogFileEnable, LogLevelFile,
		LogFilePath, LogFileMaxSize, LogFileMaxBackups, LogFileMaxAge,
	}
	for _, v := range envVars {
		os.Unsetenv(v)
	}
}

// --- getLogFileEnable ---

func TestGetLogFileEnable(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		want     bool
	}{
		{"empty", "", false},
		{"true", "true", true},
		{"false", "false", false},
		{"1", "1", true},
		{"0", "0", false},
		{"invalid", "notabool", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogFileEnable, tt.envValue)
			}
			if got := getLogFileEnable(); got != tt.want {
				t.Errorf("getLogFileEnable() = %v, want %v", got, tt.want)
			}
		})
	}
}

// --- getLogLevelConsole ---

func TestGetLogLevelConsole(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      string
		want     string
	}{
		{"env set", "WARN", "debug", "warn"},
		{"env empty, default used", "", "info", "info"},
		{"env empty, default empty", "", "", ""},
		{"env uppercase", "ERROR", "", "error"},
		{"single char default", "", "e", "e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogLevelConsole, tt.envValue)
			}
			if got := getLogLevelConsole(tt.def); got != tt.want {
				t.Errorf("getLogLevelConsole(%q) = %q, want %q", tt.def, got, tt.want)
			}
		})
	}
}

// --- getLogLevelFile ---

func TestGetLogLevelFile(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      string
		want     string
	}{
		{"env set", "ERROR", "debug", "error"},
		{"env empty, default used", "", "warn", "warn"},
		{"both empty", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogLevelFile, tt.envValue)
			}
			if got := getLogLevelFile(tt.def); got != tt.want {
				t.Errorf("getLogLevelFile(%q) = %q, want %q", tt.def, got, tt.want)
			}
		})
	}
}

// --- getLogFilePath ---

func TestGetLogFilePath(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      string
		want     string
	}{
		{"env set", "/var/log/app.log", "default.log", "/var/log/app.log"},
		{"env empty, default used", "", "default.log", "default.log"},
		{"both empty", "", "", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogFilePath, tt.envValue)
			}
			if got := getLogFilePath(tt.def); got != tt.want {
				t.Errorf("getLogFilePath(%q) = %q, want %q", tt.def, got, tt.want)
			}
		})
	}
}

// --- getLogFileMaxSize ---

func TestGetLogFileMaxSize(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      int
		want     int
	}{
		{"env set", "200", 100, 200},
		{"env empty, default used", "", 100, 100},
		{"env invalid", "abc", 100, 0}, // strconv.Atoi returns 0 for invalid
		{"both zero", "", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogFileMaxSize, tt.envValue)
			}
			if got := getLogFileMaxSize(tt.def); got != tt.want {
				t.Errorf("getLogFileMaxSize(%d) = %d, want %d", tt.def, got, tt.want)
			}
		})
	}
}

// --- getLogFileMaxBackups ---

func TestGetLogFileMaxBackups(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      int
		want     int
	}{
		{"env set", "5", 7, 5},
		{"env empty, default used", "", 7, 7},
		{"both zero", "", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogFileMaxBackups, tt.envValue)
			}
			if got := getLogFileMaxBackups(tt.def); got != tt.want {
				t.Errorf("getLogFileMaxBackups(%d) = %d, want %d", tt.def, got, tt.want)
			}
		})
	}
}

// --- getLogFileMaxAge ---

func TestGetLogFileMaxAge(t *testing.T) {
	tests := []struct {
		name     string
		envValue string
		def      int
		want     int
	}{
		{"env set", "30", 5, 30},
		{"env empty, default used", "", 5, 5},
		{"both zero", "", 0, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clearEnv(t)
			if tt.envValue != "" {
				os.Setenv(LogFileMaxAge, tt.envValue)
			}
			if got := getLogFileMaxAge(tt.def); got != tt.want {
				t.Errorf("getLogFileMaxAge(%d) = %d, want %d", tt.def, got, tt.want)
			}
		})
	}
}
