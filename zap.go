package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"sync"
)

var (
	logger *zap.SugaredLogger
	once   sync.Once
)

func logLevel() zapcore.Level {
	if level, err := zapcore.ParseLevel(getLogLevelConsole("debug")); err != nil {
		return zapcore.InfoLevel
	} else {
		return level
	}
}

func logLevelFile() zapcore.Level {
	if level, err := zapcore.ParseLevel(getLogLevelFile("debug")); err != nil {
		return zapcore.InfoLevel
	} else {
		return level
	}
}

func textEncoder() zapcore.Encoder {
	devConfig := zap.NewDevelopmentEncoderConfig()
	//devConfig.FunctionKey = "func"
	devConfig.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewConsoleEncoder(devConfig)
}

func jsonEncoder() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncodeCaller = zapcore.FullCallerEncoder
	return zapcore.NewJSONEncoder(cfg)
}

func GetLogger() (*zap.SugaredLogger, error) {

	once.Do(func() {
		cores := make([]zapcore.Core, 0)

		core := zapcore.NewCore(textEncoder(), zapcore.Lock(os.Stdout), logLevel())
		cores = append(cores, core)

		if getLogFileEnable() {
			writer := zapcore.AddSync(&lumberjack.Logger{
				Filename:   getLogFilePath("logs/example.log"),
				MaxSize:    getLogFileMaxSize(100),
				MaxAge:     getLogFileMaxAge(5),
				MaxBackups: getLogFileMaxBackups(7),
				Compress:   true,
			})
			core = zapcore.NewCore(jsonEncoder(), writer, logLevelFile())
			cores = append(cores, core)
		}

		combinedCore := zapcore.NewTee(cores...)

		l := zap.New(
			combinedCore,
			zap.AddCallerSkip(1),
			zap.AddCaller(),
		)
		logger = l.Sugar()

	})

	return logger, nil
}
