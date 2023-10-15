package store

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v2/authn/internal/sessions/model"
)

type SessionStore interface {
	CreateSession(context context.Context, session model.Session) (err error, sessionOut model.Session, ok bool)
}
