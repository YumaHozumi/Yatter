package dao

import (
	"context"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	status struct {
		db *sqlx.DB
	}
)

func NewStatus(db *sqlx.DB) repository.Status {
	return &status{db: db}
}

// ステータスを追加
func (r *status) AddStatus(ctx context.Context, status *object.Status) error {
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

	panic("stop")

	if _, err = tx.Exec(`INSERT INTO account (id, username, password_hash) VALUES (?, ?, ?)`, account.ID, account.Username, account.PasswordHash); err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}
