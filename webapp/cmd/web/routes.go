package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	mux := chi.NewRouter()

	// register middleware
	mux.Use(middleware.Recoverer)
	mux.Use(app.addIPToContext)

	// register routers
	mux.Get("/", app.Home) // home page

	// statics assets
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
