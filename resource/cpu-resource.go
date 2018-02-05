package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// CPUResource .
type CPUResource struct {
}

// InitRoutes .
func (cpuResource CPUResource) InitRoutes(router *gin.Engine) {
	cpuService := service.CPUService{}

	router.GET("/api/cpu/info", func(context *gin.Context) {
		body, err := cpuService.GetInfo()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.Error(err)
		}
	})

	router.GET("/api/cpu/percent", func(context *gin.Context) {
		body, err := cpuService.GetPercent()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.Error(err)
		}
	})
}
