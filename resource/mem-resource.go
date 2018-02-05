package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// MemResource .
type MemResource struct {
}

// InitRoutes .
func (memResource MemResource) InitRoutes(router *gin.Engine) {
	memService := service.MemService{}

	router.GET("/api/mem/virtual", func(context *gin.Context) {
		body, err := memService.GetVirtualMemory()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.Error(err)
		}
	})

}
