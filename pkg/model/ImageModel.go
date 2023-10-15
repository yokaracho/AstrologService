package model

type ImageModel struct {
	ID          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title"`
	Date        string `json:"date" db:"date"`
	URL         string `json:"url" db:"url"`
	HDURL       string `json:"hdurl" db:"hd_url"`
	MediaType   string `json:"media_type" db:"media_type"`
	Explanation string `json:"explanation" db:"explanation"`
	ThumbURL    string `json:"thumbnail_url" db:"thumbnail_url"`
	Copyright   string `json:"copyright" db:"copyright"`
	RAW         []byte `json:"raw" db:"-"`
}
