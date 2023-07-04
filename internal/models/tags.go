package models

import "time"

type Tag struct {
	ID      int
	Name    string
	UserId  int
	Created time.Time
}
