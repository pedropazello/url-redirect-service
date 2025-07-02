package repositories_test

import (
	"context"
	"errors"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/mocks"
	"github.com/pedropazello/url-redirect-service/repositories"
)

var mockDB *mocks.IDB
var redirectRepository *repositories.RedirectsRepository
var tableName string

func setup(t *testing.T) {
	mockDB = mocks.NewIDB(t)
	redirectRepository = repositories.NewRedirectsRepository(mockDB)
	tableName = "Redirects"
}

func TestGetItem_Success(t *testing.T) {
	setup(t)

	dbResult := map[string]any{
		"Id":            "1",
		"RedirectToURL": "http://foo.com",
	}

	mockDB.EXPECT().GetItem(context.Background(), tableName, "1").Return(dbResult, nil)
	resp, err := redirectRepository.GetItem(context.Background(), "1")

	assert.Equal(t, nil, err)
	assert.Equal(t, "http://foo.com", resp.RedirectToURL)
}

func TestGetItem_Error(t *testing.T) {
	setup(t)

	var dbResult map[string]any
	expectedErr := errors.New("Failed to get item: not found")

	mockDB.EXPECT().GetItem(context.Background(), tableName, "1").Return(dbResult, expectedErr)
	resp, err := redirectRepository.GetItem(context.Background(), "1")

	assert.Equal(t, expectedErr, err)
	assert.Equal(t, "", resp.RedirectToURL)
}

func TestGetItem_DifferentDocumentFormat(t *testing.T) {
	setup(t)

	dbResult := map[string]any{
		"Id": "1",
	}

	mockDB.EXPECT().GetItem(context.Background(), tableName, "1").Return(dbResult, nil)
	resp, err := redirectRepository.GetItem(context.Background(), "1")

	assert.Equal(t, errors.New("[RedirectToURL] field not found"), err)
	assert.Equal(t, "", resp.RedirectToURL)
}

func TestCreateItem_Success(t *testing.T) {
	setup(t)

	redirect := entities.Redirect{RedirectToURL: "http://foo.com"}

	dbInsert := map[string]any{
		"Id":            "",
		"RedirectToURL": "http://foo.com",
	}

	dbResult := map[string]any{
		"Id":            "1",
		"RedirectToURL": "http://foo.com",
	}

	mockDB.EXPECT().CreateItem(context.Background(), tableName, dbInsert).Return(dbResult, nil)
	resp, err := redirectRepository.CreateItem(context.Background(), redirect)

	assert.Equal(t, nil, err)
	assert.Equal(t, "1", resp.Id)
	assert.Equal(t, "http://foo.com", resp.RedirectToURL)
}
