package main

import (
	"bytes"
	"net/http"
	"net/url"
	"testing"
)

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
		{"Bookmark List", "/app/bookmark", http.StatusTemporaryRedirect, nil},
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
		{"Bookmark List", "/app/bookmark", http.StatusOK, []byte("Bookmark Title Here")},
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
