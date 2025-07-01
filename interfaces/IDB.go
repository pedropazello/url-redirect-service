package interfaces

import "context"

type IDB interface {
	GetItem(context context.Context, Id string) (map[string]any, error)
	CreateItem(context context.Context, insertion map[string]any) (map[string]any, error)
}
