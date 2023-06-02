package mocks

import (
	"html/template"
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
)

var mockBookmark = &models.Bookmark{
	ID:      1,
	UserId:  1,
	Title:   "An old silent pond",
	Url:     "An old silent pond",
	Content: "An old silent pond...",
	Created: time.Now(),
}

type BookmarkModel struct{}

func (m *BookmarkModel) Insert(title string, url string, content template.HTML, userId int, tags []int) (int, error) {
	return 2, nil
}

func (m *BookmarkModel) Get(userId int, id int) (*models.Bookmark, error) {
	switch id {
	case 1:
		return mockBookmark, nil
	default:
		return nil, models.ErrNoRecord
	}
}

func (m *BookmarkModel) Latest(userId int) ([]*models.Bookmark, error) {
	return []*models.Bookmark{mockBookmark}, nil
}
