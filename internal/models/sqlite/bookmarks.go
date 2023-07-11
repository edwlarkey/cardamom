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

	result, err := tx.Exec(stmt, title, url, content, 0, userId, time.Now())
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
	stmt := `SELECT id, title, url, content, created, read FROM bookmarks
    WHERE user_id = ? AND id = ?`

	row := m.DB.QueryRow(stmt, userId, id)

	s := &models.Bookmark{}

	err := row.Scan(&s.ID, &s.Title, &s.Url, &s.Content, &s.Created, &s.Read)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *BookmarkModel) Delete(userId int, id int) error {
	tx, err := m.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	stmt := `DELETE FROM bookmarks WHERE id = ? and user_id = ?;`

	_, err = tx.Exec(stmt, id, userId)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (m *BookmarkModel) Latest(userId int) ([]*models.Bookmark, error) {
	stmt := `SELECT id, title, url, content, created, read FROM bookmarks
    WHERE user_id = ? ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookmarks := []*models.Bookmark{}

	for rows.Next() {
		b := &models.Bookmark{}

		err := rows.Scan(&b.ID, &b.Title, &b.Url, &b.Content, &b.Created, &b.Read)
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

func (m *BookmarkModel) GetBookmarksForTag(userId int, tagId int) ([]*models.Bookmark, error) {
	stmt := `SELECT bookmark_id FROM bookmark_tags
    WHERE tag_id = ? AND user_id = ?`

	rows, err := m.DB.Query(stmt, tagId, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bookmarks := []*models.Bookmark{}
	var bookmarkIds []int

	for rows.Next() {
		var t int

		err := rows.Scan(&t)
		if err != nil {
			return nil, err
		}

		bookmarkIds = append(bookmarkIds, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	for _, id := range bookmarkIds {
		bookmark, err := m.Get(userId, id)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

func (m *BookmarkModel) MarkRead(userId int, id int) error {
	tx, err := m.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	stmt := `UPDATE bookmarks SET read = 1 WHERE id = ? and user_id = ?;`

	_, err = tx.Exec(stmt, id, userId)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (m *BookmarkModel) MarkUnread(userId int, id int) error {
	tx, err := m.DB.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}
	stmt := `UPDATE bookmarks SET read = 0 WHERE id = ? and user_id = ?;`

	_, err = tx.Exec(stmt, id, userId)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
