package main

import (
	"log"
	"net/http"
	"os"

	"snip-it/internal/depend"
	"snip-it/internal/dotenv"

	"github.com/seanzhengw/fileonlyserver"
)

func main() {
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	app := &depend.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}
	cfg, errCfg := dotenv.Load()
	if errCfg != nil {
		errorLog.Fatal(errCfg)
	}

	mux := http.NewServeMux()

	fileServer := fileonlyserver.Serve(http.Dir("./ui/static"))

	mux.HandleFunc("/", home(app))
	mux.HandleFunc("/snippet/view", snippetView(app))
	mux.HandleFunc("/snippet/create", snippetCreate(app))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     cfg["PORT"],
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on: http://localhost%s", cfg["PORT"])
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
