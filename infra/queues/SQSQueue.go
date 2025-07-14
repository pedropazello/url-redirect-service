package queues

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type SQSQueue struct {
	client   *sqs.Client
	queueURL string
}

func NewSQSQueue(client *sqs.Client, queueURL string) *SQSQueue {
	return &SQSQueue{
		client:   client,
		queueURL: queueURL,
	}
}

func (q *SQSQueue) ReceiveMessage(ctx context.Context) (*sqs.ReceiveMessageOutput, error) {
	return q.client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(q.queueURL),
		MaxNumberOfMessages: 5,
		WaitTimeSeconds:     10,
	})
}

func (q *SQSQueue) DeleteMessage(ctx context.Context, receiptHandle *string) error {
	_, err := q.client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(q.queueURL),
		ReceiptHandle: receiptHandle,
	})

	return err
}
