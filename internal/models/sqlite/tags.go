package sqlite

import (
	"database/sql"
	"errors"
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
	"github.com/jmoiron/sqlx"
)

type TagModel struct {
	DB *sqlx.DB
}

func (m *TagModel) Insert(name string, userId int) (int, error) {
	stmt := `INSERT INTO tags (name, user_id, created) VALUES (
		?,
		?,
		?
	);`

	result, err := m.DB.Exec(stmt, name, userId, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TagModel) Get(userId int, id int) (*models.Tag, error) {
	stmt := `SELECT id, name, created FROM tags
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, userId, id)

	s := &models.Tag{}

	err := row.Scan(&s.ID, &s.Name, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *TagModel) GetTagsForBookmark(userId int, bookmarkId int) ([]*models.Tag, error) {
	stmt := `SELECT tag_id FROM bookmark_tags
    WHERE bookmark_id = ? AND user_id = ?`

	rows, err := m.DB.Query(stmt, bookmarkId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*models.Tag{}
	var tagIds []int

	for rows.Next() {
		var t int

		err := rows.Scan(&t)
		if err != nil {
			return nil, err
		}

		tagIds = append(tagIds, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for _, id := range tagIds {
		tag, err := m.Get(userId, id)
		if err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

func (m *TagModel) All(user_id int) ([]*models.Tag, error) {
	stmt := `SELECT id, name, created FROM tags
    WHERE user_id = ? ORDER BY name ASC`

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tags := []*models.Tag{}

	for rows.Next() {
		s := &models.Tag{}

		err := rows.Scan(&s.ID, &s.Name, &s.Created)
		if err != nil {
			return nil, err
		}

		tags = append(tags, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
