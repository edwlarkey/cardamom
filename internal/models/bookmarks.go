package models

import (
	"time"
)

type BookmarkModelInterface interface {
	Insert(title string, url string, content string, user_id int) (int, error)
	Get(user_id int, id int) (*Bookmark, error)
	Latest(user_id int) ([]*Bookmark, error)
}

type Bookmark struct {
	ID      int
	Title   string
	Url     string
	Content string
	UserId  int
	Created time.Time
}
