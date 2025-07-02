package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/controllers"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/repositories"
	"github.com/pedropazello/url-redirect-service/routes"
	"github.com/pedropazello/url-redirect-service/usecases"
)

func main() {
	ctx := context.Background()
	router := gin.Default()

	dynamoClient := db.NewDynamoDBClient(ctx)
	db := db.NewDynamoDB(dynamoClient)
	redirectRepository := repositories.NewRedirectsRepository(db)
	redirectUseCase := usecases.NewRedirectURLtUseCase(redirectRepository)
	redirectController := controllers.NewRedirectController(redirectUseCase)
	routes := routes.NewRoutes(redirectController)

	routes.RegisterRoutes(router)
	router.Run(":8080")
}
