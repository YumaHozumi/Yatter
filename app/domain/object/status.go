package object

import "time"

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	AccountID int64 `json:"account_id" db:"account_id"`

	// The username of the account
	Username string `json:"username,omitempty"`

	// The content of the status
	Content string `json:"content,omitempty"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"create_at"`
}
