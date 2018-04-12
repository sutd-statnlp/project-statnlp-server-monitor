package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// NetResource .
type NetResource struct {
}

// InitRoutes .
func (netResource NetResource) InitRoutes(router *gin.Engine) {
	netService := service.NetService{}

	router.GET("/api/net/interface", func(context *gin.Context) {
		body, err := netService.GetInterfaces()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

	router.GET("/api/net/connection", func(context *gin.Context) {
		body, err := netService.GetConnections()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})
}
