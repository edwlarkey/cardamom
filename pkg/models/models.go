package models

import (
	"encoding/gob"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Bookmark struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Title     string
	URL       string `gorm:"uniqueIndex"`
	Read      int    `gorm:"default 0"`
	Tags      []*Tag `gorm:"many2many:bookmark_tags;jointForeignKey:tag_id"`
}

type Tag struct {
	ID        int64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"uniqueIndex"`
}

type User struct {
	ID             int64 `gorm:"primaryKey"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Name           string
	Email          string `gorm:"uniqueIndex"`
	HashedPassword []byte
}

func init() {
	gob.Register(&User{})
}
