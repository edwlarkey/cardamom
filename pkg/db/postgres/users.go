package postgres

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

	stmt := "INSERT INTO users (email, password, created_at) VALUES($1, $2, now())"

	_, err = d.DB.Exec(stmt, email, string(hashedPassword))
	if err != nil {
		return err
	}
	return err
}

func (d *DB) AuthenticateUser(email, password string) (*models.User, error) {
	user := &models.User{}

	stmt := "SELECT id, email, password, created_at FROM users WHERE email = $1"
	err := d.DB.Get(user, stmt, email)
	if err == sql.ErrNoRows {
		return nil, models.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, models.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	user.Password = nil

	return user, nil
}

func (d *DB) GetUser(id int64) (*models.User, error) {
	user := &models.User{}

	stmt := "SELECT id, email, created_at FROM users WHERE id = $1"
	err := d.DB.Get(user, stmt, id)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return user, nil
}
