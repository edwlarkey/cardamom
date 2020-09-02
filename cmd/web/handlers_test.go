package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

func TestPing(t *testing.T) {

	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
}

func TestShowBookmarkAnon(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/app/bookmark/1", http.StatusTemporaryRedirect, nil},
		{"Non-existent ID", "/app/bookmark/2", http.StatusTemporaryRedirect, nil},
		{"Negative ID", "/app/bookmark/-1", http.StatusTemporaryRedirect, nil},
		{"Decimal ID", "/app/bookmark/1.23", http.StatusTemporaryRedirect, nil},
		{"String ID", "/app/bookmark/foo", http.StatusTemporaryRedirect, nil},
		{"Empty ID", "/app/bookmark/", http.StatusNotFound, nil},
		{"Trailing slash", "/app/bookmark/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}

		})
	}
}

func TestShowBookmark(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, loginbody := ts.get(t, "/login")
	csrfToken := extractCSRFToken(t, loginbody)

	form := url.Values{}
	form.Add("email", "alice@example.com")
	form.Add("password", "whatever")
	form.Add("gorilla.csrf.Token", csrfToken)

	_, _, _ = ts.postForm(t, "/login", form)

	tests := []struct {
		name     string
		urlPath  string
		wantCode int
		wantBody []byte
	}{
		{"Valid ID", "/app/bookmark/1", http.StatusOK, []byte("Bookmark Title Here")},
		{"Non-existent ID", "/app/bookmark/2", http.StatusNotFound, nil},
		{"Negative ID", "/app/bookmark/-1", http.StatusNotFound, nil},
		{"Decimal ID", "/app/bookmark/1.23", http.StatusNotFound, nil},
		{"String ID", "/app/bookmark/foo", http.StatusNotFound, nil},
		{"Empty ID", "/app/bookmark/", http.StatusNotFound, nil},
		{"Trailing slash", "/app/bookmark/1/", http.StatusNotFound, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			code, _, body := ts.get(t, tt.urlPath)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}

		})
	}
}

func TestSignupUser(t *testing.T) {
	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	_, _, body := ts.get(t, "/signup")
	csrfToken := extractCSRFToken(t, body)

	tests := []struct {
		name         string
		userEmail    string
		userPassword string
		csrfToken    string
		wantCode     int
		wantBody     []byte
	}{
		{"Valid submission", "bob@example.com", "validPa$$word", csrfToken, http.StatusSeeOther, nil},
		{"Empty email", "", "validPa$$word", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Empty password", "bob@example.com", "", csrfToken, http.StatusOK, []byte("This field cannot be blank")},
		{"Invalid email (incomplete domain)", "bob@example.", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing @)", "bobexample.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Invalid email (missing local part)", "@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("This field is invalid")},
		{"Short password", "bob@example.com", "pa$$", csrfToken, http.StatusOK, []byte("This field is too short (minimum is 8 characters)")},
		{"Duplicate email", "dupe@example.com", "validPa$$word", csrfToken, http.StatusOK, []byte("Address is already in use")},
		{"Invalid CSRF Token", "", "", "wrongToken", http.StatusForbidden, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			form := url.Values{}
			form.Add("email", tt.userEmail)
			form.Add("password", tt.userPassword)
			form.Add("gorilla.csrf.Token", tt.csrfToken)

			code, _, body := ts.postForm(t, "/signup", form)

			if code != tt.wantCode {
				t.Errorf("want %d; got %d", tt.wantCode, code)
			}

			if !bytes.Contains(body, tt.wantBody) {
				t.Errorf("want body to contain %q", tt.wantBody)
			}

		})
	}
}
