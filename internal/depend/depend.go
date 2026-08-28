package depend

import (
	"html/template"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"snip-it/internal/models"
)

type Application struct {
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	DB            *pgxpool.Pool
	SnippetModel  *models.SnippetModel
	TemplateCache map[string]*template.Template
}