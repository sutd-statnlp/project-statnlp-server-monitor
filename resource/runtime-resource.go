package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// RuntimeResource .
type RuntimeResource struct {
}

// InitRoutes .
func (runtimeResource RuntimeResource) InitRoutes(router *gin.Engine) {
	runtimeService := service.RuntimeService{}

	router.GET("/api/runtime/goos", func(context *gin.Context) {
		body := runtimeService.GetGOOS()
		context.JSON(200, body)
	})

}
