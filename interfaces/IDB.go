package interfaces

import "context"

type IDB interface {
	GetItem(context context.Context, Id string) (map[string]any, error)
}
