package postgresql

import (
	"database/sql"

	"github.com/edwlarkey/cardamom/pkg/models"
)

// InsertBookmark adds a bookmark to the DB
func (d *DB) InsertBookmark(bookmark *models.Bookmark) error {
	tx, err := d.DB.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	r, err := tx.Exec("INSERT INTO bookmarks (title, url) VALUES (?, ?)", bookmark.Title, bookmark.URL)
	if err != nil {
		tx.Rollback()
		return err
	}

	bookmark.ID, err = r.LastInsertId()
	if err != nil {
		return err
	}

	return tx.Commit()
}

// UpdateBookmark updates a single bookmark in the DB
func (d *DB) UpdateBookmark(bookmark *models.Bookmark) error {
	tx := d.DB.MustBegin()

	tx.NamedExec("UPDATE bookmarks (title, url, read, content) VALUES (:title, :url, :read, :content)", bookmark)

	// res := m.DB.Model(&bookmark).Association("Tags").Clear()

	// for _, tag_name := range tags {
	// 	tag, err := m.CreateIfNotExists(tag_name)
	// 	if err != nil {
	// 		tx.Rollback()
	// 		return nil, err
	// 	}
	// 	res := m.DB.Model(&bookmark).Association("Tags").Append(tag)

	// 	if res.Error != nil {
	// 		tx.Rollback()
	// 		return nil, res.Error
	// 	}
	// }

	return tx.Commit()
}

// BookmarkByID gets a single bookmark from the DB
func (d *DB) BookmarkByID(bookmarkID int64) (*models.Bookmark, error) {
	bookmark := &models.Bookmark{}

	stmt := `SELECT * FROM bookmarks WHERE id = ? LIMIT 1`
	err := d.DB.Get(&bookmark, stmt, bookmarkID)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return bookmark, nil
}

// LatestBookmarks gets the 100 most recent bookmarks from the DB
func (d *DB) LatestBookmarks() ([]*models.Bookmark, error) {
	bookmarks := []*models.Bookmark{}

	stmt := `SELECT * FROM bookmarks LIMIT 10`
	err := d.DB.Select(&bookmarks, stmt)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return bookmarks, nil
}
