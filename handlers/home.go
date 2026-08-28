package handlers

import (
	"net/http"

	"snip-it/internal/depend"
	"snip-it/internal/helper"
)

func Home(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		// Get latest snippets.
		snippets, err := app.SnippetModel.Latest(r.Context())
		if err != nil {
			helper.ServerError(app, w, err)
			return
		}

		// Create template data with common data.
		data := NewTemplateData()

		// Add home-specific data.
		data.Snippets = snippets

		// Render cached template.
		Render(
			app,
			w,
			http.StatusOK,
			"home.html",
			data,
		)
	}
}