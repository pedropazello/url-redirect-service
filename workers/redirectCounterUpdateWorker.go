package workers

import (
	"fmt"

	"github.com/pedropazello/url-redirect-service/entities"
)

func RedirectCounterUpdateWorkerPerform(redirects chan entities.Redirect) {
	fmt.Println("running redirectPerformedCounterUpdateWorker...")

	for redirect := range redirects {
		fmt.Printf("updating counter for: %s\n", redirect)
	}
}
