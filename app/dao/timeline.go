package dao

import (
	"context"
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

func (r *timeline) GetPublic(ctx context.Context, limit int, maxID int64, sinceID int64) (*object.Timeline, error) {
	entity := new(object.Timeline)
	//status一覧を取得するクエリ
	query := `
		SELECT 
		status.id AS "id", 
		account.id AS "account.id",
		account.username AS "account.username",
		account.create_at AS "account.create_at",
		status.content AS "content",
		status.created_at AS "created_at"
		FROM 
			status
		INNER JOIN 
		account ON status.account_id = account.id
		`
	//maxID未満のアカウントのstatusだけ
	if maxID > 0 {
		query += fmt.Sprintf("WHERE account.id < %d ", maxID)
	}
	//sinceIDより大きいアカウントのstatusだけ
	if sinceID > 0 {
		query += fmt.Sprintf("WHERE account.id > %d ", sinceID)
	}
	//取得する件数を制限
	query += fmt.Sprintf("LIMIT %d ", limit)

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	//statusをtimelineに追加
	for rows.Next() {
		status := new(object.Status)
		if err := rows.StructScan(&status); err != nil {
			return nil, err
		}
		entity.Statuses = append(entity.Statuses, *status)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return entity, nil
}
