package repositories

import (
	"context"
	"log"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

func NewRedirectsRepository(db interfaces.IDB) *RedirectsRepository {
	return &RedirectsRepository{
		db: db,
	}
}

type RedirectsRepository struct {
	db interfaces.IDB
}

func (r RedirectsRepository) GetItem(context context.Context, Id string) (entities.Redirect, error) {
	redirect := entities.Redirect{}

	result, err := r.db.GetItem(context, Id)
	if err != nil {
		log.Fatalf("Failed to get item: %v", err)
	}

	redirect.RedirectToURL = result["RedirectToURL"].(string)
	redirect.Id = result["Id"].(string)

	return redirect, err
}
