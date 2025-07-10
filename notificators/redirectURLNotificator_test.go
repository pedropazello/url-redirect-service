package notificators_test

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/interfaces"
	"github.com/pedropazello/url-redirect-service/mocks"
	"github.com/pedropazello/url-redirect-service/notificators"
	"github.com/stretchr/testify/assert"
)

var notificator interfaces.IRedirectPerformedNotificator
var redirect entities.Redirect
var topicMock *mocks.ITopic
var redirectAsJSONString string

func setup(t *testing.T) {
	topicMock = mocks.NewITopic(t)
	notificator = notificators.NewRedirectPerformedNotificator(topicMock)

	redirect = entities.Redirect{Id: "1", RedirectToURL: "http://foo.com"}

	redirectAsBytes, _ := json.Marshal(redirect)
	redirectAsJSONString = string(redirectAsBytes)
}

func TestNotificate_Success(t *testing.T) {
	setup(t)

	topicMock.EXPECT().Publish(context.Background(), redirectAsJSONString).Return("1", nil)
	err := notificator.Notificate(context.Background(), redirect)

	assert.NoError(t, err)
}

func TestNotificate_Error(t *testing.T) {
	setup(t)
	expectedErr := errors.New("failed to notificate")

	topicMock.EXPECT().Publish(context.Background(), redirectAsJSONString).Return("", expectedErr)
	err := notificator.Notificate(context.Background(), redirect)

	assert.Equal(t, expectedErr, err)
}
