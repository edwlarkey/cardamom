package sqlite

import (
	"context"
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

func (m *BookmarkModel) Insert(title string, url string, content template.HTML, userId int, tags []int) (int, error) {
	tx, err := m.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return 0, err
	}
	stmt := `INSERT INTO bookmarks (title, url, content, read, user_id, created) VALUES (
		?,
		?,
		?,
		?,
		?,
		?
	);`

	result, err := tx.Exec(stmt, title, url, content, 1, userId, time.Now())
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	tagStmt := `INSERT INTO bookmark_tags (bookmark_id, tag_id, user_id) VALUES (
		?,
		?,
		?
	);`

	for _, tagId := range tags {
		_, err := tx.Exec(tagStmt, id, tagId, userId)
		if err != nil {
			return 0, err
		}
	}

	tx.Commit()

	return int(id), nil
}

func (m *BookmarkModel) Get(userId int, id int) (*models.Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, userId, id)

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

func (m *BookmarkModel) Latest(userId int) ([]*models.Bookmark, error) {
	stmt := `SELECT id, title, url, content, created FROM bookmarks
    WHERE user_id = ? ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookmarks := []*models.Bookmark{}

	for rows.Next() {
		b := &models.Bookmark{}

		err := rows.Scan(&b.ID, &b.Title, &b.Url, &b.Content, &b.Created)
		if err != nil {
			return nil, err
		}

		bookmarks = append(bookmarks, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return bookmarks, nil
}
