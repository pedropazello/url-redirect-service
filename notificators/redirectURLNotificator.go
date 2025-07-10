package notificators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

type redirectPerformedNotificator struct {
	topic interfaces.ITopic
}

func NewRedirectPerformedNotificator(topic interfaces.ITopic) interfaces.IRedirectPerformedNotificator {
	return &redirectPerformedNotificator{
		topic: topic,
	}
}

func (r *redirectPerformedNotificator) Notificate(ctx context.Context, redirect entities.Redirect) error {
	redirectAsJSON, _ := json.Marshal(redirect)

	msgID, err := r.topic.Publish(ctx, string(redirectAsJSON))

	fmt.Println(msgID)

	return err
}
