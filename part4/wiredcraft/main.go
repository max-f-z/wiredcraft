package main

import (
	"fmt"
	"sync"
	"time"

	"wiredcraft/logger"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var name string
var mutex sync.RWMutex

type Param struct {
	Name string `json:"name"`
}

func main() {
	r := gin.New()

	r.Use(ginzap.Ginzap(logger.Logger(), time.RFC3339, true))

	r.GET("/welcome", func(c *gin.Context) {
		if name == "" {
			c.String(200, "hello wiredcraft")
		} else {
			mutex.RLock()
			defer mutex.RUnlock()

			c.String(200, fmt.Sprintf("hello %s", name))
		}
	})

	r.PUT("/welcome", func(c *gin.Context) {
		params := Param{}
		err := c.BindJSON(&params)

		if err != nil {
			logger.Logger().Error("error unmarhsalling request", zap.Any("error", err))
			c.String(400, "bad request")
			return
		}

		mutex.Lock()
		defer mutex.Unlock()
		name = params.Name

		c.String(200, "success")
	})

	r.Run()
}
