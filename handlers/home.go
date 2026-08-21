package handlers

import (
	"html/template"
	"net/http"

	"snip-it/internal/depend"
	"snip-it/internal/helper"
)

func Home(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			helper.NotFound(w)
			return
		}

		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			"./ui/html/pages/home.html",
		}

		tm, err := template.ParseFiles(files...)
		if err != nil {
			helper.ServerError(app, w, err)
			return
		}

		err = tm.ExecuteTemplate(w, "base", nil)
		if err != nil {
			helper.ServerError(app, w, err)
			return
		}
	}
}
