package models

import (
	"encoding/gob"
	"errors"
	"html/template"

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
	Excerpt string
	User    User
	UserID  uint `gorm:"index:unique_url_user,unique"`
	Content template.HTML
	URL     string `gorm:"index:unique_url_user,unique"`
	Read    int    `gorm:"default 0"`
	Tags    []*Tag `gorm:"many2many:bookmark_tags"`
}

type Tag struct {
	gorm.Model
	Name   string `gorm:"index:unique_tag_user,unique"`
	User   User
	UserID uint `gorm:"index:unique_tag_user,unique"`
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
