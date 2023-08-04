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
	timeline struct {
		db *sqlx.DB
	}
)

func NewTimeline(db *sqlx.DB) repository.Timeline {
	return &timeline{db: db}
}

func (r *timeline) GetPublic(ctx context.Context) (*object.Timeline, error) {
	entity := new(object.Timeline)
	//status一覧を取得するクエリ
	query := `
		SELECT 
		status.id AS "id", 
		accounts.id AS "account.id",
		accounts.username AS "account.username",
		accounts.create_at AS "account.created_at",
		status.content AS "content",
		status.created_at AS "create_at"
		FROM 
			status
		INNER JOIN 
		accounts ON status.account_id = accounts.id
	`
	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to get timeline from db: %w", err)
	}

	return entity, nil
}
