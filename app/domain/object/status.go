package object

import "time"

type Status struct {
	// The internal ID of the status
	ID int64 `json:"id,omitempty"`

	AccountID int64 `json:"-" db:"account_id"`

	Account Account `json:"account"`

	// The content of the status
	Content string `json:"content,omitempty"`

	URL *string `json:"url,omitempty" db:"url"`

	// The time the account was created
	CreateAt time.Time `json:"create_at,omitempty" db:"created_at"`
}
