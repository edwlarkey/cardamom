package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
)

// Insert adds an tag to the DB
func (m *DB) InsertTag(tag *models.Tag) (uint, error) {

	m.DB.Create(&tag)

	return tag.ID, nil
}

// Get gets a single tag from the DB
func (m *DB) GetTag(id uint) (*models.Tag, error) {
	tag := &models.Tag{}

	err := m.DB.First(&tag, id).Error
	if err != nil {
		return nil, models.ErrNoRecord
	}

	return tag, nil
}

// Get gets a single tag from the DB by name
func (m *DB) GetTagByName(name string, user *models.User) (*models.Tag, error) {
	tag := &models.Tag{}

	err := m.DB.Where("name = ?", name).Scopes(currentUser(user)).First(&tag).Error
	if err != nil {
		return nil, models.ErrNoRecord
	}

	return tag, nil
}

// CreateIfNotExists gets a single tag from the DB or create it if it doesn't exist
func (m *DB) CreateIfNotExists(name string) (*models.Tag, error) {
	var tag models.Tag

	res := m.DB.FirstOrCreate(&tag, &models.Tag{Name: name})
	if res.Error != nil {
		return nil, res.Error
	}

	return &tag, nil
}

// GetTags gets all tags from the DB
func (m *DB) GetTags(user *models.User) ([]*models.Tag, error) {
	t := []models.Tag{}
	m.DB.Scopes(currentUser(user)).Find(&t)

	tags := []*models.Tag{}

	for i := 0; i < len(t); i++ {
		tags = append(tags, &t[i])
	}

	return tags, nil
}

// Latest gets the 10 most recent tags from the DB
func (m *DB) LatestTags() ([]*models.Tag, error) {
	v := []models.Tag{}
	m.DB.Limit(10).Find(&v)

	tags := []*models.Tag{}

	for i := 0; i < len(v); i++ {
		tags = append(tags, &v[i])
	}

	return tags, nil
}
