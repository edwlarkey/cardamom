package models

import (
	"encoding/gob"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Bookmark struct {
	gorm.Model
	Title   string
	Content string
	URL     string `gorm:"uniqueIndex"`
	Read    int    `gorm:"default 0"`
	Tags    []*Tag `gorm:"many2many:bookmark_tags"`
}

type Tag struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`
}

type User struct {
	gorm.Model
	Name           string
	Email          string `gorm:"uniqueIndex"`
	HashedPassword []byte
}

func init() {
	gob.Register(&User{})
}
