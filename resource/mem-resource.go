package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// InitMemoryRoutes .
func InitMemoryRoutes(router *gin.Engine) {
	memService := service.MemService{}

	router.GET("/api/mem/virtual", func(context *gin.Context) {
		body, err := memService.GetVirtualMemory()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

	router.GET("/api/mem/swap", func(context *gin.Context) {
		body, err := memService.GetSwapMemory()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

}
