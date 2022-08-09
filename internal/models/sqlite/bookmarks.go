package sqlite

import (
	"database/sql"
	"errors"
	"html/template"
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
	"github.com/jmoiron/sqlx"
)

type BookmarkModel struct {
	DB *sqlx.DB
}

func (m *BookmarkModel) Insert(title string, url string, content template.HTML, user_id int) (int, error) {
	stmt := `INSERT INTO bookmarks (title, url, content, user_id, created) VALUES (
		?,
		?,
		?,
		?,
		?
	);`

	result, err := m.DB.Exec(stmt, title, url, content, user_id, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *BookmarkModel) Get(user_id int, id int) (*models.Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, user_id, id)

	s := &models.Bookmark{}

	err := row.Scan(&s.ID, &s.Title, &s.Url, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *BookmarkModel) Latest(user_id int) ([]*models.Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookmarks := []*models.Bookmark{}

	for rows.Next() {
		s := &models.Bookmark{}

		err := rows.Scan(&s.ID, &s.Title, &s.Url, &s.Content, &s.Created)
		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}
