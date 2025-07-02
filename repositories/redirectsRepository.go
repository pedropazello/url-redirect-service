package repositories

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/interfaces"
)

const TableName string = "Redirects"

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
	var err error

	result, err := r.db.GetItem(context, TableName, Id)
	if err == nil {
		redirect, err = dbResultToRedirect(result)
	}

	return redirect, err
}

func (r RedirectsRepository) CreateItem(context context.Context, redirect entities.Redirect) (entities.Redirect, error) {
	insertDB, err := redirectToHash(redirect)
	if err != nil {
		return redirect, err
	}

	result, err := r.db.CreateItem(context, TableName, insertDB)
	if err == nil {
		redirect, err = dbResultToRedirect(result)
	}

	return redirect, err
}

func redirectToHash(redirect entities.Redirect) (map[string]any, error) {
	var insertDB map[string]any
	var err error

	data, err := json.Marshal(redirect)
	if err == nil {
		err = json.Unmarshal(data, &insertDB)
	}

	return insertDB, err
}

func dbResultToRedirect(result map[string]any) (entities.Redirect, error) {
	redirect := entities.Redirect{}
	var err error

	if val, ok := result["RedirectToURL"]; ok && val != nil {
		redirect.RedirectToURL = val.(string)
	} else {
		err = errors.New("[RedirectToURL] field not found")
	}

	if val, ok := result["Id"]; ok && val != nil {
		redirect.Id = val.(string)
	} else {
		err = errors.New("[Id] field not found")
	}

	return redirect, err
}
