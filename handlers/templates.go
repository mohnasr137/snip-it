package handlers

import (
	"html/template"
	"path/filepath"
	"time"

	"snip-it/internal/models"
)

// TemplateData acts as the holding structure for any dynamic data
// that we want to pass to our HTML templates.
type TemplateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
}

// humanDate returns a nicely formatted date and time.
func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

// Template functions available inside HTML templates.
var functions = template.FuncMap{
	"humanDate": humanDate,
}

// NewTemplateCache parses all HTML templates and stores them
// in an in-memory cache.
func NewTemplateCache() (map[string]*template.Template, error) {
	cache := make(map[string]*template.Template)

	// Find all page templates.
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// Create a new template and register custom functions
		// before parsing any templates.
		ts, err := template.
			New(name).
			Funcs(functions).
			ParseFiles("./ui/html/base.html")
		if err != nil {
			return nil, err
		}

		// Parse all partial templates.
		ts, err = ts.ParseGlob("./ui/html/partials/*.html")
		if err != nil {
			return nil, err
		}

		// Parse the page template.
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// Store the complete template set in the cache.
		cache[name] = ts
	}

	return cache, nil
}