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

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	user_id := app.authenticatedUser(r)
	bookmarks, err := app.bookmarks.Latest(user_id)
	if err != nil {
		app.serverError(w, err)
		return
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

	user_id := app.authenticatedUser(r)

	bookmark, err := app.bookmarks.Get(user_id, id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := app.newTemplateData(r)
	data.Bookmark = bookmark

	app.render(w, http.StatusOK, "view.tmpl", data)
}

type bookmarkCreateForm struct {
	Title               string `form:"title"`
	Name                string `form:"name"`
	Url                 string `form:"url"`
	validator.Validator `form:"-"`
}

func (app *application) bookmarkCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = bookmarkCreateForm{}

	if r.Header.Get("HX-Request") == "true" {
		user_id := app.authenticatedUser(r)
		tags, err := app.tags.All(user_id)
		if err != nil {
			app.serverError(w, err)
			return
		}
		data.Tags = tags

		app.renderPartial(w, http.StatusOK, "create_bookmark.tmpl", data)
	} else {
		app.render(w, http.StatusOK, "create.tmpl", data)
	}
}

func (app *application) bookmarkCreatePost(w http.ResponseWriter, r *http.Request) {
	var form bookmarkCreateForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	// form.CheckField(validator.NotBlank(form.Url), "url", "This field cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		app.render(w, http.StatusUnprocessableEntity, "create_bookmark.tmpl", data)
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

	id, err := app.bookmarks.Insert(title, form.Url, content, userId)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Bookmark successfully created!")

	http.Redirect(w, r, fmt.Sprintf("/bookmark/view/%d", id), http.StatusSeeOther)
}

type tagCreateForm struct {
	Name                string `form:"name"`
	validator.Validator `form:"-"`
}

func (app *application) tagCreate(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	data.Form = tagCreateForm{}

	app.renderPartial(w, http.StatusOK, "create_tag.tmpl", data)
}

func (app *application) tagGetHtmx(w http.ResponseWriter, r *http.Request) {
	user_id := app.authenticatedUser(r)
	data := app.newTemplateData(r)
	tags, err := app.tags.All(user_id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Tags = tags

	app.renderPartial(w, http.StatusOK, "tag_options.tmpl", data)
}

func (app *application) tagCreatePost(w http.ResponseWriter, r *http.Request) {
	var form bookmarkCreateForm

	user_id := app.authenticatedUser(r)
	data := app.newTemplateData(r)
	tags, err := app.tags.All(user_id)
	if err != nil {
		app.serverError(w, err)
		return
	}
	data.Tags = tags

	err = app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Name, 100), "title", "This field cannot be more than 100 characters long")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		app.renderPartial(w, http.StatusUnprocessableEntity, "create_bookmark.tmpl", data)
		return
	}

	_, err = app.tags.Insert(form.Name, user_id)
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Tag successfully created!")

	app.renderPartial(w, http.StatusOK, "create_bookmark.tmpl", data)
}

type userSignupForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userSignupForm{}
	app.render(w, http.StatusOK, "signup.tmpl", data)
}

func (app *application) userSignupPost(w http.ResponseWriter, r *http.Request) {
	var form userSignupForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Name), "name", "This field cannot be blank")
	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "This field must be at least 8 characters long")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already in use")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "signup.tmpl", data)
		} else {
			app.serverError(w, err)
		}

		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your signup was successful. Please log in.")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)
}

type userLoginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = userLoginForm{}
	app.render(w, http.StatusOK, "login.tmpl", data)
}

func (app *application) userLoginPost(w http.ResponseWriter, r *http.Request) {
	var form userLoginForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Email), "email", "This field cannot be blank")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "This field must be a valid email address")
	form.CheckField(validator.NotBlank(form.Password), "password", "This field cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email or password is incorrect")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, http.StatusUnprocessableEntity, "login.tmpl", data)
		} else {
			app.serverError(w, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *application) userLogoutPost(w http.ResponseWriter, r *http.Request) {
	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, err)
		return
	}

	app.sessionManager.Remove(r.Context(), "authenticatedUserID")

	app.sessionManager.Put(r.Context(), "flash", "You've been logged out successfully!")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
