package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"snip-it/internal/depend"
	"snip-it/internal/dotenv"

	"snip-it/routes"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seanzhengw/fileonlyserver"
)

func main() {
	// Create loggers
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	// Load dotenv
	cfg, errCfg := dotenv.Load()
	if errCfg != nil {
		errorLog.Fatal(errCfg)
	}
	infoLog.Println("Dotenv loaded successfully")

	// Create database connection pool
	db, errDB := pgxpool.New(context.Background(), cfg["DATABASE_URL"])
	if errDB != nil {
		errorLog.Fatal(errDB)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Database connected successfully")

	// Create application
	app := &depend.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
		DB:       db,
	}

	// Create router
	mux := http.NewServeMux()

	// Serve static files
	fileServer := fileonlyserver.Serve(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Register routes
	routes.Home(app, mux)
	routes.Snippet(app, mux)

	// Create HTTP server
	srv := &http.Server{
		Addr:     cfg["PORT"],
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting on: http://localhost%s", cfg["PORT"])

	// Start server
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}