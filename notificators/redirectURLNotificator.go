package notificators

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/infra/topics"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

type redirectPerformedNotificator struct {
}

func NewRedirectPerformedNotificator() interfaces.IRedirectPerformedNotificator {
	return &redirectPerformedNotificator{}
}

func (r *redirectPerformedNotificator) Notificate(ctx context.Context, redirect entities.Redirect) error {
	redirectAsJSON, _ := json.Marshal(redirect)

	client := topics.NewSNSClient(ctx)
	topic := topics.NewSNSTopic(client, "arn:aws:sns:us-east-1:000000000000:redirect_performed_topic")
	msgID, err := topic.Publish(ctx, string(redirectAsJSON))

	fmt.Println(msgID)

	return err
}
