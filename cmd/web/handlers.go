package main

import (
	"fmt"
	"net/http"

	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
	"github.com/gorilla/mux"
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

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	v, err := app.db.LatestBookmarks()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Bookmarks: v,
	})
}

func (app *application) listBookmark(w http.ResponseWriter, r *http.Request) {
	v, err := app.db.GetBookmarks()
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.render(w, r, "home.page.tmpl", &templateData{
		Bookmarks: v,
	})
}

func (app *application) showBookmark(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := routeUintParam(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	v, err := app.db.GetBookmark(id)
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
	params := mux.Vars(r)
	id := routeUintParam(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	bookmark, err := app.db.GetBookmark(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	options, err := app.getOptions(bookmark)
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

	bookmark, err := app.db.GetBookmark(id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	bookmark.URL = form.Get("url")

	title, content, err := getPageContent(bookmark.URL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(form.Get("title")) > 0 {
		bookmark.Title = form.Get("title")
	} else {
		bookmark.Title = title
	}
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

	app.PrettyPrint(bookmark)

	err = app.db.UpdateBookmark(bookmark)
	if err != nil {
		app.serverError(w, err)
		return
	}

	session.AddFlash("Update successful!")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/app/bookmark/%d", bookmark.ID), http.StatusSeeOther)
}

func (app *application) createBookmarkForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) createBookmark(w http.ResponseWriter, r *http.Request) {
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
	form.Required("url")

	if !form.Valid() {
		app.render(w, r, "create.page.tmpl", &templateData{Form: form})
		return
	}

	bookmark := models.Bookmark{}
	bookmark.URL = form.Get("url")

	title, content, err := getPageContent(bookmark.URL)
	if err != nil {
		app.serverError(w, err)
		return
	}

	if len(form.Get("title")) > 0 {
		bookmark.Title = form.Get("title")
	} else {
		bookmark.Title = title
	}
	bookmark.Content = content

	err = app.db.InsertBookmark(&bookmark)

	if err != nil {
		app.serverError(w, err)
		return
	}

	session.AddFlash("Bookmark added successfully!")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/app/bookmark/%v", bookmark.ID), http.StatusSeeOther)
}

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

func (app *application) signupUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "signup.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) signupUser(w http.ResponseWriter, r *http.Request) {
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
	form.Required("name", "email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 8)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.db.InsertUser(form.Get("name"), form.Get("email"), form.Get("password"))
	if err == models.ErrDuplicateEmail {
		form.Errors.Add("email", "Address is already in use")
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	session.AddFlash("Your signup was successful. Please log in.")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *application) loginUserForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
}

func (app *application) loginUser(w http.ResponseWriter, r *http.Request) {
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

	user, err := app.db.AuthenticateUser(form.Get("email"), form.Get("password"))
	if err == models.ErrInvalidCredentials {
		form.Errors.Add("generic", "Email or Password is incorrect")
		app.render(w, r, "login.page.tmpl", &templateData{Form: form})
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	session.Values["user"] = user
	session.AddFlash("Welcome!")

	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, "/", 303)
}

func (app *application) logoutUser(w http.ResponseWriter, r *http.Request) {
	session, err := app.store.Get(r, "cardamom-session")
	if err != nil {
		app.serverError(w, err)
		return
	}
	session.Values["user"] = nil
	session.AddFlash("You've been logged out successfully!")
	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
