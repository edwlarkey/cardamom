package main

import (
	"errors"
	"net/http"
	"strconv"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
	"git.sr.ht/~edwlarkey/cardamom/internal/validator"
	"github.com/julienschmidt/httprouter"
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

func (app *application) tagView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	tag, err := app.tags.Get(userId, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	bookmarks, err := app.bookmarks.GetBookmarksForTag(userId, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	for _, bookmark := range bookmarks {
		tags, err := app.tags.GetTagsForBookmark(userId, bookmark.ID)
		if err != nil {
			app.serverError(w, err)
			return
		}
		bookmark.Tags = tags
	}

	data := app.newTemplateData(r)
	data.Tag = tag
	data.Bookmarks = bookmarks

	app.render(w, http.StatusOK, "view_tag.tmpl", data)
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

func (app *application) tagDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	err = app.tags.Delete(userId, id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Tag successfully deleted!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
