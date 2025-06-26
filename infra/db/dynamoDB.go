package db

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func NewDynamoDB() *DynamoDB {
	return &DynamoDB{}
}

type DynamoDB struct {
}

func (d DynamoDB) GetItem(context context.Context, Id string) (map[string]any, error) {
	var result map[string]any

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: Id},
		},
	}

	client := createClient(context)
	output, err := client.GetItem(context, input)
	if err != nil {
		return result, err
	}

	dynamoItem := output.Item
	err = attributevalue.UnmarshalMap(dynamoItem, &result)

	return result, err
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
