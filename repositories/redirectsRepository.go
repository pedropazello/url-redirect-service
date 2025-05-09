package repositories

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/infra/db"
)

func NewRedirectsRepository() *RedirectsRepository {
	return &RedirectsRepository{}
}

type RedirectsRepository struct{}

func (r RedirectsRepository) GetItem(context context.Context, Id string) (entities.Redirect, error) {
	redirect := entities.Redirect{}

	input := &dynamodb.GetItemInput{
		TableName: aws.String("Redirects"),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberS{Value: Id},
		},
	}

	dynamoDB := db.NewDynamoDB()

	result, err := dynamoDB.GetItem(context, input)
	if err != nil {
		log.Fatalf("Failed to get item: %v", err)
	}

	if val, ok := result.Item["RedirectToURL"]; ok {
		redirect.RedirectToURL = val.(*types.AttributeValueMemberS).Value
	}

	if val, ok := result.Item["Id"]; ok {
		redirect.Id = val.(*types.AttributeValueMemberS).Value
	}

	return redirect, err
}
