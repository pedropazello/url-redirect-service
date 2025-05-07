package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRedirects(c *gin.Context) {
	path := c.Param("path")

	if path == "demo" {
		c.Redirect(http.StatusFound, "https://example.com")
		return
	}

	c.String(http.StatusNotFound, "No redirection found for path: %s", path)
}
