package main

import (
	"myapp/pkg/config"
	"myapp/pkg/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	//	mux := pat.New()

	//	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	//	mux.Get("/about", http.HandlerFunc(handlers.Repo.Home))

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux

}
