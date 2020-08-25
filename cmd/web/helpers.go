package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gorilla/csrf"
	"gitlab.com/edwlarkey/cardamom/pkg/models"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

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
	t, err := app.db.GetTags()
	if err != nil {
		return nil, err
	}

	options := []Option{}

	for _, item := range t {
		option := Option{}

		for i, bt := range bookmark.Tags {
			if item.ID == bt.ID {
				option = Option{item.Name, true}
				options = append(options, option)
				break
			}

			if i == len(bookmark.Tags)-1 {
				option = Option{item.Name, false}
				options = append(options, option)
			}
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
	var user = &models.User{}
	user, ok := val.(*models.User)
	if !ok {
		return nil
	}
	return user
}
