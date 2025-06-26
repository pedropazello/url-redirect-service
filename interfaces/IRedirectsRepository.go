package interfaces

import (
	"context"

	"github.com/pedropazello/url-redirect-service/entities"
)

type IRedirectsRepository interface {
	GetItem(context context.Context, Id string) (entities.Redirect, error)
}
