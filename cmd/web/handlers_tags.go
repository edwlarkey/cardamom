package main

import (
	"net/http"

	"github.com/edwlarkey/cardamom/pkg/forms"
)

func (app *application) createTag(w http.ResponseWriter, r *http.Request) {
	session, err := app.store.Get(r, "cardamom-session")
	if err != nil {
		app.serverError(w, err)
		return
	}
	err = r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("name")

	if !form.Valid() {
		app.render(w, r, "create_tag.page.tmpl", &templateData{Form: form})
		return
	}

	_, err = app.db.InsertTag(form.Get("name"))
	if err != nil {
		app.serverError(w, err)
		return
	}

	session.AddFlash("Tag added successfully!")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/app/tag", http.StatusSeeOther)
}

func (app *application) tagList(w http.ResponseWriter, r *http.Request) {
	v, err := app.db.GetTags()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "tags.page.tmpl", &templateData{
		Tags: v,
		Form: forms.New(nil),
	})
}
