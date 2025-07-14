package workers

import (
	"fmt"

	"github.com/pedropazello/url-redirect-service/entities"
)

func SendRedirectMetricsWorkerPerform(redirects chan entities.Redirect) {
	fmt.Println("running redirectPerformedMetricsWorker...")

	for redirect := range redirects {
		fmt.Printf("saving metrics for: %s\n", redirect)
	}
}
