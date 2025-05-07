package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/controllers"
)

func RegisterRoutes(r *gin.Engine) {
	redirectsGroup := r.Group("/redirects")
	{
		redirectsGroup.GET("/:path", controllers.GetRedirects)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

}
