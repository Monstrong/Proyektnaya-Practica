package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
	"github.com/monstrong/proyektnaya-practica/src/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.ServerPage))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/journal", http.HandlerFunc(handlers.Repo.Journal))
	mux.Get("/team", http.HandlerFunc(handlers.Repo.Team))
	mux.Get("/resources", http.HandlerFunc(handlers.Repo.Resources))

	return mux
}
