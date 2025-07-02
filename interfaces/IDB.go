package interfaces

import "context"

type IDB interface {
	GetItem(context context.Context, tableName string, Id string) (map[string]any, error)
	CreateItem(context context.Context, tableName string, insertion map[string]any) (map[string]any, error)
}
