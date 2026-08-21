package routes

import (
	"net/http"

	"snip-it/internal/depend"
	"snip-it/handlers"
)

func Home(app *depend.Application, mux *http.ServeMux) {
	mux.HandleFunc("/", handlers.Home(app))
}
