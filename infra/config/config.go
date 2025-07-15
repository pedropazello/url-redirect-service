package config

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
)

func LoadAWSConfig(ctx context.Context) (aws.Config, error) {
	if IsDevelopment() {
		return config.LoadDefaultConfig(
			ctx,
			config.WithRegion("us-east-1"),
			config.WithBaseEndpoint("http://localstack:4566"),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
		)
	}

	return config.LoadDefaultConfig(ctx)
}

func Environment() string {
	return os.Getenv("APP_ENV")
}

func IsProduction() bool {
	return Environment() == "production"
}

func IsDevelopment() bool {
	return Environment() == "development"
}
