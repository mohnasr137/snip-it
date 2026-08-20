package main

import (
	"log"
	"net/http"

	"snip-it/internal/dotenv"

	"github.com/seanzhengw/fileonlyserver"
)

func main() {
	mux := http.NewServeMux()
	cfg := dotenv.Load()

	fileServer := fileonlyserver.Serve(http.Dir("./ui/static"))

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on : http://localhost%s", cfg.Port)
	err := http.ListenAndServe(cfg.Port, mux)
	log.Fatal(err)
}
