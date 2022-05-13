package logging

// ConfigurationInterface ...
type ConfigurationInterface interface {
	GetEnableConsole() bool
	GetConsoleJSONFormat() bool
	GetConsoleLevel() string
	GetEnableFile() bool
	GetFileJSONFormat() bool
	GetFileLevel() string
	GetFileLocation() string
}

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

func (c *Configuration) GetEnableConsole() bool {
	return c.EnableConsole
}

func (c *Configuration) GetConsoleJSONFormat() bool {
	return c.ConsoleJSONFormat
}

func (c *Configuration) GetConsoleLevel() string {
	return c.ConsoleLevel
}

func (c *Configuration) GetEnableFile() bool {
	return c.EnableFile
}

func (c *Configuration) GetFileJSONFormat() bool {
	return c.FileJSONFormat
}

func (c *Configuration) GetFileLevel() string {
	return c.FileLevel
}

func (c *Configuration) GetFileLocation() string {
	return c.FileLocation
}
