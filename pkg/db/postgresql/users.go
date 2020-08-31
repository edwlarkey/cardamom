package postgresql

import (
	"database/sql"

	"github.com/edwlarkey/cardamom/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (d *DB) InsertUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (username, email, hashed_password, created_at)
    VALUES($1, $2, $3, UTC_TIMESTAMP())`

	_, err = d.DB.Exec(stmt, name, email, string(hashedPassword))
	if err != nil {
		return err
	}
	return err
}

func (d *DB) AuthenticateUser(email, password string) (int, error) {
	var id int
	var hashedPassword []byte
	row := d.DB.QueryRow("SELECT id, hashed_password FROM users WHERE email = ?", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, models.ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}

	return id, nil
}

func (d *DB) GetUser(id int) (*models.User, error) {
	s := &models.User{}

	stmt := `SELECT id, username, email, created_at FROM users WHERE id = ?`
	err := d.DB.Get(&s, stmt, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}
