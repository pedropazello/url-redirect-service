package usecases

import (
	"context"

	"github.com/pedropazello/url-redirect-service/interfaces"
)

func NewRedirectURLtUseCase(redirectsRepository interfaces.IRedirectsRepository) *RedirectURLUseCase {
	return &RedirectURLUseCase{
		redirectsRepository: redirectsRepository,
	}
}

type RedirectURLUseCase struct {
	redirectsRepository interfaces.IRedirectsRepository
}

func (r RedirectURLUseCase) Execute(context context.Context, path string) (string, error) {
	redirect, err := r.redirectsRepository.GetItem(context, path)

	return redirect.RedirectToURL, err
}
