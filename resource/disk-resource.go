package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// DiskResource .
type DiskResource struct {
}

// InitRoutes .
func (diskResource DiskResource) InitRoutes(router *gin.Engine) {
	diskService := service.DiskService{}

	router.GET("/api/disk/usage", func(context *gin.Context) {
		body, err := diskService.GetUsage()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

}
