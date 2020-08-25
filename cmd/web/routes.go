package main

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"gitlab.com/edwlarkey/cardamom/pkg/assets"
)

func (app *application) routes() http.Handler {
	csrfMiddleware := csrf.Protect([]byte(app.config.Web.CSRF), csrf.Secure(app.config.Web.CSRFSecure))

	r := mux.NewRouter()
	r.Use(app.recoverPanic)
	r.Use(app.logRequest)
	r.Use(secureHeaders)
	r.Use(csrfMiddleware)
	auth := r.PathPrefix("/app").Subrouter()
	auth.Use(app.requireAuthenticatedUser)

	// Bookmarks
	auth.HandleFunc("/bookmark/create", app.createBookmarkForm).Methods("GET")
	auth.HandleFunc("/bookmark/create", app.createBookmark).Methods("POST")
	auth.HandleFunc("/bookmark/{id}/edit", app.editBookmarkForm).Methods("GET")
	auth.HandleFunc("/bookmark/{id}/edit", app.updateBookmark).Methods("POST")
	auth.HandleFunc("/bookmark/{id}", app.showBookmark).Methods("GET")

	// Tags
	auth.HandleFunc("/tag/create", app.createTagForm).Methods("GET")
	auth.HandleFunc("/tag/create", app.createTag).Methods("POST")

	// User
	r.HandleFunc("/signup", app.signupUserForm).Methods("GET")
	r.HandleFunc("/signup", app.signupUser).Methods("POST")
	r.HandleFunc("/login", app.loginUserForm).Methods("GET")
	r.HandleFunc("/login", app.loginUser).Methods("POST")
	auth.HandleFunc("/user/logout", app.logoutUser).Methods("POST")

	// Static assets
	r.PathPrefix("/static/").Handler(assets.Handler)
	r.HandleFunc("/ping", ping)

	// Home
	r.HandleFunc("/", app.home)

	return r

}
