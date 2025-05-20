package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
	"github.com/monstrong/proyektnaya-practica/src/pkg/handlers"
	"github.com/monstrong/proyektnaya-practica/src/pkg/render"
)

var portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache(); if err != nil {
		log.Fatal("cannot create template cache, ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// Обработчик статических файлов (CSS, JS, изображения)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/home", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/journal", handlers.Repo.Journal)
	http.HandleFunc("/team", handlers.Repo.Team)
	http.HandleFunc("/resources", handlers.Repo.Resources)
	http.HandleFunc("/", handlers.Repo.ServerPage)

	fmt.Printf("Starting application on port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
