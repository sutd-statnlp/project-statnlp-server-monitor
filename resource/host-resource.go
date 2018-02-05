package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// HostResource .
type HostResource struct {
}

// InitRoutes .
func (hostResource HostResource) InitRoutes(router *gin.Engine) {
	hostService := service.HostService{}

	router.GET("/api/host/info", func(context *gin.Context) {
		body, err := hostService.GetInfo()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.Error(err)
		}
	})

}
