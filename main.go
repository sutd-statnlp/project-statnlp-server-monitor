package main

import (
	"github.com/gin-gonic/gin"

	"./controller"
	"./resource"
)

func main() {
	router := setupRoutes()

	router.Run(":8080")
}

func setupRoutes() *gin.Engine {
	router := gin.Default()

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
