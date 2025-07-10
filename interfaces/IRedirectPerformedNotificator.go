package interfaces

import (
	"context"

	"github.com/pedropazello/url-redirect-service/entities"
)

type IRedirectPerformedNotificator interface {
	Notificate(context.Context, entities.Redirect) error
}
