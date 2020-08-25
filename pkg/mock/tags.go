package mock

import (
	"time"

	"gitlab.com/edwlarkey/cardamom/pkg/models"
)

var mockTag = &models.Tag{
	ID:        1,
	CreatedAt: time.Now(),
	Name:      "Tag Title Here",
}

func (m *DB) InsertTag(name string) (int, error) {
	return 2, nil
}

func (m *DB) GetTag(id int) (*models.Tag, error) {
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
