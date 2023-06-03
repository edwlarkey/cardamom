package mocks

import (
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
)

var mockTag = &models.Tag{
	ID:      1,
	UserId:  1,
	Name:    "Tag1",
	Created: time.Now(),
}

type TagModel struct{}

func (m *TagModel) Insert(name string, userId int) (int, error) {
	return 2, nil
}

func (m *TagModel) Get(id int, userId int) (*models.Tag, error) {
	switch id {
	case 1:
		return mockTag, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *TagModel) GetTagsForBookmark(id int, userId int) ([]*models.Tag, error) {
	switch id {
	case 1:
		return []*models.Tag{mockTag}, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *TagModel) All(userId int) ([]*models.Tag, error) {
	return []*models.Tag{mockTag}, nil
}
