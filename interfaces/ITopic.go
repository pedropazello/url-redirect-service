package interfaces

import "context"

type ITopic interface {
	Publish(ctx context.Context, msg string) (string, error)
}
