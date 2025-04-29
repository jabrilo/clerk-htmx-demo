package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type app struct {
	templates map[string]*template.Template
	logger    *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	templates, err := newTemplateCache()
	if err != nil {
		logger.Info("template cache failed", "message", err.Error())
		os.Exit(1)
	}

	app := &app{
		logger:    logger,
		templates: templates,
	}

	mux := http.NewServeMux()

	mux.Handle("GET /{$}", app.logRequest(http.HandlerFunc(app.index)))
	mux.Handle("GET /sign-in", app.logRequest(http.HandlerFunc(app.signIn)))
	mux.Handle("GET /sign-up", app.logRequest(http.HandlerFunc(app.signUp)))

	server := http.Server{
		Addr:     "localhost:3000",
		ErrorLog: slog.NewLogLogger(app.logger.Handler(), slog.LevelDebug),
		Handler:  mux,
	}

	app.logger.Info("starting server...", "address", server.Addr)
	err = server.ListenAndServe()
	logger.Error("server stopped", "message", err.Error())
	os.Exit(1)
}
