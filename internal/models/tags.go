package models

import "time"

type TagModelInterface interface {
	Insert(name string, user_id int) (int, error)
	Get(user_id int, id int) (*Tag, error)
	All(user_id int) ([]*Tag, error)
}

type Tag struct {
	ID      int
	Name    string
	UserId  int
	Created time.Time
}
