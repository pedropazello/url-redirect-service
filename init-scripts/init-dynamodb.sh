#!/bin/bash

echo "Initializing DynamoDB table..."

awslocal dynamodb create-table \
  --table-name Redirects \
  --attribute-definitions AttributeName=Id,AttributeType=S \
  --key-schema AttributeName=Id,KeyType=HASH \
  --provisioned-throughput ReadCapacityUnits=5,WriteCapacityUnits=5

echo "Table Redirects created ."


awslocal dynamodb put-item \
    --table-name Redirects \
    --item '{
        "Id": {"S": "1"},
        "RedirectToURL": {"S": "https://example.com"}
    }'

awslocal dynamodb put-item \
    --table-name Redirects \
    --item '{
        "Id": {"S": "2"},
        "RedirectToURL": {"S": "https://google.com"}
    }'

echo "Table Redirects initialized with sample data."