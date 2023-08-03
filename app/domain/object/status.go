package object

import "time"

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	AccountID int64 `json:"account_id" db:"account_id"`

	// The content of the status
	Content string `json:"content,omitempty"`

	URL *string `json:"url,omitempty" db:"url"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"created_at"`
}
