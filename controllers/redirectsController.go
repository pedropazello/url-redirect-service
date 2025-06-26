package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

type RedirectController struct {
	usecase interfaces.IRedirectUsecase
}

func NewRedirectController(usecase interfaces.IRedirectUsecase) *RedirectController {
	return &RedirectController{
		usecase: usecase,
	}
}

func (rc *RedirectController) GetRedirects(c *gin.Context) {
	path := c.Param("path")

	redirectToURL, err := rc.usecase.Execute(c.Request.Context(), path)

	if err == nil {
		fmt.Printf("Retrieved name: %v\n", redirectToURL)
		c.Redirect(http.StatusFound, redirectToURL)
		return
	}

	c.String(http.StatusNotFound, "No redirection found for path: %s", path)
}
