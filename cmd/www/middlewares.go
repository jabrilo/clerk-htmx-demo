package main

import "net/http"

func (app *app) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.Info("new request incoming", "method", r.Method, "url", r.URL.String(), "ip", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
