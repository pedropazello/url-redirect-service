package db_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/go-playground/assert/v2"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/mocks"
)

func TestGetItem_Success(t *testing.T) {
	clientMock := mocks.IDynamodbClient{}
	dynamoDB := db.NewDynamoDB(&clientMock)

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: "1"},
		},
	}

	output := &dynamodb.GetItemOutput{
		Item: map[string]types.AttributeValue{
			"Id":            &types.AttributeValueMemberS{Value: "1"},
			"RedirectToURL": &types.AttributeValueMemberS{Value: "http://foo.com"},
		},
	}

	clientMock.EXPECT().GetItem(context.Background(), input).Return(output, nil)

	result, err := dynamoDB.GetItem(context.Background(), "1")

	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["Id"].(string))
	assert.Equal(t, "http://foo.com", result["RedirectToURL"].(string))
}
