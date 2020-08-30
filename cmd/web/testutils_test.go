package main

import (
	"html"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"net/url"
	"regexp"
	"testing"

	"github.com/edwlarkey/cardamom/pkg/mock"

	"github.com/gorilla/sessions"
)

var csrfTokenRX = regexp.MustCompile(`<input type="hidden" name="gorilla.csrf.Token" value="(.+)">`)

func extractCSRFToken(t *testing.T, body []byte) string {
	// FindSubmatch returns an array with entire match in first element
	// and captured value in the second
	matches := csrfTokenRX.FindSubmatch(body)
	if len(matches) < 2 {
		t.Fatal("no csrf token found in body")
	}

	return html.UnescapeString(string(matches[1]))
}

func newTestApplication(t *testing.T) *application {

	store := sessions.NewCookieStore([]byte("3dSm5MnygFHh7XidAtbskXrjbwfoJcbJ"))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}

	return &application{
		errorLog:  log.New(ioutil.Discard, "", 0),
		infoLog:   log.New(ioutil.Discard, "", 0),
		store:     store,
		db:        &mock.DB{},
		templates: initTemplates("base"),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	jar, err := cookiejar.New(nil)

	if err != nil {
		t.Fatal(err)
	}

	ts.Client().Jar = jar

	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, []byte) {
	rs, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	// Read the response body.
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Return the response status, headers and body.
	return rs.StatusCode, rs.Header, body
}

func (ts *testServer) postForm(t *testing.T, urlPath string, form url.Values) (int, http.Header, []byte) {
	rs, err := ts.Client().PostForm(ts.URL+urlPath, form)
	if err != nil {
		t.Fatal(err)
	}

	// Read the response body.
	defer rs.Body.Close()
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Return the response status, headers and body.
	return rs.StatusCode, rs.Header, body
}
