package postgres

import (
	"database/sql"

	"github.com/edwlarkey/cardamom/pkg/models"
)

// Insert adds an tag to the DB
func (d *DB) InsertTag(name string) (int, error) {
	r, err := d.DB.Exec("INSERT INTO tags (name) VALUES (?)", name)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get gets a single tag from the DB
func (d *DB) GetTag(id int) (*models.Tag, error) {
	tag := &models.Tag{}
	stmt := `SELECT * FROM tags WHERE id = ? LIMIT 1`
	err := d.DB.Get(&tag, stmt, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return tag, nil
}

// CreateIfNotExists gets a single tag from the DB or create it if it doesn't exist
func (d *DB) CreateIfNotExists(name string) (*models.Tag, error) {
	var tag *models.Tag

	stmt := `SELECT * FROM tags WHERE name = ? LIMIT 1`
	err := d.DB.Get(tag, stmt, name)
	if err == sql.ErrNoRows {
		id, err := d.InsertTag(name)
		if err != nil {
			return nil, err
		}
		tag, err = d.GetTag(id)
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return tag, nil
}

// GetTags gets all tags from the DB
func (d *DB) GetTags() ([]*models.Tag, error) {
	tags := []*models.Tag{}

	stmt := `SELECT * FROM tags`
	err := d.DB.Select(&tags, stmt)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return tags, nil
}
