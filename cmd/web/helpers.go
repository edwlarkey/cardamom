package main

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	nurl "net/url"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
	readability "github.com/go-shiori/go-readability"
	"github.com/justinas/nosurf"
	"github.com/microcosm-cc/bluemonday"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear:     time.Now().Year(),
		Flash:           app.sessionManager.PopString(r.Context(), "flash"),
		IsAuthenticated: app.isAuthenticated(r),
		CSRFToken:       nosurf.Token(r),
	}
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)

	buf.WriteTo(w)
}

func (app *application) decodePostForm(r *http.Request, dst any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dst, r.PostForm)
	if err != nil {

		var invalidDecoderError *form.InvalidDecoderError

		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}

func (app *application) isAuthenticated(r *http.Request) bool {
	isAuthenticated, ok := r.Context().Value(isAuthenticatedContextKey).(bool)
	if !ok {
		return false
	}

	return isAuthenticated
}

func (app *application) authenticatedUser(r *http.Request) int {
	user_id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
	return user_id
}

func getPageContent(Url string) (title string, excerpt string, content template.HTML, err error) {
	p := bluemonday.UGCPolicy()
	p.AddTargetBlankToFullyQualifiedLinks(true)

	/* #nosec */
	resp, err := http.Get(Url)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	u, err := nurl.Parse(Url)
	if err != nil {
		return "", "", "", err
	}

	page, err := readability.FromReader(resp.Body, u)
	if err != nil {
		return "", "", "", err
	}

	title = page.Title

	/* #nosec */
	content = template.HTML(p.Sanitize(page.Content))

	return title, page.Excerpt, content, nil
}
