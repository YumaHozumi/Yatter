package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	CreateStatus(ctx context.Context, status *object.Status, account *object.Account) (int64, error)
	FindStatusByID(ctx context.Context, id int64) (*object.Status, error)
}
