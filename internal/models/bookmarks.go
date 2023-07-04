package models

import (
	"html/template"
	"time"
)

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
