package main

import (
	"net/http"

	"git.sr.ht/~edwlarkey/cardamom/internal/validator"
)

type tagCreateForm struct {
	Name                string `form:"name"`
	validator.Validator `form:"-"`
}

func (app *application) tagList(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("HX-Request") == "true" {
		userId := app.authenticatedUser(r)
		data := app.newTemplateData(r)
		tags, err := app.tags.All(userId)
		if err != nil {
			app.serverError(w, err)
			return
		}
		data.Tags = tags

		app.renderPartial(w, http.StatusOK, "tag-options", data)
	} else {
		app.clientError(w, http.StatusBadRequest)
		return
	}
}

func (app *application) tagCreatePost(w http.ResponseWriter, r *http.Request) {
	var form bookmarkCreateForm

	userId := app.authenticatedUser(r)
	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	_, err = app.tags.Insert(form.Name, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := app.newTemplateData(r)
	tags, err := app.tags.All(userId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Tags = tags

	app.renderPartial(w, http.StatusOK, "tag-options", data)
}
