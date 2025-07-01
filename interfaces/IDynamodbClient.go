package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type IDynamodbClient interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
}
