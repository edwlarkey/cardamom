package main

import (
	"net/http"
)

type Option struct {
	Value    string
	Selected bool
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if app.authenticatedUser(r) != nil {
		http.Redirect(w, r, "/app/bookmark", http.StatusTemporaryRedirect)
		return
	}
	app.render(w, r, "home.page.tmpl", &templateData{})
}

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		app.serverError(w, err)
		return
	}
}
