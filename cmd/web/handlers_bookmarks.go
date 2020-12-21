package main

import (
	"fmt"
	"net/http"

	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
	"github.com/gorilla/mux"
)

func (app *application) listBookmark(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)

	v, err := app.db.GetBookmarks(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	options, err := app.getAllOptions(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "bookmarks.page.tmpl", &templateData{
		Bookmarks: v,
		Options:   options,
		Form:      forms.New(nil),
	})
}

func (app *application) showBookmark(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)
	params := mux.Vars(r)

	id := routeUintParam(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	v, err := app.db.GetBookmark(id, user)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "show.page.tmpl", &templateData{
		Bookmark: v,
	})

}

func (app *application) editBookmarkForm(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)
	params := mux.Vars(r)
	id := routeUintParam(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	bookmark, err := app.db.GetBookmark(id, user)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	options, err := app.getOptions(bookmark, user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "edit.page.tmpl", &templateData{
		Bookmark: bookmark,
		Options:  options,
		Form:     forms.New(nil),
	})

}

func (app *application) updateBookmark(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("title")

	if !form.Valid() {
		app.render(w, r, "edit.page.tmpl", &templateData{Form: form})
		return
	}

	id := routeUintParam(form.Get("id"))
	if id < 1 {
		app.notFound(w)
		return
	}

	bookmark, err := app.db.GetBookmark(id, user)
	if err != nil {
		app.serverError(w, err)
		return
	}
	bookmark.URL = form.Get("url")

	title, excerpt, content, err := getPageContent(bookmark.URL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(form.Get("title")) > 0 {
		bookmark.Title = form.Get("title")
	} else {
		bookmark.Title = title
	}
	bookmark.Excerpt = excerpt
	bookmark.Content = content

	// Remove all the tags
	bookmark.Tags = nil
	// Add back tags that were selected
	for _, t := range r.Form["tags"] {
		tag, err := app.db.CreateIfNotExists(t)
		if err != nil {
			app.serverError(w, err)
			return
		}
		bookmark.Tags = append(bookmark.Tags, tag)
	}

	err = app.db.UpdateBookmark(bookmark)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Update successful!")

	http.Redirect(w, r, fmt.Sprintf("/app/bookmark/%d", bookmark.ID), http.StatusSeeOther)
}

func (app *application) createBookmark(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("url")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	user := app.authenticatedUser(r)

	bookmark := models.Bookmark{}
	bookmark.URL = form.Get("url")
	bookmark.UserID = user.ID

	title, excerpt, content, err := getPageContent(bookmark.URL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(form.Get("title")) > 0 {
		bookmark.Title = form.Get("title")
	} else {
		bookmark.Title = title
	}
	bookmark.Excerpt = excerpt
	bookmark.Content = content

	for _, t := range r.Form["tags"] {
		tag, err := app.db.CreateIfNotExists(t)
		if err != nil {
			app.serverError(w, err)
			return
		}
		bookmark.Tags = append(bookmark.Tags, tag)
	}

	err = app.db.InsertBookmark(&bookmark)

	if err != nil {
		app.serverError(w, err)
		return
	}

	app.session.Put(r, "flash", "Bookmark added successfully!")

	http.Redirect(w, r, fmt.Sprintf("/app/bookmark/%v", bookmark.ID), http.StatusSeeOther)
}

func (app *application) createBookmarkForm(w http.ResponseWriter, r *http.Request) {
	user := app.authenticatedUser(r)
	options, err := app.getAllOptions(user)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "create.page.tmpl", &templateData{
		Options: options,
		Form:    forms.New(r.URL.Query()),
	})
}
