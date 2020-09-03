package postgres_test

import (
	"testing"

	"github.com/edwlarkey/cardamom/pkg/models"
)

var pw = "secret"

var u = &models.User{
	Email: "alice@example.com",
}

func TestInsertUser(t *testing.T) {
	err := db.InsertUser(u.Email, pw)
	if err != nil {
		t.Errorf("Inserting User failed %s", err)
	}
}

func TestAuthenticateUser(t *testing.T) {
	tests := []struct {
		name      string
		email     string
		password  string
		wantError bool
	}{
		{"Correct Login", u.Email, pw, false},
		{"Incorrect Password", u.Email, "otherpassword", true},
		{"Incorrect Email", "bob@example.com", pw, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := db.AuthenticateUser(tt.email, tt.password)
			if !tt.wantError && err != nil {
				t.Errorf("Authenticateing User failed %s", err)
			}

			if user != nil {
				if user.Email != u.Email {
					t.Errorf("want %s; got %s", tt.email, user.Email)
				}
			}

		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name      string
		id        int64
		wantError bool
	}{
		{"Correct User ID", 1, false},
		{"Incorrect User ID", 2, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := db.GetUser(tt.id)
			if !tt.wantError && err != nil {
				t.Errorf("Authenticateing User failed %s", err)
			}

			if user != nil {
				if user.Email != u.Email {
					t.Errorf("want %d; got %d", tt.id, user.ID)
				}
			}

		})
	}
}
