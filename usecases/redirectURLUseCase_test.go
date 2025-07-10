// go
package usecases_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/mocks"
	"github.com/pedropazello/url-redirect-service/usecases"
	"github.com/stretchr/testify/assert"
)

var mockRepo *mocks.IRedirectsRepository
var mockNotificator *mocks.IRedirectPerformedNotificator
var useCase *usecases.RedirectURLUseCase

func setup(t *testing.T) {
	mockRepo = mocks.NewIRedirectsRepository(t)
	mockNotificator = mocks.NewIRedirectPerformedNotificator(t)
	useCase = usecases.NewRedirectURLtUseCase(mockRepo, mockNotificator)
}

func TestRedirectURLUseCase_Execute_Success(t *testing.T) {
	setup(t)

	expectedURL := "https://example.com"
	expectedRedirect := entities.Redirect{RedirectToURL: expectedURL}
	mockRepo.EXPECT().GetItem(context.Background(), "/foo").Return(expectedRedirect, nil)
	mockNotificator.EXPECT().Notificate(context.Background(), expectedRedirect).Return(nil)

	url, err := useCase.Execute(context.Background(), "/foo")
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url)
}

func TestRedirectURLUseCase_Execute_Error(t *testing.T) {
	setup(t)

	err := errors.New("not found")
	mockRepo.EXPECT().GetItem(context.Background(), "/bar").Return(entities.Redirect{}, err)

	url, err := useCase.Execute(context.Background(), "/bar")
	assert.Error(t, err)
	assert.Empty(t, url)
}

func TestRedirectURLUseCase_Execute_Notificator_Failed(t *testing.T) {
	setup(t)

	err := errors.New("notification failed")

	expectedURL := "https://example.com"
	expectedRedirect := entities.Redirect{RedirectToURL: expectedURL}
	mockRepo.EXPECT().GetItem(context.Background(), "/foo").Return(expectedRedirect, nil)
	mockNotificator.EXPECT().Notificate(context.Background(), expectedRedirect).Return(err)

	url, err := useCase.Execute(context.Background(), "/foo")
	assert.NoError(t, err)
	assert.Equal(t, expectedURL, url)
}
