package store

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/model"
)

type UserStore interface {
	CheckIfEmailIsAlreadyInUse(ctx context.Context, email string) (err error, ok bool)
	CreateUser(background context.Context, username string, password string) (err error, user model.User, ok bool)
}
