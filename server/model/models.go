package model

import "time"

type URL struct {
	// ID        string    `json:"id,omitempty"`
	LongURL   string    `json:"longUrl"`
	URLHash   string    `json:"urlHash"`
	ClickNum  int       `json:"clickNum"`
	QrCode    string    `json:"qrCode"`
	CreatedAt time.Time `json:"createdAt"`
	ExpireAt  time.Time `json:"expireAt"`
}
