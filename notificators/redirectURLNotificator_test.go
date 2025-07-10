package notificators_test

import (
	"context"
	"errors"
	"testing"

	"github.com/pedropazello/url-redirect-service/entities"
	"github.com/pedropazello/url-redirect-service/interfaces"
	"github.com/pedropazello/url-redirect-service/notificators"
	"github.com/stretchr/testify/assert"
)

var notificator interfaces.IRedirectPerformedNotificator
var redirect entities.Redirect

func Setup(_ *testing.T) {
	notificator = notificators.NewRedirectPerformedNotificator()
	redirect = entities.Redirect{Id: "1", RedirectToURL: "http://foo.com"}
}

func TestNotificate_Success(t *testing.T) {
	err := notificator.Notificate(context.Background(), redirect)

	assert.NoError(t, err)
}

func TestNotificate_Error(t *testing.T) {
	expectedErr := errors.New("failed to notificate")
	err := notificator.Notificate(context.Background(), redirect)

	assert.Equal(t, expectedErr, err)
}
