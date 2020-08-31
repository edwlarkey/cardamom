package mock

import (
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
)

var mockBookmark = &models.Bookmark{
	ID:        1,
	CreatedAt: time.Now(),
	Title:     "Bookmark Title Here",
	URL:       "https://www.archives.gov/founding-docs/constitution-transcript",
	Read:      0,
}

func (m *DB) InsertBookmark(*models.Bookmark) error {
	return nil
}

func (m *DB) UpdateBookmark(*models.Bookmark) error {
	return nil
}

func (m *DB) BookmarkByID(id int64) (*models.Bookmark, error) {
	switch id {
	case 1:
		return mockBookmark, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *DB) LatestBookmarks() ([]*models.Bookmark, error) {
	return []*models.Bookmark{mockBookmark}, nil
}
