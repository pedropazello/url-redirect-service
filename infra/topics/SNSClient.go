package topics

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/pedropazello/url-redirect-service/infra/config"
)

func NewSNSClient(ctx context.Context) *sns.Client {
	cfg, err := config.LoadAWSConfig(ctx)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return sns.NewFromConfig(cfg)
}
