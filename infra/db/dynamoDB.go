package db

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

func NewDynamoDB(client interfaces.IDynamodbClient) *DynamoDB {
	return &DynamoDB{
		client: client,
	}
}

type DynamoDB struct {
	client interfaces.IDynamodbClient
}

func (d DynamoDB) GetItem(context context.Context, Id string) (map[string]any, error) {
	var result map[string]any

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: Id},
		},
	}

	output, err := d.client.GetItem(context, input)
	if err != nil {
		return result, err
	}

	dynamoItem := output.Item
	err = attributevalue.UnmarshalMap(dynamoItem, &result)

	return result, err
}
