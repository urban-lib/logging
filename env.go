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

func getLogFileEnable() bool {
	strValue := os.Getenv(LogFileEnable)
	value, _ := strconv.ParseBool(strValue)
	return value
}

func getLogLevelConsole(def string) string {
	strValue := os.Getenv(LogLevelConsole)
	if len(strValue) == 0 && len(def) > 1 {
		return strings.ToLower(def)
	}
	return strings.ToLower(strValue)
}

func getLogLevelFile(def string) string {
	strValue := os.Getenv(LogLevelFile)
	if strValue == "" && def != "" {
		return strings.ToLower(def)
	}
	return strings.ToLower(strValue)
}

func getLogFilePath(def string) string {
	strValue := os.Getenv(LogFilePath)
	if strValue == "" && def != "" {
		return def
	}
	return strValue
}

func getLogFileMaxSize(def int) int {
	strValue := os.Getenv(LogFileMaxSize)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

func getLogFileMaxBackups(def int) int {
	strValue := os.Getenv(LogFileMaxBackups)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

func getLogFileMaxAge(def int) int {
	strValue := os.Getenv(LogFileMaxAge)
	if strValue == "" && def != 0 {
		return def
	}
	value, _ := strconv.Atoi(strValue)
	return value
}

func CheckEnvironments() {
	log.Println(strings.Repeat("*", 50))
	log.Printf("%s\t=\t%s\n", LogLevelConsole, getLogLevelConsole(""))
	log.Printf("%s\t=\t%v\n", LogFileEnable, getLogFileEnable())
	log.Printf("%s\t=\t%s\n", LogLevelFile, getLogLevelFile(""))
	log.Printf("%s\t=\t%s\n", LogFilePath, getLogFilePath(""))
	log.Printf("%s\t=\t%d\n", LogFileMaxSize, getLogFileMaxSize(0))
	log.Printf("%s=\t%d\n", LogFileMaxBackups, getLogFileMaxBackups(0))
	log.Printf("%s\t=\t%d\n", LogFileMaxAge, getLogFileMaxAge(0))
	log.Println(strings.Repeat("*", 50))
}
