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

var ctx context.Context
var clientMock mocks.IDynamodbClient
var dynamoDB *db.DynamoDB
var getInput *dynamodb.GetItemInput
var putInput *dynamodb.PutItemInput

func setup(_ *testing.T) {
	ctx = context.Background()
	clientMock = mocks.IDynamodbClient{}
	dynamoDB = db.NewDynamoDB(&clientMock)
	getInput = &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: "1"},
		},
	}

	putInput = &dynamodb.PutItemInput{
		TableName: aws.String("Redirects"),
		Item: map[string]types.AttributeValue{
			"Id":            &types.AttributeValueMemberS{Value: "1"},
			"RedirectToURL": &types.AttributeValueMemberS{Value: "http://foo.com"},
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

	clientMock.EXPECT().GetItem(ctx, getInput).Return(output, nil)

	result, err := dynamoDB.GetItem(ctx, "1")

	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["Id"].(string))
	assert.Equal(t, "http://foo.com", result["RedirectToURL"].(string))
}

func TestGetItem_Error(t *testing.T) {
	setup(t)

	output := &dynamodb.GetItemOutput{}
	expectedErr := errors.New("connection error")

	clientMock.EXPECT().GetItem(ctx, getInput).Return(output, expectedErr)
	result, err := dynamoDB.GetItem(ctx, "1")

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, nil, result["Id"])
	assert.Equal(t, nil, result["RedirectToURL"])
}

func TestCreateItem_Success(t *testing.T) {
	setup(t)

	clientMock.EXPECT().PutItem(ctx, putInput).Return(&dynamodb.PutItemOutput{}, nil)

	dbInsert := map[string]any{
		"Id":            "1",
		"RedirectToURL": "http://foo.com",
	}

	result, err := dynamoDB.CreateItem(ctx, dbInsert)

	assert.Equal(t, nil, err)
	assert.Equal(t, "1", result["Id"])
	assert.Equal(t, "http://foo.com", result["RedirectToURL"])
}

func TestCreateItem_Error(t *testing.T) {
	setup(t)
	expectedErr := errors.New("connection error")

	clientMock.EXPECT().PutItem(ctx, putInput).Return(&dynamodb.PutItemOutput{}, expectedErr)

	dbInsert := map[string]any{
		"Id":            "1",
		"RedirectToURL": "http://foo.com",
	}

	_, err := dynamoDB.CreateItem(ctx, dbInsert)

	assert.Equal(t, expectedErr, err)
}
