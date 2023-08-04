package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetPublic(ctx context.Context) (*object.Timeline, error)
}
