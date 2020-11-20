package controllers

import (
	"SearchEngine/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Search searches on Internet
// @Summary Returns search results on Internet
// @Params GET /search?engine=google&engine=bing&query=<phrase_to_search>
// @Produce json
// @Success 200 {object}
// @Router /search [get]
func Search(c *gin.Context) {
	queryParameters := c.Request.URL.Query()
	engines := queryParameters["engine"]
	query := queryParameters["query"]

	value := services.Search(query[0], engines)
	fmt.Println("Value: ", value)
	c.JSON(http.StatusOK, value)
}
