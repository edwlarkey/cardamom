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

func (m *TagModel) Insert(name string, user_id int) (int, error) {
	stmt := `INSERT INTO tags (name, user_id, created) VALUES (
		?,
		?,
		?
	);`

	result, err := m.DB.Exec(stmt, name, user_id, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *TagModel) Get(user_id int, id int) (*models.Tag, error) {
	stmt := `SELECT id, name, created FROM tags
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, user_id, id)

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

func (m *TagModel) All(user_id int) ([]*models.Tag, error) {
	stmt := `SELECT id, name, created FROM tags
    WHERE user_id = ? ORDER BY id DESC LIMIT 10`

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
