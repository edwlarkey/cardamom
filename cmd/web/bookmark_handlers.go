package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
	"git.sr.ht/~edwlarkey/cardamom/internal/validator"
	"github.com/julienschmidt/httprouter"
)

func (app *application) bookmarkList(w http.ResponseWriter, r *http.Request) {
	userId := app.authenticatedUser(r)
	bookmarks, err := app.bookmarks.Latest(userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, bookmark := range bookmarks {
		tags, err := app.tags.GetTagsForBookmark(bookmark.ID, userId)
		if err != nil {
			app.serverError(w, err)
			return
		}
		bookmark.Tags = tags
	}

	data := app.newTemplateData(r)
	data.Bookmarks = bookmarks

	app.render(w, http.StatusOK, "home.tmpl", data)
}

func (app *application) bookmarkView(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	bookmark, err := app.bookmarks.Get(userId, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	tags, err := app.tags.GetTagsForBookmark(id, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	bookmark.Tags = tags

	data := app.newTemplateData(r)
	data.Bookmark = bookmark

	app.render(w, http.StatusOK, "view.tmpl", data)
}

func (app *application) bookmarkDelete(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	err = app.bookmarks.Delete(userId, id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bookmark successfully deleted!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

type bookmarkCreateForm struct {
	Title               string `form:"title"`
	Name                string `form:"name"`
	Url                 string `form:"url"`
	Tags                []int  `form:"tags"`
	validator.Validator `form:"-"`
}

func (app *application) bookmarkCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = bookmarkCreateForm{}

	app.render(w, http.StatusOK, "create.tmpl", data)
}

func (app *application) bookmarkCreatePost(w http.ResponseWriter, r *http.Request) {
	var form bookmarkCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		// app.clientError(w, http.StatusBadRequest)
		app.serverError(w, err)
		return
	}

	form.CheckField(validator.NotBlank(form.Url), "url", "URL cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	userId := app.authenticatedUser(r)

	title, _, content, err := getPageContent(form.Url)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(form.Title) > 0 {
		title = form.Title
	}

	id, err := app.bookmarks.Insert(title, form.Url, content, userId, form.Tags)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bookmark successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/bookmark/view/%d", id), http.StatusSeeOther)
}

func (app *application) bookmarkMarkRead(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	err = app.bookmarks.MarkRead(userId, id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	bookmark, err := app.bookmarks.Get(userId, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	tags, err := app.tags.GetTagsForBookmark(id, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	bookmark.Tags = tags

	data := app.newTemplateData(r)
	data.Bookmark = bookmark

	http.Redirect(w, r, fmt.Sprintf("/bookmark/view/%d", id), http.StatusSeeOther)
}

func (app *application) bookmarkMarkUnread(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	userId := app.authenticatedUser(r)

	err = app.bookmarks.MarkUnread(userId, id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	bookmark, err := app.bookmarks.Get(userId, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	tags, err := app.tags.GetTagsForBookmark(id, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}
	bookmark.Tags = tags

	data := app.newTemplateData(r)
	data.Bookmark = bookmark

	http.Redirect(w, r, fmt.Sprintf("/bookmark/view/%d", id), http.StatusSeeOther)
}
