package depend

import (
	"log"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Application struct {
	InfoLog *log.Logger
	ErrorLog *log.Logger
	DB *pgxpool.Pool
}