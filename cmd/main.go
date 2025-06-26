package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/controllers"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/repositories"
	"github.com/pedropazello/url-redirect-service/routes"
	"github.com/pedropazello/url-redirect-service/usecases"
)

func main() {
	router := gin.Default()

	db := db.NewDynamoDB()
	redirectRepository := repositories.NewRedirectsRepository(db)
	redirectUseCase := usecases.NewRedirectURLtUseCase(redirectRepository)
	redirectController := controllers.NewRedirectController(redirectUseCase)
	routes := routes.NewRoutes(redirectController)

	routes.RegisterRoutes(router)
	router.Run(":8080")
}
