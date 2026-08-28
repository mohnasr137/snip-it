package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"snip-it/handlers"
	"snip-it/internal/depend"
	"snip-it/internal/dotenv"
	"snip-it/internal/models"
	"snip-it/routes"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/seanzhengw/fileonlyserver"
)


func main() {
	// Create loggers.
	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	// Load dotenv.
	cfg, err := dotenv.Load()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Dotenv loaded successfully")

	// Create database connection pool.
	db, err := pgxpool.New(context.Background(), cfg["DATABASE_URL"])
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(context.Background()); err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Database connected successfully")

	// Create snippet model.
	snippetModel := &models.SnippetModel{
		DB: db,
	}

	// Create template cache.
	templateCache, err := handlers.NewTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}
	infoLog.Println("Template cache created successfully")

	// Create application.
	app := &depend.Application{
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		DB:            db,
		SnippetModel:  snippetModel,
		TemplateCache: templateCache,
	}

	// Create router.
	mux := http.NewServeMux()

	// Serve static files.
	fileServer := fileonlyserver.Serve(http.Dir("./ui/static"))
	mux.Handle(
		"/static/",
		http.StripPrefix("/static", fileServer),
	)

	// Register routes.
	routes.Home(app, mux)
	routes.Snippet(app, mux)

	// Create HTTP server.
	srv := &http.Server{
		Addr:     cfg["PORT"],
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf(
		"Starting on: http://localhost%s",
		cfg["PORT"],
	)

	// Start server.
	if err := srv.ListenAndServe(); err != nil {
		errorLog.Fatal(err)
	}
}