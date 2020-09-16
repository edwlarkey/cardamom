package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	"github.com/antchfx/goreadly"
	"github.com/edwlarkey/cardamom/pkg/models"
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

func getPageContent(URL string) (title string, content template.HTML, err error) {
	p := bluemonday.UGCPolicy()
	p.AddTargetBlankToFullyQualifiedLinks(true)

	ps := bluemonday.StrictPolicy()

	/* #nosec */
	resp, err := http.Get(URL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	page, err := goreadly.ParseResponse(resp)
	if err != nil {
		return "", "", err
	}
	title = page.Title

	/* #nosec */
	content = template.HTML(p.Sanitize(page.Body))

	desc := ps.Sanitize(page.Body[0:256])
	fmt.Println(desc)

	return title, content, nil
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

func (app *application) getOptions(bookmark *models.Bookmark) ([]Option, error) {
	var options []Option
	s := make(map[uint]struct{})

	t, err := app.db.GetTags()
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
	session, err := app.store.Get(r, "cardamom-session")
	if err != nil {
		return td
	}
	td.CSRFField = csrf.TemplateField(r)
	td.AuthenticatedUser = app.authenticatedUser(r)
	td.CurrentYear = time.Now().Year()
	td.SiteTitle = "Cardamom"
	td.Flash = session.Flashes()

	err = session.Save(r, w)
	if err != nil {
		app.serverError(w, err)
		return nil
	}

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
	session, err := app.store.Get(r, "cardamom-session")
	if err != nil {
		app.errorLog.Println("Could not retrieve session")
		return nil
	}

	val := session.Values["user"]
	var user *models.User
	user, ok := val.(*models.User)
	if !ok {
		return nil
	}
	return user
}
