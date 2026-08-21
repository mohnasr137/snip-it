package main

import (
	"net/http"

	"snip-it/internal/depend"
	"github.com/seanzhengw/fileonlyserver"
)

func routers(app *depend.Application, mux *http.ServeMux) {
	fileServer := fileonlyserver.Serve(http.Dir("./ui/static"))


	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
}
