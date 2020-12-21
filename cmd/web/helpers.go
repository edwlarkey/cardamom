package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/edwlarkey/cardamom/pkg/models"
	readability "github.com/go-shiori/go-readability"
	"github.com/gorilla/csrf"
	"github.com/microcosm-cc/bluemonday"
)

// routeUintParam returns an URL route parameter as uint
func routeUintParam(param string) uint {
	value, err := strconv.ParseUint(param, 0, 0)
	if err != nil {
		return 0
	}

	return uint(value)
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

	page, err := readability.FromReader(resp.Body, Url)
	if err != nil {
		return "", "", "", err
	}

	title = page.Title

	/* #nosec */
	content = template.HTML(p.Sanitize(page.Content))

	return title, page.Excerpt, content, nil
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	_ = app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		app.infoLog.Println(string(b))
	}
	return
}

func (app *application) getOptions(bookmark *models.Bookmark, user *models.User) ([]Option, error) {
	var options []Option
	s := make(map[uint]struct{})

	t, err := app.db.GetTags(user)
	if err != nil {
		return nil, err
	}

	for _, bt := range bookmark.Tags {
		s[bt.ID] = struct{}{}
	}

	for _, item := range t {
		var option Option

		_, ok := s[item.ID]
		if ok {
			option = Option{item.Name, true}
			options = append(options, option)
		} else {
			option = Option{item.Name, false}
			options = append(options, option)
		}

	}

	return options, nil
}

func (app *application) getAllOptions(user *models.User) ([]Option, error) {
	var options []Option

	t, err := app.db.GetTags(user)
	if err != nil {
		return nil, err
	}

	for _, item := range t {
		option := Option{item.Name, false}
		options = append(options, option)

	}

	return options, nil
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) addDefaultData(td *templateData, w http.ResponseWriter, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}

	td.CSRFField = csrf.TemplateField(r)
	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CurrentYear = time.Now().Year()
	td.SiteTitle = "Cardamom"
	td.Flash = app.session.PopString(r, "flash")

	return td
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	app.templates.Lock()
	defer app.templates.Unlock()

	template, ok := app.templates.templates[name]
	if !ok {
		app.errorLog.Printf("template %s not found", name)
	}

	err := template.ExecuteTemplate(w, app.templates.base, app.addDefaultData(td, w, r))
	if err != nil {
		app.errorLog.Printf("error parsing template %s: %s", name, err)
		app.serverError(w, err)
	}
}

func (app *application) authenticatedUser(r *http.Request) *models.User {
	user, ok := r.Context().Value(contextKeyUser).(*models.User)
	if !ok {
		return nil
	}
	return user
}
