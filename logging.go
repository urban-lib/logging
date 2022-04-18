package logging

var log Logger

type Fields map[string]interface{}

const (
	// Debug string
	Debug = "debug"
	// Info string
	Info = "info"
	// Warn string
	Warn = "warning"
	// Error string
	Error = "error"
	// Fatal string
	Fatal = "fatal"
)

// Configuration ...
type Configuration struct {
	EnableConsole     bool   `mapstructure:"enableConsole"`
	ConsoleJSONFormat bool   `mapstructure:"consoleJsonFormat"`
	ConsoleLevel      string `mapstructure:"consoleLevel"`
	EnableFile        bool   `mapstructure:"enableFile"`
	FileJSONFormat    bool   `mapstructure:"fileJsonFormat"`
	FileLevel         string `mapstructure:"fileLevel"`
	FileLocation      string `mapstructure:"fileLocation"`
}

type Logger interface {
	Debugf(f string, args ...interface{})
	Infof(f string, args ...interface{})
	Warnf(f string, args ...interface{})
	Errorf(f string, args ...interface{})
	Fatalf(f string, args ...interface{})
	Panicf(f string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

func New(config Configuration) error {
	logger, err := newZapLogger(config)
	if err != nil {
		return err
	}
	log = logger
	return nil
}

func GetLog() Logger {
	return log
}
