package router

import (
	"fmt"
	"wiredcraft/docs"
	"wiredcraft/logger"
	"wiredcraft/store"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

type Param struct {
	Name string `json:"name"`
}

// @BasePath /

// GetWelcome godoc
// @Summary get welcome
// @Schemes
// @Description return hello name
// @Tags example
// @Accept json
// @Produce text/plain
// @Success 200 {string} hello name
// @Router /welcome [get]
func GetWelcome(c *gin.Context) {
	name := store.GetStore().Get()

	if name == "" {
		c.String(200, "hello wiredcraft")
	} else {
		c.String(200, fmt.Sprintf("hello %s", name))
	}
}

// @BasePath /

// PetWelcome godoc
// @Summary put welcome
// @Schemes
// @Description return success
// @Param name body Param true "name"
// @Tags example
// @Accept json
// @Produce text/plain
// @Success 200 {string} success
// @Router /welcome [put]
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

	docs.SwaggerInfo.BasePath = "/"

	r.GET("/welcome", GetWelcome)
	r.PUT("/welcome", PutWelcome)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1)))

	return r
}
