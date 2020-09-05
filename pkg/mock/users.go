package mock

import (
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
	"gorm.io/gorm"
)

var mockUser = &models.User{
	Model: gorm.Model{
		ID:        1,
		CreatedAt: time.Now(),
	},
	Name:  "Alice",
	Email: "alice@example.com",
}

func (m *DB) InsertUser(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (m *DB) AuthenticateUser(email, password string) (*models.User, error) {
	switch email {
	case "alice@example.com":
		return mockUser, nil
	default:
		return nil, models.ErrInvalidCredentials
	}
}

func (m *DB) GetUser(id uint) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
