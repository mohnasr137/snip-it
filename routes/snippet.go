package routes

import (
	"net/http"

	"snip-it/internal/depend"
	"snip-it/handlers"
)

func Snippet(app *depend.Application, mux *http.ServeMux) {
	mux.HandleFunc("/snippet/view", handlers.SnippetView(app))
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate(app))
}