package interfaces

import "context"

type IRedirectUsecase interface {
	Execute(context context.Context, path string) (string, error)
}
