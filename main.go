package main

import (
	"github.com/gin-gonic/gin"

	"./controller"
	"./resource"
)

func main() {
	router := setupRoutes()

	router.Run(":8210")
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(func(context *gin.Context) {
		// add header Access-Control-Allow-Origin
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Next()
	})

	// Home
	controller.InitHomeRoutes(router)

	// Runtime
	resource.InitRuntimeRoutes(router)

	// Host
	resource.InitHostRoutes(router)

	// CPU
	resource.InitCPURoutes(router)

	// GPU
	resource.InitGPURoutes(router)

	// Memory
	resource.InitMemoryRoutes(router)

	// Disk
	resource.InitDiskRoutes(router)

	// Load
	resource.InitLoadRoutes(router)

	// Network
	resource.InitNetRoutes(router)

	return router
}
