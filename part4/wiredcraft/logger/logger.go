package logger

import "go.uber.org/zap"

var logger *zap.Logger

func Logger() *zap.Logger {
	if logger != nil {
		return logger
	}

	loggerConfig := zap.NewDevelopmentConfig()

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
