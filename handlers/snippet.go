package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"snip-it/internal/depend"
	"snip-it/internal/helper"
)

func SnippetView(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "All snippets...")
	}
}

func SnippetCreate(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// method
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// logic
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			helper.ServerError(app, w, err)
			return
		}
		if id < 1 || id > 999 {
			helper.ServerError(app, w, err)
			return

		}

		// process
		idStr := fmt.Sprintf("%03d", id)

		// response
		fmt.Fprintf(w, "Snippet created with ID: %v", idStr)
	}
}
