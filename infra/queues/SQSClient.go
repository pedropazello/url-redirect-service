package queues

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/pedropazello/url-redirect-service/infra/config"
)

func NewSQSClient(ctx context.Context) *sqs.Client {
	cfg, err := config.LoadAWSConfig(ctx)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return sqs.NewFromConfig(cfg)
}
