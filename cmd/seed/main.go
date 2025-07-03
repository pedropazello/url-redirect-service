package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/infra/db"
	"github.com/pedropazello/url-redirect-service/repositories"
)

const totalItems = 1_000
const maxRetries = 4
const retryDelayMs = 200
const maxConcurrent = 1_000

var client *dynamodb.Client
var dynamoDB *db.DynamoDB
var redirectRepository *repositories.RedirectsRepository

func setup() context.Context {
	go monitorGoroutines()
	ctx := context.Background()
	client = db.NewDynamoDBClient(ctx)
	dynamoDB = db.NewDynamoDB(client)
	redirectRepository = repositories.NewRedirectsRepository(dynamoDB)

	return ctx
}

func runWithoutGoroutines() {
	ctx := setup()

	for i := 0; i < totalItems; i++ {
		newRedirect := createRandomRedirect()
		saveRedirect(ctx, newRedirect)
	}

	fmt.Println("All redirects inserted.")
}

func runWithGoroutines() {
	ctx := setup()

	redirectsChannel := make(chan entities.Redirect)
	var wg sync.WaitGroup

	for i := 0; i < maxConcurrent; i++ {
		wg.Add(1)
		go worker(ctx, i, redirectsChannel, &wg)
	}

	for i := 0; i < totalItems; i++ {
		redirect := createRandomRedirect()
		redirectsChannel <- redirect
	}

	close(redirectsChannel)
	wg.Wait()

	fmt.Println("All redirects inserted.")
}

func worker(ctx context.Context, workerID int, redirects <-chan entities.Redirect, wg *sync.WaitGroup) {
	defer wg.Done()
	for redirect := range redirects {
		success := false
		for attempt := 1; attempt <= maxRetries; attempt++ {
			fmt.Printf("Worker %d processing redirect %s\n", workerID, redirect)
			err := saveRedirect(ctx, redirect)

			if err == nil {
				success = true
				break
			}
			fmt.Printf("Worker %d: redirect %s failed on attempt %d: %v\n", workerID, redirect, attempt, err)
			time.Sleep(time.Millisecond * retryDelayMs)
		}
		if !success {
			fmt.Printf("Worker %d: redirect %s failed after %d attempts\n", workerID, redirect, maxRetries)
		}
	}
}

func monitorGoroutines() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		fmt.Printf("Active goroutines: %d\n", runtime.NumGoroutine())
	}
}

func createRandomRedirect() entities.Redirect {
	return entities.Redirect{
		Id:            uuid.New().String(),
		RedirectToURL: gofakeit.URL(),
	}
}

func saveRedirect(ctx context.Context, redirect entities.Redirect) error {
	resp, err := redirectRepository.CreateItem(ctx, redirect)
	if err == nil {
		fmt.Println("Inserted:", resp)
	}
	return err
}
