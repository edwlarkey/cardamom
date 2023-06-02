package models

import "time"

type TagModelInterface interface {
	Insert(name string, userId int) (int, error)
	Get(id int, userId int) (*Tag, error)
	GetTagsForBookmark(bookmarkId int, userId int) ([]*Tag, error)
	All(userId int) ([]*Tag, error)
}

type Tag struct {
	ID      int
	Name    string
	UserId  int
	Created time.Time
}
