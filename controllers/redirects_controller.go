package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetRedirects(c *gin.Context) {
	path := c.Param("path")

	cfg, err := config.LoadDefaultConfig(
		c.Request.Context(),
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint("http://localstack:4566"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: path},
		},
	}

	result, err := client.GetItem(c.Request.Context(), input)
	if err != nil {
		log.Fatalf("Failed to get item: %v", err)
	}

	if val, ok := result.Item["RedirectToURL"]; ok {
		redirectToURL := val.(*types.AttributeValueMemberS).Value

		fmt.Printf("Retrieved name: %v\n", redirectToURL)
		c.Redirect(http.StatusFound, redirectToURL)
		return
	}

	c.String(http.StatusNotFound, "No redirection found for path: %s", path)
}
