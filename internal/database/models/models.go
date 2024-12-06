package models

import "time"

type User struct {
	ID int64 `json:"id"`

	Username  string    `json:"username"`
	Gender    string    `json:"gender"`
	Birthday  string    `json:"birthday"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserDiary struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type UserEvent struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	EventType string    `json:"event_type"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
type ArticleTag struct {
	ID        int    `json:"id"`
	ArticleID int    `json:"article_id"`
	Tag       string `json:"tag"`
}

type Article struct {
	ID           int    `json:"id"`
	ImageURLs    string `json:"image_urls"`
	Title        string `json:"title"`
	UploadedTime string `json:"uploaded_time"`
	Content      string `json:"content"`
	Tags         string `json:"tags"`
	CreatedAt    string `json:"created_at"`
}
type GlobalConfig struct {
	ID    int    `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
