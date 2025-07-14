package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pedropazello/url-redirect-service/infra/config"
)

func NewDynamoDBClient(ctx context.Context) *dynamodb.Client {
	cfg, err := config.LoadAWSConfig(ctx)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return dynamodb.NewFromConfig(cfg)
}
