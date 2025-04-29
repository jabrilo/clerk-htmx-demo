package main

import "net/http"

func (app *app) index(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, http.StatusOK, "index.html", templateData{})
}

func (app *app) signIn(w http.ResponseWriter, r *http.Request) {
	// app.render(w, r, http.StatusOK, "index.html", content{})
	w.Write([]byte("Sign In page"))
}

func (app *app) signUp(w http.ResponseWriter, r *http.Request) {
	// app.render(w, r, http.StatusOK, "index.html", content{})
	w.Write([]byte("Sign Up page"))
}
