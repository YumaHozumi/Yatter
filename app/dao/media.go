package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	media struct {
		db *sqlx.DB
	}
)

func NewMedia(db *sqlx.DB) repository.Media {
	return &media{db: db}
}

func (r *media) UploadMedia(ctx context.Context, media *object.Media) (int64, error) {
	//トランザクションの開始
	tx, _ := r.db.Begin()

	var err error
	defer func() {
		switch r := recover(); {
		case r != nil:
			tx.Rollback()
			panic(r)
		case err != nil:
			tx.Rollback()
		}
	}()

	res, err := tx.Exec(`INSERT INTO media (media_url) VALUES (?)`, media.MediaURL)
	if err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *media) FindMedia(ctx context.Context, id int64) (*object.Media, error) {
	entity := new(object.Media)

	err := r.db.QueryRowxContext(ctx, "select * from media where id = ?", id).StructScan(entity)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find media from db: %w", err)
	}

	return entity, nil
}
