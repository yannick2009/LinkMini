package model

import "time"

type URL struct {
	LongURL   string    `json:"longUrl"`
	URLHash   string    `json:"urlHash"`
	ClickNum  int       `json:"clickNum"`
	QrCode    []byte    `json:"qrCode"`
	CreatedAt time.Time `json:"createdAt"`
	ExpireAt  time.Time `json:"expireAt"`
}
