package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB interface {
	Connect(string, string) error
	Migrate() error
	Close()
	LatestBookmarks() ([]*models.Bookmark, error)
	BookmarkByID(int64) (*models.Bookmark, error)
	InsertBookmark(*models.Bookmark) error
	UpdateBookmark(*models.Bookmark) error
	GetTags() ([]*models.Tag, error)
	InsertTag(string) (int, error)
	InsertUser(string, string) error
	AuthenticateUser(string, string) (*models.User, error)
	GetUser(int64) (*models.User, error)
}
