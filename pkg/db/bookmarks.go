package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/gorm/clause"
)

// InsertBookmark adds a bookmark to the DB
func (m *DB) InsertBookmark(bookmark *models.Bookmark) error {
	m.DB.Create(bookmark)

	return nil
}

// UpdateBookmark updates a single bookmark in the DB
func (m *DB) UpdateBookmark(bookmark *models.Bookmark) error {
	err := m.DB.Model(bookmark).Association("Tags").Replace(bookmark.Tags)
	if err != nil {
		return err
	}

	err = m.DB.Omit(clause.Associations).Save(bookmark).Error
	if err != nil {
		return err
	}

	return nil
}

// GetBookmark gets a single bookmark from the DB
func (m *DB) GetBookmark(id uint) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{}

	err := m.DB.Preload("Tags").First(&bookmark, id).Error
	if err != nil {
		return nil, models.ErrNoRecord
	}

	return bookmark, nil
}

// GetBookmarks gets all bookmarks
func (m *DB) GetBookmarks() ([]*models.Bookmark, error) {
	v := []models.Bookmark{}
	m.DB.Preload("Tags").Find(&v)

	bookmarks := []*models.Bookmark{}

	for i := 0; i < len(v); i++ {
		bookmarks = append(bookmarks, &v[i])
	}

	return bookmarks, nil
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
