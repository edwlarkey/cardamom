package models

import (
	"encoding/gob"
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Bookmark struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string
	URL       string `gorm:"unique"`
	Read      int    `gorm:"DEFAULT 0"`
	Tags      []*Tag `gorm:"association_autoupdate:false;many2many:bookmark_tags;association_joinable_foreignkey:tag_id"`
}

type Tag struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"unique_index"`
}

type User struct {
	ID             int `gorm:"primary_key"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time `sql:"index"`
	Name           string
	Email          string `gorm:"unique_index"`
	HashedPassword []byte
}

func init() {
	gob.Register(&User{})
}
