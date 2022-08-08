package models

import (
	"database/sql"
	"errors"
	"time"
)

type BookmarkModelInterface interface {
	Insert(title string, url string, content string, user_id int) (int, error)
	Get(user_id int, id int) (*Bookmark, error)
	Latest(user_id int) ([]*Bookmark, error)
}

type Bookmark struct {
	ID      int
	Title   string
	Url     string
	Content string
	UserId  int
	Created time.Time
}

type BookmarkModel struct {
	DB *sql.DB
}

func (m *BookmarkModel) Insert(title string, url string, content string, user_id int) (int, error) {
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

func (m *BookmarkModel) Get(user_id int, id int) (*Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, user_id, id)

	s := &Bookmark{}

	err := row.Scan(&s.ID, &s.Title, &s.Url, &s.Content, &s.Created)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *BookmarkModel) Latest(user_id int) ([]*Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookmarks := []*Bookmark{}

	for rows.Next() {
		s := &Bookmark{}

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
