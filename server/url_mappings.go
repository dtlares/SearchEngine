package server

import (
	"SearchEngine/controllers"

	"github.com/gin-gonic/gin"
)

// URL mapping declarations
func mapUrlsToControllers(router *gin.Engine) {

	// Health check
	router.GET("/status", controllers.Status)

	// Api Routes
	inputJSON := router.Group("/")
	inputJSON.GET("/search", controllers.Search)

	// System routes
	router.NoRoute(controllers.NotFound)
	router.NoMethod(controllers.NoMethods)

	// Debugging
	router.StaticFile("/profiling", "./profiling")
}
