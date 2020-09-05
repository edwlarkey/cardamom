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

func (m *DB) InsertTag(name string) (uint, error) {
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

func (m *DB) GetTags() ([]*models.Tag, error) {
	return []*models.Tag{mockTag}, nil
}
