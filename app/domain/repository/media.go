package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Media interface {
	UploadMedia(ctx context.Context, media *object.Media) (int64, error)
	FindMedia(ctx context.Context, id int64) (*object.Media, error)
}
