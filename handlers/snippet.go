package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"snip-it/internal/depend"
	"snip-it/internal/helper"
	"snip-it/internal/models"
)

func SnippetView(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Get snippet ID from URL.
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}

		// Get snippet from database.
		snippet, err := app.SnippetModel.Get(
			r.Context(),
			id,
		)
		if err != nil {
			if errors.Is(err, models.ErrNoRecord) {
				http.NotFound(w, r)
			} else {
				helper.ServerError(app, w, err)
			}
			return
		}

		// Create template data with common data.
		data := NewTemplateData()

		// Add snippet-specific data.
		data.Snippet = snippet

		// Render cached template.
		Render(
			app,
			w,
			http.StatusOK,
			"view.html",
			data,
		)
	}
}

func SnippetCreate(app *depend.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// Method.
		if r.Method != http.MethodPost {
			w.Header().Set("Allow", http.MethodPost)
			http.Error(
				w,
				"Method not allowed",
				http.StatusMethodNotAllowed,
			)
			return
		}

		// Decode request body.
		var data map[string]json.RawMessage

		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(
				w,
				"Invalid request body",
				http.StatusBadRequest,
			)
			return
		}

		// Check required fields.
		if data["title"] == nil {
			http.Error(
				w,
				"Title is required",
				http.StatusBadRequest,
			)
			return
		}

		if data["content"] == nil {
			http.Error(
				w,
				"Content is required",
				http.StatusBadRequest,
			)
			return
		}

		if data["expires"] == nil {
			http.Error(
				w,
				"Expires is required",
				http.StatusBadRequest,
			)
			return
		}

		// Create variables.
		var title string
		var content string
		var expires int

		// Decode title.
		if err := json.Unmarshal(data["title"], &title); err != nil {
			http.Error(
				w,
				"Invalid title",
				http.StatusBadRequest,
			)
			return
		}

		// Decode content.
		if err := json.Unmarshal(data["content"], &content); err != nil {
			http.Error(
				w,
				"Invalid content",
				http.StatusBadRequest,
			)
			return
		}

		// Decode expires.
		if err := json.Unmarshal(data["expires"], &expires); err != nil {
			http.Error(
				w,
				"Invalid expires",
				http.StatusBadRequest,
			)
			return
		}

		// Validate input.
		if title == "" {
			http.Error(
				w,
				"Title is required",
				http.StatusBadRequest,
			)
			return
		}

		if content == "" {
			http.Error(
				w,
				"Content is required",
				http.StatusBadRequest,
			)
			return
		}

		if expires < 1 {
			http.Error(
				w,
				"Expires must be greater than 0",
				http.StatusBadRequest,
			)
			return
		}

		// Insert snippet.
		id, err := app.SnippetModel.Insert(
			r.Context(),
			title,
			content,
			expires,
		)
		if err != nil {
			helper.ServerError(app, w, err)
			return
		}

		// Response.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(map[string]any{
			"message": "Snippet created successfully",
			"id":      id,
		}); err != nil {
			app.ErrorLog.Println(err)
		}
	}
}