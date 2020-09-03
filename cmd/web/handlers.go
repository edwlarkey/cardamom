package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/antchfx/goreadly"
	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/bluemonday"
)

type Option struct {
	Value    string
	Selected bool
}

// RouteInt64Param returns an URL route parameter as int64.
func RouteInt64Param(param string) int64 {
	value, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return 0
	}

	if value < 0 {
		return 0
	}

	return value
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

func (app *application) showBookmark(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := RouteInt64Param(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	v, err := app.db.BookmarkByID(id)
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
	id := RouteInt64Param(params["id"])
	if id < 1 {
		app.notFound(w)
		return
	}

	bookmark, err := app.db.BookmarkByID(id)
	if err == models.ErrNoRecord {
		app.notFound(w)
		return
	} else if err != nil {
		app.serverError(w, err)
		return
	}

	app.PrettyPrint(bookmark)

	options, err := app.getOptions(bookmark)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.PrettyPrint(options)

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

	id := RouteInt64Param(form.Get("id"))
	if id < 1 {
		app.notFound(w)
		return
	}

	bookmark, err := app.db.BookmarkByID(id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	bookmark.Title = form.Get("title")
	bookmark.URL = form.Get("url")

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

	if len(form.Get("title")) > 0 {
		bookmark.Title = form.Get("title")
	} else {
		p := bluemonday.UGCPolicy()
		p.AddTargetBlankToFullyQualifiedLinks(true)

		resp, err := http.Get(bookmark.URL)
		if err != nil {
			app.serverError(w, err)
			return
		}
		defer resp.Body.Close()

		page, err := goreadly.ParseResponse(resp)
		if err != nil {
			app.serverError(w, err)
			return
		}

		bookmark.Title = page.Title
	}

	app.PrettyPrint(&bookmark)

	err = app.db.InsertBookmark(&bookmark)

	app.PrettyPrint(&bookmark)

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

func (app *application) createTagForm(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "create_tag.page.tmpl", &templateData{
		Form: forms.New(nil),
	})
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

	id, err := app.db.InsertTag(form.Get("name"))
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

	http.Redirect(w, r, fmt.Sprintf("/app/tag/%d", id), http.StatusSeeOther)
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
	form.Required("email", "password")
	form.MatchesPattern("email", forms.EmailRX)
	form.MinLength("password", 8)

	if !form.Valid() {
		app.render(w, r, "signup.page.tmpl", &templateData{Form: form})
		return
	}

	err = app.db.InsertUser(form.Get("email"), form.Get("password"))
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

	app.PrettyPrint(user)

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
