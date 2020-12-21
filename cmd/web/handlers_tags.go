package main

import (
	"net/http"

	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
)

func (app *application) createTag(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
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

	user := app.authenticatedUser(r)

	tag := models.Tag{}
	tag.Name = form.Get("name")
	tag.UserID = user.ID

	_, err = app.db.InsertTag(&tag)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Tag added successfully!")

	http.Redirect(w, r, "/app/tag", http.StatusSeeOther)
}

func (app *application) tagList(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)

	v, err := app.db.GetTags(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "tags.page.tmpl", &templateData{
		Tags: v,
		Form: forms.New(nil),
	})
}
