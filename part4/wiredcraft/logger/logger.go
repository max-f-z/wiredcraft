package logger

import (
	"os"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

func Logger() *zap.Logger {
	if logger != nil {
		return logger
	}

	var loggerConfig zap.Config

	if string(gin.ReleaseMode) == os.Getenv("GIN_MODE") {
		loggerConfig = zap.NewProductionConfig()
	} else {
		loggerConfig = zap.NewDevelopmentConfig()
	}

	logger, err := loggerConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
