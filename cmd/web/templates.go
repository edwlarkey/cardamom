package main

import (
	"html/template"
	"net/url"
	"path/filepath"
	"sync"
	"time"

	"github.com/edwlarkey/cardamom/pkg/assets"
	"github.com/edwlarkey/cardamom/pkg/forms"
	"github.com/edwlarkey/cardamom/pkg/models"
	"github.com/shurcooL/httpfs/html/vfstemplate"
	"github.com/shurcooL/httpfs/path/vfspath"
)

type templateData struct {
	AuthenticatedUser *models.User
	CSRFField         template.HTML
	CurrentYear       int
	SiteTitle         string
	Hostname          string
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

func domainName(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}

	return u.Host
}

func monthDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format("Jan 2006")
}

var functions = template.FuncMap{
	"humanDate":  humanDate,
	"monthDate":  monthDate,
	"domainName": domainName,
}

type templates struct {
	sync.Mutex

	base      string
	templates map[string]*template.Template
}

func initTemplates() (*templates, error) {
	cache := map[string]*template.Template{}

	pages, err := vfspath.Glob(assets.Assets, "/templates/*.page.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts := template.New(name).Funcs(functions)

		ts, err = vfstemplate.ParseFiles(assets.Assets, ts, page)
		if err != nil {
			return nil, err
		}

		ts, err = vfstemplate.ParseGlob(assets.Assets, ts, "/templates/*.layout.tmpl")
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return &templates{
		base:      "base",
		templates: cache,
	}, nil

}
