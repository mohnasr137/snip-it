package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"snip-it/internal/depend"
	"snip-it/internal/helper"
)

// NewTemplateData returns a TemplateData struct containing
// the common dynamic data used across templates.
func NewTemplateData() *TemplateData {
	return &TemplateData{
		CurrentYear: time.Now().Year(),
	}
}

// Render renders a cached template safely.
func Render(
	app *depend.Application,
	w http.ResponseWriter,
	status int,
	page string,
	data *TemplateData,
) {
	// Get the template from the cache.
	ts, ok := app.TemplateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		helper.ServerError(app, w, err)
		return
	}

	// Render the template into a buffer first.
	var buf bytes.Buffer

	if err := ts.ExecuteTemplate(&buf, "base", data); err != nil {
		helper.ServerError(app, w, err)
		return
	}

	// Only send the status code after successful rendering.
	w.WriteHeader(status)

	// Write the rendered template to the response.
	if _, err := buf.WriteTo(w); err != nil {
		app.ErrorLog.Println(err)
	}
}