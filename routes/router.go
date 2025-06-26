package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/controllers"
)

type Routes struct {
	redirectController *controllers.RedirectController
}

func NewRoutes(redirectController *controllers.RedirectController) *Routes {
	return &Routes{
		redirectController: redirectController,
	}
}

func (r Routes) RegisterRoutes(engine *gin.Engine) {
	redirectsGroup := engine.Group("/redirects")

	{
		redirectsGroup.GET("/:path", r.redirectController.GetRedirects)
	}

	engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

}
