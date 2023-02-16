package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	docs "github.com/mooncool/gin-swagger-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type PingResponse struct {
	Message string `json:"message"` // return message
}

func main() {
	r := gin.Default()
	r.GET("/ping", ping)
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} PingResponse
// @Router /ping [get]
func ping(g *gin.Context) {
	g.JSON(http.StatusOK, PingResponse{"pong"})
}
