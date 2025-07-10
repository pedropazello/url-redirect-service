package topics

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type snsTopic struct {
	client    *sns.Client
	topicName string
}

func NewSNSTopic(client *sns.Client, topicName string) *snsTopic {
	return &snsTopic{
		client:    client,
		topicName: topicName,
	}
}

func (t *snsTopic) Publish(ctx context.Context, msg string) (string, error) {
	out, err := t.client.Publish(ctx, &sns.PublishInput{
		Message:  aws.String(msg),
		TopicArn: &t.topicName,
	})

	if err != nil {
		return "", err
	}

	return *out.MessageId, nil
}
