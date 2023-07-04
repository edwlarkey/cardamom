package sqlite

import (
	"fmt"
	"testing"

	"git.sr.ht/~edwlarkey/cardamom/internal/assert"
)

func TestUserModelInsert(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name     string
		email    string
		password string
		error    bool
	}{
		{
			name:     "Valid Name",
			email:    "name@example.com",
			password: "password",
			error:    false,
		},
		{
			name:     "Duplicate",
			email:    "alice@example.com",
			password: "passwordd",
			error:    true,
		},
	}

	for idx, tt := range tests {
		idx := idx
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t, fmt.Sprintf("UserModelInsert%d", idx))

			m := UserModel{DB: db}

			err := m.Insert(tt.name, tt.email, tt.password)

			if !tt.error {
				assert.NilError(t, err)
			}
		})
	}
}

func TestUserModelAuthenticate(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name     string
		email    string
		password string
		error    bool
	}{
		{
			name:     "Valid Login",
			email:    "alice@example.com",
			password: "pa$$word",
			error:    false,
		},
		{
			name:     "Invalid Password",
			email:    "alice@example.com",
			password: "password",
			error:    true,
		},
		{
			name:     "Invalid Email",
			email:    "bob@example.com",
			password: "password",
			error:    true,
		},
	}

	for idx, tt := range tests {
		idx := idx
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t, fmt.Sprintf("UserModelAuthenticate%d", idx))

			m := UserModel{DB: db}

			_, err := m.Authenticate(tt.email, tt.password)

			if !tt.error {
				assert.NilError(t, err)
			}
		})
	}
}

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}

	tests := []struct {
		name   string
		userID int
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userID: 2,
			want:   false,
		},
	}

	for idx, tt := range tests {
		idx := idx
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t, fmt.Sprintf("UserModelExists%d", idx))

			m := UserModel{DB: db}

			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
