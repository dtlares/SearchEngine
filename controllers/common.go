package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFound route
// @Summary Redirection when resource is not found
// @Produce json
func NotFound(c *gin.Context) {
	c.Status(http.StatusBadRequest)
}

// NoMethods resource no methods
func NoMethods(c *gin.Context) {
	c.Status(http.StatusBadRequest)
}
