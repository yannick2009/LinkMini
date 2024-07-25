package model

type URL struct {
	ID       string `json:"id,omitempty"`
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
	ClickNum int    `json:"clickNum"`
}
