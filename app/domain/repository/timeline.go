package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Timeline interface {
	GetPublic(ctx context.Context, limit int, maxID int64, sinceID int64) (*object.Timeline, error)
}
