package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"snip-it/internal/depend"
)

func home(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./ui/html/base.html",
			"./ui/html/partials/nav.html",
			"./ui/html/pages/home.html",
		}

		tm, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Printf("Err:%v & Loc:/handeler.go/home.fn", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		err = tm.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.ErrorLog.Printf("Err:%v & Loc:/handeler.go/home.fn", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}

func snippetView(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "All snippets...")
	}
}

func snippetCreate(app *depend.Application) http.HandlerFunc {
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
			app.ErrorLog.Printf("Err:%v & Loc:/handler.go/snippetCreate.fn", err)
			http.Error(w, "Invalid ID: ID must be number", http.StatusBadRequest)
			return
		}
		if id < 1 || id > 999 {
			app.ErrorLog.Printf("Err:%v & Loc:/handler.go/snippetCreate.fn", err)
			http.Error(w, "Invalid ID: ID must be number between 0 and 1000", http.StatusBadRequest)
			return

		}

		// process
		idStr := fmt.Sprintf("%03d", id)

		// response
		fmt.Fprintf(w, "Snappet created with ID: %v", idStr)
	}
}
