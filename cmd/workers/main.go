package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/infra/config"
	"github.com/pedropazello/url-redirect-service/infra/queues"
	"github.com/pedropazello/url-redirect-service/workers"
)

func main() {
	ctx := context.Background()

	cfg, err := config.LoadAWSConfig(ctx)
	if err != nil {
		fmt.Println(err)
	}

	client := sqs.NewFromConfig(cfg)

	redirectPerformedCounterUpdateChannel := make(chan entities.Redirect)
	redirectPerformedMetricsChannel := make(chan entities.Redirect)

	for {
		go workers.RedirectCounterUpdateWorkerPerform(redirectPerformedCounterUpdateChannel)
		go workers.SendRedirectMetricsWorkerPerform(redirectPerformedMetricsChannel)

		go poolMessagesFor(ctx, client, redirectPerformedCounterUpdateChannel, "http://sqs.us-east-1.localhost:4566/000000000000/redirect_performed_counter_update_queue")
		go poolMessagesFor(ctx, client, redirectPerformedMetricsChannel, "http://sqs.us-east-1.localhost:4566/000000000000/redirect_performed_metrics_queue")

		time.Sleep(2 * time.Second)
	}
}

func poolMessagesFor(ctx context.Context, client *sqs.Client, redirects chan<- entities.Redirect, queueURL string) {
	sqsQueue := queues.NewSQSQueue(client, queueURL)

	// Receive messages
	resp, err := sqsQueue.ReceiveMessage(ctx)
	if err != nil {
		log.Printf("From (%s) failed to receive messages: %v", queueURL, err)
	}

	if len(resp.Messages) == 0 {
		fmt.Printf("From (%s), No messages...\n", queueURL)
	}

	for _, msg := range resp.Messages {
		fmt.Printf("From (%s) Message: %s\n", queueURL, aws.ToString(msg.Body))

		var outer map[string]any
		if err := json.Unmarshal([]byte(*msg.Body), &outer); err != nil {
			log.Fatalf("Failed to parse outer JSON: %v", err)
		}

		msgStr, ok := outer["Message"].(string)
		if !ok {
			log.Fatal("Message field not found or is not a string")
		}

		redirect := entities.Redirect{}
		b := []byte(msgStr)
		json.Unmarshal(b, &redirect)

		redirects <- redirect

		// Delete the message after processing
		err := sqsQueue.DeleteMessage(ctx, msg.ReceiptHandle)
		if err != nil {
			log.Printf("From (%s) failed to delete message: %v\n", queueURL, err)
		}
	}
}
