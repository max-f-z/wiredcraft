package main

import (
	"time"

	"wiredcraft/logger"
	"wiredcraft/router"
	"wiredcraft/store"

	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"
)

func main() {
	store.InitValue()

	r := router.SetupRouter()
	r.Use(ginzap.Ginzap(logger.Logger(), time.RFC3339, true))
	err := r.Run()

	logger.Logger().Error("error occurred", zap.Any("error", err))
}
