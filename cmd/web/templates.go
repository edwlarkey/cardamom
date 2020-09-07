package main

import (
	"html/template"
	"log"
	"sync"
	"time"

	"github.com/edwlarkey/cardamom/pkg/assets"
	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
)

type templateData struct {
	AuthenticatedUser *models.User
	CSRFField         template.HTML
	CurrentYear       int
	SiteTitle         string
	Flash             []interface{}
	Form              *forms.Form
	Bookmark          *models.Bookmark
	Bookmarks         []*models.Bookmark
	Tags              []*models.Tag
	Options           []Option
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("02 Jan 2006 at 15:04")
}

func monthDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("Jan 2006")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
	"monthDate": monthDate,
}

type templates struct {
	sync.Mutex

	base      string
	templates map[string]*template.Template
}

func initTemplates(base string) *templates {
	templateMap := make(map[string]*template.Template)
	templateFiles := make(map[string]string)

	ts, err := assets.WalkDirs("templates", false)
	if err != nil {
		log.Fatalln("Could not list templates")
	}

	for _, t := range ts {
		b, _ := assets.ReadFile(t)
		if err != nil {
			log.Fatalf("Could not read template: %s", t)
		}
		templateFiles[t] = string(b)
	}

	// Templates
	homeTemplate := template.New("home.page.tmpl").Funcs(functions)
	template.Must(homeTemplate.Parse(templateFiles["templates/home.page.tmpl"]))
	template.Must(homeTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	showTemplate := template.New("show.page.tmpl").Funcs(functions)
	template.Must(showTemplate.Parse(templateFiles["templates/show.page.tmpl"]))
	template.Must(showTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	editTemplate := template.New("edit.page.tmpl").Funcs(functions)
	template.Must(editTemplate.Parse(templateFiles["templates/edit.page.tmpl"]))
	template.Must(editTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	createTemplate := template.New("create.page.tmpl").Funcs(functions)
	template.Must(createTemplate.Parse(templateFiles["templates/create.page.tmpl"]))
	template.Must(createTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	tagListTemplate := template.New("tags.page.tmpl").Funcs(functions)
	template.Must(tagListTemplate.Parse(templateFiles["templates/tags.page.tmpl"]))
	template.Must(tagListTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	signupTemplate := template.New("signup.page.tmpl").Funcs(functions)
	template.Must(signupTemplate.Parse(templateFiles["templates/signup.page.tmpl"]))
	template.Must(signupTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	loginTemplate := template.New("login.page.tmpl").Funcs(functions)
	template.Must(loginTemplate.Parse(templateFiles["templates/login.page.tmpl"]))
	template.Must(loginTemplate.Parse(templateFiles["templates/base.layout.tmpl"]))

	templateMap["home.page.tmpl"] = homeTemplate
	templateMap["show.page.tmpl"] = showTemplate
	templateMap["edit.page.tmpl"] = editTemplate
	templateMap["create.page.tmpl"] = createTemplate
	templateMap["tags.page.tmpl"] = tagListTemplate
	templateMap["signup.page.tmpl"] = signupTemplate
	templateMap["login.page.tmpl"] = loginTemplate

	return &templates{
		base:      base,
		templates: templateMap,
	}

}
