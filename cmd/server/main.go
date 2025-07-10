package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pedropazello/url-redirect-service/controllers"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/infra/topics"
	"github.com/pedropazello/url-redirect-service/notificators"
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

	snsClient := topics.NewSNSClient(ctx)
	topic := topics.NewSNSTopic(snsClient, "arn:aws:sns:us-east-1:000000000000:redirect_performed_topic")
	redirectPerformedNotificator := notificators.NewRedirectPerformedNotificator(topic)

	redirectUseCase := usecases.NewRedirectURLtUseCase(redirectRepository, redirectPerformedNotificator)
	redirectController := controllers.NewRedirectController(redirectUseCase)
	routes := routes.NewRoutes(redirectController)

	routes.RegisterRoutes(router)
	router.Run(":8080")
}
