package main

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
)

type templateData struct {
}

func (app *app) render(w http.ResponseWriter, r *http.Request, status int, template string, content templateData) {
	tmpl, ok := app.templates[template]
	if !ok {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	buffer := bytes.Buffer{}
	err := tmpl.ExecuteTemplate(&buffer, "base", content)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(status)

	buffer.WriteTo(w)
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./assets/templates/pages/*.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		tmpl, err := template.New(name).ParseFiles("./assets/templates/base.html")
		if err != nil {
			return nil, err
		}

		// add partials if any
		partials, err := filepath.Glob("./assets/templates/partials/*.html")
		if err != nil {
			return nil, err
		}

		if len(partials) > 0 {
			tmpl, err = tmpl.ParseFiles(partials...)
			if err != nil {
				return nil, err
			}
		}

		tmpl, err = tmpl.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[name] = tmpl
	}

	// htmx partials if any
	partials, err := filepath.Glob("./assets/templates/partials/htmx/*.html")
	if err != nil {
		return nil, err
	}

	if len(partials) > 0 {
		for _, partial := range partials {
			name := filepath.Base(partial)

			tmpl, err := template.New(name).ParseFiles(partial)
			if err != nil {
				return nil, err
			}

			cache[name] = tmpl
		}
	}

	return cache, nil
}
