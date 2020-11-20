package server

import (
	"SearchEngine/config"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

// New gin server
func new() *gin.Engine {
	router := gin.New()

	// Add loging
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Add middlewares

	// All other routes
	mapUrlsToControllers(router)

	return router
}

// StartApp starts the application
func StartApp() {
	// Start endpoint
	router = new()
	router.Run(config.PortNumber)
}
