package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/usecases"
)

func GetRedirects(c *gin.Context) {
	path := c.Param("path")

	redirectURLUseCase := usecases.NewRedirectURLtUseCase()
	redirectToURL, err := redirectURLUseCase.Execute(c.Request.Context(), path)

	if err == nil {
		fmt.Printf("Retrieved name: %v\n", redirectToURL)
		c.Redirect(http.StatusFound, redirectToURL)
		return
	}

	c.String(http.StatusNotFound, "No redirection found for path: %s", path)
}
