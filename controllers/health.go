package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Status healthcheck
// @Summary Provides the current api status
// @Produce json
// @Success 200 {object}
// @Router /status [get]
func Status(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
