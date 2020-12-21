package mock

import (
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/gorm"
)

var mockTag = &models.Tag{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
	},
	Name: "Tag Title Here",
}

func (m *DB) InsertTag(tag *models.Tag) (uint, error) {
	return 2, nil
}

func (m *DB) GetTag(id uint) (*models.Tag, error) {
	switch id {
	case 1:
		return mockTag, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *DB) GetTagByName(name string, user *models.User) (*models.Tag, error) {
	switch name {
	case "Tag Title Here":
		return mockTag, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *DB) CreateIfNotExists(name string) (*models.Tag, error) {
	switch name {
	case "Tag Title Here":
		return mockTag, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *DB) GetTags(*models.User) ([]*models.Tag, error) {
	return []*models.Tag{mockTag}, nil
}
