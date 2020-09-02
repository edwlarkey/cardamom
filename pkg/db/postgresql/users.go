package postgresql

import (
	"database/sql"

	"github.com/edwlarkey/cardamom/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *DB) InsertUser(email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (email, password, created_at)
    VALUES($1, $2, now())`

	_, err = d.DB.Exec(stmt, email, string(hashedPassword))
	if err != nil {
		return err
	}
	return err
}

func (d *DB) AuthenticateUser(email, password string) (*models.User, error) {
	var id int64
	var hashedPassword []byte
	row := d.DB.QueryRow("SELECT id, password FROM users WHERE email = $1", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return nil, models.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, models.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *DB) GetUser(id int64) (*models.User, error) {
	s := &models.User{}

	stmt := `SELECT id, username, email, created_at FROM users WHERE id = $1`
	err := d.DB.Get(&s, stmt, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
