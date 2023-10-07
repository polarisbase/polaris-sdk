package store

import (
	"context"
	"github.com/polarisbase/polaris-sdk/v2/authentication/internal/users/model"
)

type UserStore interface {
	CheckIfEmailIsAlreadyInUse(ctx context.Context, email string) (err error, ok bool)
	CreateUser(background context.Context, email string, passwordHash string, profileData map[string]string) (err error, user model.User, ok bool)
	GetUsers(ctx context.Context, limit int, offset int) (err error, users []model.User, ok bool)
	FindUserByEmail(ctx context.Context, email string) (err error, user model.User, ok bool)
}
