package models

import (
	"html/template"
	"time"
)

type BookmarkModelInterface interface {
	Insert(title string, url string, content template.HTML, userId int, tags []int) (int, error)
	Get(userId int, id int) (*Bookmark, error)
	Latest(userId int) ([]*Bookmark, error)
	MarkRead(userId int, id int) error
	MarkUnread(userId int, id int) error
}

type Bookmark struct {
	ID      int
	Title   string
	Url     string
	Content template.HTML
	UserId  int
	Read    bool
	Created time.Time
	Tags    []*Tag
}
