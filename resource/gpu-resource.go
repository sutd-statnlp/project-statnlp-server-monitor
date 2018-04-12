package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
	"../util"
)

// InitGPURoutes .
func InitGPURoutes(router *gin.Engine) {
	gpuService := service.GPUService{}

	router.GET("/api/gpu/info", func(context *gin.Context) {
		body, err := gpuService.GetInfo()
		if err == nil {
			context.Header(util.GetJSONHeader())
			context.String(200, body)
		} else {
			context.String(400, err.Error())
		}
	})
}
