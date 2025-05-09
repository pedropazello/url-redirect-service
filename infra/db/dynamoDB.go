package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDB() *DynamoDB {
	return &DynamoDB{}
}

type DynamoDB struct {
}

func (d DynamoDB) GetItem(context context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	client := createClient(context)
	output, err := client.GetItem(context, input)

	return output, err
}

func createClient(context context.Context) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(
		context,
		config.WithRegion("us-east-1"),
		config.WithBaseEndpoint("http://localstack:4566"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return dynamodb.NewFromConfig(cfg)
}
