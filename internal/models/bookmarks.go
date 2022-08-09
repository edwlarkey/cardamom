package models

import (
	"html/template"
	"time"
)

type BookmarkModelInterface interface {
	Insert(title string, url string, content template.HTML, user_id int) (int, error)
	Get(user_id int, id int) (*Bookmark, error)
	Latest(user_id int) ([]*Bookmark, error)
}

type Bookmark struct {
	ID      int
	Title   string
	Url     string
	Content template.HTML
	UserId  int
	ToRead  int
	Created time.Time
	Tags    []Tag
}
