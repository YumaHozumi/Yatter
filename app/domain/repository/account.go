package repository

import (
	"context"

	"yatter-backend-go/app/domain/object"
)

type Account interface {
	// Fetch account which has specified username
	FindByUsername(ctx context.Context, username string) (*object.Account, error)
	// TODO: Add Other APIs
	CreateUser(ctx context.Context, account *object.Account) error
	FindByUserID(ctx context.Context, id int64) (*object.Account, error)
}
