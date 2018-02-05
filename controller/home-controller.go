package controller

import (
	"github.com/gin-gonic/gin"
)

// HomeController .
type HomeController struct {
}

// InitRoutes .
func (homeController HomeController) InitRoutes(router *gin.Engine) {
	router.GET("/", func(context *gin.Context) {
		context.String(200, "Welcome %s", "to StatNLP monitor")
	})
}
