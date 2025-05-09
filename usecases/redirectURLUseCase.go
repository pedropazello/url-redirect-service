package usecases

import (
	"context"

	"github.com/pedropazello/url-redirect-service/repositories"
)

func NewRedirectURLtUseCase() *RedirectURLUseCase {
	return &RedirectURLUseCase{}
}

type RedirectURLUseCase struct {
}

func (r RedirectURLUseCase) Execute(context context.Context, path string) (string, error) {
	redirectsRepository := repositories.NewRedirectsRepository()
	redirect, err := redirectsRepository.GetItem(context, path)

	return redirect.RedirectToURL, err
}
