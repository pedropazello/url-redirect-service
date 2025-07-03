package db

import (
	"context"
	"fmt"

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

func (d DynamoDB) GetItem(context context.Context, tableName string, Id string) (map[string]any, error) {
	var result map[string]any

	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
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

func (d DynamoDB) CreateItem(context context.Context, tableName string, insertion map[string]any) (map[string]any, error) {
	itens := make(map[string]types.AttributeValue)

	for k, v := range insertion {
		switch val := v.(type) {
		case string:
			itens[k] = &types.AttributeValueMemberS{Value: val}
		default:
			return insertion, fmt.Errorf("unsupported type for key %s: %T", k, v)
		}
	}

	input := &dynamodb.PutItemInput{
		TableName:           aws.String(tableName),
		Item:                itens,
		ConditionExpression: aws.String("attribute_not_exists(Id)"),
	}

	_, err := d.client.PutItem(context, input)

	return insertion, err
}
