#!/bin/bash

echo "Initializing SNS topic..."

awslocal sns create-topic --name redirect_performed_topic

awslocal sqs create-queue --queue-name redirect_performed_counter_update_queue
awslocal sqs create-queue --queue-name redirect_performed_metrics_queue

echo "SNS topic and SQS queues created ."

awslocal sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:000000000000:redirect_performed_topic \
  --protocol sqs \
  --notification-endpoint arn:aws:sqs:us-east-1:000000000000:redirect_performed_counter_update_queue

awslocal sns subscribe \
  --topic-arn arn:aws:sns:us-east-1:000000000000:redirect_performed_topic \
  --protocol sqs \
  --notification-endpoint arn:aws:sqs:us-east-1:000000000000:redirect_performed_metrics_queue

echo "SNS subscriptions created ."