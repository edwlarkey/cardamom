package main

import (
	"html/template"
	"io/fs"
	"path/filepath"
	"time"

	"git.sr.ht/~edwlarkey/cardamom/internal/models"
	"git.sr.ht/~edwlarkey/cardamom/ui"
)

type templateData struct {
	CurrentYear     int
	Bookmark        *models.Bookmark
	BookmarkContent template.HTML
	Bookmarks       []*models.Bookmark
	Tag             *models.Tag
	Tags            []*models.Tag
	Form            any
	Flash           string
	IsAuthenticated bool
	CSRFToken       string
}

func humanDate(t time.Time) string {
	if t.IsZero() {
		return ""
	}

	return t.UTC().Format("2006-01-02")
}

var functions = template.FuncMap{
	"humanDate": humanDate,
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := fs.Glob(ui.Files, "html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			"html/partials/*.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	patterns := []string{
		"html/partials/*.tmpl",
	}

	ts, err := template.New("partials").Funcs(functions).ParseFS(ui.Files, patterns...)
	if err != nil {
		return nil, err
	}

	cache["partials"] = ts

	return cache, nil
}
