package object

type Media struct {
	MediaID int64 `json:"media_id,omitempty" db:"id"`

	MediaURL string `json:"media_url,omitempty" db:"media_url"`
}
