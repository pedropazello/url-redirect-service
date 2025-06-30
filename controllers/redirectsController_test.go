// go
package controllers_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"github.com/pedropazello/url-redirect-service/controllers"
	"github.com/pedropazello/url-redirect-service/mocks"
	"github.com/pedropazello/url-redirect-service/testutils"
	"github.com/stretchr/testify/assert"
)

var getRedirectsPath string
var mockUsecase *mocks.IRedirectUsecase
var controller *controllers.RedirectController

func setup(t *testing.T) {
	getRedirectsPath = "/redirects/:path"
	mockUsecase = mocks.NewIRedirectUsecase(t)
	controller = controllers.NewRedirectController(mockUsecase)
}

func TestGetRedirects_Success(t *testing.T) {
	setup(t)

	mockUsecase.EXPECT().Execute(context.Background(), "1").Return("https://example.com", nil)
	w := testutils.MakeGetRequest(getRedirectsPath, controller.GetRedirects, "/redirects/1")

	assert.Equal(t, http.StatusFound, w.Code)
	assert.Equal(t, "https://example.com", w.Header().Get("Location"))
}

func TestGetRedirects_NotFound(t *testing.T) {
	setup(t)

	err := errors.New("not found")
	mockUsecase.EXPECT().Execute(context.Background(), "1").Return("", err)

	w := testutils.MakeGetRequest("/redirects/:path", controller.GetRedirects, "/redirects/1")

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "No redirection found for path: 1")
}
