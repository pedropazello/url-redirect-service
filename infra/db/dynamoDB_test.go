package db_test

import (
	"context"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/go-playground/assert/v2"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/mocks"
)

var clientMock mocks.IDynamodbClient
var dynamoDB *db.DynamoDB
var input *dynamodb.GetItemInput

func setup(_ *testing.T) {
	clientMock = mocks.IDynamodbClient{}
	dynamoDB = db.NewDynamoDB(&clientMock)
	input = &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: "1"},
		},
	}
}

func TestGetItem_Success(t *testing.T) {
	setup(t)

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

func TestGetItem_Error(t *testing.T) {
	setup(t)

	output := &dynamodb.GetItemOutput{}
	expectedErr := errors.New("connection error")

	clientMock.EXPECT().GetItem(context.Background(), input).Return(output, expectedErr)
	result, err := dynamoDB.GetItem(context.Background(), "1")

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, nil, result["Id"])
	assert.Equal(t, nil, result["RedirectToURL"])
}
