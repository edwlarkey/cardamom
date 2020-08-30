package mock

import (
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
)

var mockUser = &models.User{
	ID:        1,
	Name:      "Alice",
	Email:     "alice@example.com",
	CreatedAt: time.Now(),
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

func (m *DB) GetUser(id int) (*models.User, error) {
	switch id {
	case 1:
		return mockUser, nil
	default:
		return nil, models.ErrNoRecord
	}
}
