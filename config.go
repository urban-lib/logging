package logging

// Config holds all configuration options for the logger.
type Config struct {
	// ConsoleLevel sets the minimum log level for console output.
	// Valid values: "debug", "info", "warn", "error", "fatal", "panic".
	// Default: "debug".
	ConsoleLevel string

	// FileEnabled enables writing logs to a file.
	// Default: false.
	FileEnabled bool

	// FileLevel sets the minimum log level for file output.
	// Valid values: "debug", "info", "warn", "error", "fatal", "panic".
	// Default: "debug".
	FileLevel string

	// FilePath is the path to the log file.
	// Default: "logs/example.log".
	FilePath string

	// FileMaxSize is the maximum size in megabytes before log rotation.
	// Default: 100.
	FileMaxSize int

	// FileMaxBackups is the maximum number of old log files to retain.
	// Default: 7.
	FileMaxBackups int

	// FileMaxAge is the maximum number of days to retain old log files.
	// Default: 5.
	FileMaxAge int

	// FileCompress determines if rotated log files should be compressed (gzip).
	// Default: true.
	FileCompress bool

	// CallerSkip is the number of callers to skip when determining the caller.
	// Default: 1 (skips the logging wrapper).
	CallerSkip int
}

// DefaultConfig returns configuration with sensible defaults.
func DefaultConfig() Config {
	return Config{
		ConsoleLevel:   "debug",
		FileEnabled:    false,
		FileLevel:      "debug",
		FilePath:       "logs/example.log",
		FileMaxSize:    100,
		FileMaxBackups: 7,
		FileMaxAge:     5,
		FileCompress:   true,
		CallerSkip:     1,
	}
}

// Option is a functional option for configuring the logger.
type Option func(*Config)

// WithConsoleLevel sets the console log level.
func WithConsoleLevel(level string) Option {
	return func(c *Config) {
		c.ConsoleLevel = level
	}
}

// WithFileEnabled enables or disables file logging.
func WithFileEnabled(enabled bool) Option {
	return func(c *Config) {
		c.FileEnabled = enabled
	}
}

// WithFileLevel sets the file log level.
func WithFileLevel(level string) Option {
	return func(c *Config) {
		c.FileLevel = level
	}
}

// WithFilePath sets the log file path.
func WithFilePath(path string) Option {
	return func(c *Config) {
		c.FilePath = path
	}
}

// WithFileMaxSize sets the maximum log file size in MB before rotation.
func WithFileMaxSize(size int) Option {
	return func(c *Config) {
		c.FileMaxSize = size
	}
}

// WithFileMaxBackups sets the maximum number of old log files to retain.
func WithFileMaxBackups(backups int) Option {
	return func(c *Config) {
		c.FileMaxBackups = backups
	}
}

// WithFileMaxAge sets the maximum number of days to retain old log files.
func WithFileMaxAge(days int) Option {
	return func(c *Config) {
		c.FileMaxAge = days
	}
}

// WithFileCompress enables or disables gzip compression for rotated files.
func WithFileCompress(compress bool) Option {
	return func(c *Config) {
		c.FileCompress = compress
	}
}

// WithCallerSkip sets the number of callers to skip.
func WithCallerSkip(skip int) Option {
	return func(c *Config) {
		c.CallerSkip = skip
	}
}
