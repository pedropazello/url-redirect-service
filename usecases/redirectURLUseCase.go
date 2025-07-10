package usecases

import (
	"context"
	"fmt"

	"github.com/pedropazello/url-redirect-service/interfaces"
)

func NewRedirectURLtUseCase(redirectsRepository interfaces.IRedirectsRepository,
	redirectPerformedNotificator interfaces.IRedirectPerformedNotificator) *RedirectURLUseCase {
	return &RedirectURLUseCase{
		redirectsRepository:          redirectsRepository,
		redirectPerformedNotificator: redirectPerformedNotificator,
	}
}

type RedirectURLUseCase struct {
	redirectsRepository          interfaces.IRedirectsRepository
	redirectPerformedNotificator interfaces.IRedirectPerformedNotificator
}

func (r RedirectURLUseCase) Execute(ctx context.Context, path string) (string, error) {
	redirect, err := r.redirectsRepository.GetItem(ctx, path)
	if err != nil {
		return "", err
	}

	err = r.redirectPerformedNotificator.Notificate(ctx, redirect)
	if err != nil {
		fmt.Printf("notificator.Notificate failed: %s", err)
	}

	return redirect.RedirectToURL, nil
}
