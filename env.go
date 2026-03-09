package logging

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	LogLevelConsole   = "LOG_LEVEL_CONSOLE"
	LogFileEnable     = "LOG_FILE_ENABLE"
	LogLevelFile      = "LOG_LEVEL_FILE"
	LogFilePath       = "LOG_FILE_PATH"
	LogFileMaxSize    = "LOG_FILE_MAX_SIZE"
	LogFileMaxBackups = "LOG_FILE_MAX_BACKUPS"
	LogFileMaxAge     = "LOG_FILE_MAX_AGE"
)

// getLogFileEnable reads LOG_FILE_ENABLE from the environment.
func getLogFileEnable() bool {
	strValue := os.Getenv(LogFileEnable)
	value, _ := strconv.ParseBool(strValue)
	return value
}

// getLogLevelConsole reads LOG_LEVEL_CONSOLE from the environment.
// Returns def (lowercased) when the variable is empty.
func getLogLevelConsole(def string) string {
	strValue := os.Getenv(LogLevelConsole)
	if strValue == "" && def != "" {
		return strings.ToLower(def)
	}
	return strings.ToLower(strValue)
}

// getLogLevelFile reads LOG_LEVEL_FILE from the environment.
// Returns def (lowercased) when the variable is empty.
func getLogLevelFile(def string) string {
	strValue := os.Getenv(LogLevelFile)
	if strValue == "" && def != "" {
		return strings.ToLower(def)
	}
	return strings.ToLower(strValue)
}

// getLogFilePath reads LOG_FILE_PATH from the environment.
// Returns def when the variable is empty.
func getLogFilePath(def string) string {
	strValue := os.Getenv(LogFilePath)
	if strValue == "" && def != "" {
		return def
	}
	return strValue
}

// getLogFileMaxSize reads LOG_FILE_MAX_SIZE (megabytes) from the environment.
// Returns def when the variable is empty.
func getLogFileMaxSize(def int) int {
	strValue := os.Getenv(LogFileMaxSize)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

// getLogFileMaxBackups reads LOG_FILE_MAX_BACKUPS from the environment.
// Returns def when the variable is empty.
func getLogFileMaxBackups(def int) int {
	strValue := os.Getenv(LogFileMaxBackups)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

// getLogFileMaxAge reads LOG_FILE_MAX_AGE (days) from the environment.
// Returns def when the variable is empty.
func getLogFileMaxAge(def int) int {
	strValue := os.Getenv(LogFileMaxAge)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

// CheckEnvironments prints the current values of all logging-related
// environment variables to the standard logger. Useful for startup diagnostics.
func CheckEnvironments() {
	log.Println(strings.Repeat("*", 50))
	log.Printf("%s\t=\t%s\n", LogLevelConsole, getLogLevelConsole(""))
	log.Printf("%s\t=\t%v\n", LogFileEnable, getLogFileEnable())
	log.Printf("%s\t=\t%s\n", LogLevelFile, getLogLevelFile(""))
	log.Printf("%s\t=\t%s\n", LogFilePath, getLogFilePath(""))
	log.Printf("%s\t=\t%d\n", LogFileMaxSize, getLogFileMaxSize(0))
	log.Printf("%s\t=\t%d\n", LogFileMaxBackups, getLogFileMaxBackups(0))
	log.Printf("%s\t=\t%d\n", LogFileMaxAge, getLogFileMaxAge(0))
	log.Println(strings.Repeat("*", 50))
}
