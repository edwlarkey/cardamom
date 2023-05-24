package models

import (
	"html/template"
	"time"
)

type BookmarkModelInterface interface {
	Insert(title string, url string, content template.HTML, userId int) (int, error)
	Get(userId int, id int) (*Bookmark, error)
	Latest(userId int) ([]*Bookmark, error)
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
