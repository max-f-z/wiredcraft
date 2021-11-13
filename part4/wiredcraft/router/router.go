package router

import (
	"fmt"
	"wiredcraft/logger"
	"wiredcraft/store"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Param struct {
	Name string `json:"name"`
}

func GetWelcome(c *gin.Context) {
	name := store.GetStore().Get()

	if name == "" {
		c.String(200, "hello wiredcraft")
	} else {
		c.String(200, fmt.Sprintf("hello %s", name))
	}
}

func PutWelcome(c *gin.Context) {
	params := Param{}
	err := c.BindJSON(&params)

	if err != nil {
		logger.Logger().Error("error unmarhsalling request", zap.Any("error", err))
		c.String(400, "bad request")
		return
	}

	err = store.GetStore().Set(params.Name)
	if err != nil {
		logger.Logger().Error("error setting to redis", zap.Any("error", err))
		c.String(500, "internal error")
		return
	}

	c.String(200, "success")
}

func SetupRouter() *gin.Engine {
	r := gin.New()

	r.GET("/welcome", GetWelcome)
	r.PUT("/welcome", PutWelcome)

	return r
}
