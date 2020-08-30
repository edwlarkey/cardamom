package db

import (
	"github.com/edwlarkey/cardamom/pkg/models"
	"golang.org/x/crypto/bcrypt"
)

func (m *DB) InsertUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	user := models.User{Name: name, Email: email, HashedPassword: hashedPassword}

	m.DB.Create(&user)

	return err
}

func (m *DB) AuthenticateUser(email, password string) (*models.User, error) {
	user := models.User{}
	if m.DB.Where("email = ?", email).First(&user).RecordNotFound() {
		return nil, models.ErrInvalidCredentials
	}

	err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil, models.ErrInvalidCredentials
	} else if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *DB) GetUser(id int) (*models.User, error) {
	user := &models.User{}
	if err := m.DB.First(&user, id).Error; err != nil {
		return nil, models.ErrNoRecord
	}

	return user, nil
}
