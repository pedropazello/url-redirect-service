package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	router.GET("/redirect/:path", func(c *gin.Context) {
		path := c.Param("path")

		if path == "demo" {
			c.Redirect(http.StatusFound, "https://example.com")
			return
		}

		c.String(http.StatusNotFound, "No redirection found for path: %s", path)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	serverAddr := ":" + port
	fmt.Printf("Starting server on %s\n", serverAddr)
	if err := router.Run(serverAddr); err != nil {
		fmt.Printf("Server failed to start: %v\n", err)
	}
}
