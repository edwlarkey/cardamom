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

	return t.UTC().Format("02 Jan 2006 at 15:04")
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

	partials, err := fs.Glob(ui.Files, "html/partials/*.tmpl")
	if err != nil {
		return nil, err
	}

	htmx, err := fs.Glob(ui.Files, "html/htmx/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		patterns := []string{
			"html/base.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	for _, part := range partials {
		name := filepath.Base(part)

		patterns := []string{
			part,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	for _, h := range htmx {
		name := filepath.Base(h)

		patterns := []string{
			h,
		}

		ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
		if err != nil {
			return nil, err
		}

		cache[name] = ts
	}

	return cache, nil
}
