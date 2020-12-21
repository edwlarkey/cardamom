package main

import (
	"net/http"
)

type Option struct {
	Value    string
	Selected bool
}

func (app *application) ping(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) settings(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "settings.page.tmpl", &templateData{
		Hostname: app.config.Web.Hostname,
	})
}

func (app *application) search(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("OK"))
	if err != nil {
		app.serverError(w, err)
		return
	}
}
