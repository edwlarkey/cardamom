package mock

import (
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/gorm"
)

var mockBookmark = &models.Bookmark{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
	},
	UserID: 1,
	Title:  "Bookmark Title Here",
	URL:    "https://www.archives.gov/founding-docs/constitution-transcript",
	Read:   0,
}

func (m *DB) InsertBookmark(*models.Bookmark) error {
	return nil
}

func (m *DB) UpdateBookmark(*models.Bookmark) error {
	return nil
}

func (m *DB) GetBookmarks(user *models.User) ([]*models.Bookmark, error) {
	return []*models.Bookmark{mockBookmark}, nil
}

func (m *DB) GetBookmark(id uint, user *models.User) (*models.Bookmark, error) {
	switch id {
	case 1:
		if user.ID != mockBookmark.UserID {
			return nil, models.ErrNoRecord
		}
		return mockBookmark, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *DB) LatestBookmarks() ([]*models.Bookmark, error) {
	return []*models.Bookmark{mockBookmark}, nil
}
