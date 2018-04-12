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

	homeController := controller.HomeController{}
	homeController.InitRoutes(router)

	// Runtime
	runtimeResource := resource.RuntimeResource{}
	runtimeResource.InitRoutes(router)

	// Host
	hostResource := resource.HostResource{}
	hostResource.InitRoutes(router)

	// CPU
	cpuResource := resource.CPUResource{}
	cpuResource.InitRoutes(router)

	// GPU
	gpuResource := resource.GPUResource{}
	gpuResource.InitRoutes(router)

	// Memory
	memResource := resource.MemResource{}
	memResource.InitRoutes(router)

	// Disk
	diskResource := resource.DiskResource{}
	diskResource.InitRoutes(router)

	// Load
	loadResource := resource.LoadResource{}
	loadResource.InitRoutes(router)

	// Network
	netResource := resource.NetResource{}
	netResource.InitRoutes(router)

	return router
}
