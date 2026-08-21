package helper

import (
	"net/http"

	"snip-it/internal/depend"
)

func ServerError(app *depend.Application, w http.ResponseWriter, err error) {
	app.ErrorLog.Printf("Err: %v", err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func NotFound(w http.ResponseWriter) {
	ClientError(w, http.StatusNotFound)
}
