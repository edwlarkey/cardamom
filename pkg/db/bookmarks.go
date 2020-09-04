package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
)

// InsertBookmark adds a bookmark to the DB
func (m *DB) InsertBookmark(bookmark *models.Bookmark) error {
	m.DB.Create(bookmark)

	return nil
}

// UpdateBookmark updates a single bookmark in the DB
func (m *DB) UpdateBookmark(id int64, title string, tags []string) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{}
	err := m.DB.First(&bookmark, id).Error
	if err != nil {
		return nil, models.ErrNoRecord
	}

	bookmark.Title = title
	m.DB.Model(&bookmark).Association("Tags").Clear()

	for _, tag_name := range tags {
		tag, err := m.CreateIfNotExists(tag_name)
		if err != nil {
			return nil, err
		}
		m.DB.Model(&bookmark).Association("Tags").Append(tag)
	}

	m.DB.Save(&bookmark)

	return bookmark, nil
}

// GetBookmark gets a single bookmark from the DB
func (m *DB) GetBookmark(id int64) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{}

	err := m.DB.Preload("Tags").First(&bookmark, id).Error
	if err != nil {
		return nil, models.ErrNoRecord
	}

	return bookmark, nil
}

// LatestBookmarks gets the 100 most recent bookmarks from the DB
func (m *DB) LatestBookmarks() ([]*models.Bookmark, error) {
	v := []models.Bookmark{}
	m.DB.Preload("Tags").Limit(100).Find(&v)

	bookmarks := []*models.Bookmark{}

	for i := 0; i < len(v); i++ {
		bookmarks = append(bookmarks, &v[i])
	}

	return bookmarks, nil
}
