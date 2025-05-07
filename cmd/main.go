package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/routes"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)
	router.Run(":8080")
}
