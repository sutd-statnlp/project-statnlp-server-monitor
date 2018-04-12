package resource

import (
	"github.com/gin-gonic/gin"

	"../service"
)

// InitLoadRoutes .
func InitLoadRoutes(router *gin.Engine) {
	loadService := service.LoadService{}

	router.GET("/api/load/average", func(context *gin.Context) {
		body, err := loadService.GetAverage()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

	router.GET("/api/load/misc", func(context *gin.Context) {
		body, err := loadService.GetMisc()
		if err == nil {
			context.JSON(200, body)
		} else {
			context.String(400, err.Error())
		}
	})

}
