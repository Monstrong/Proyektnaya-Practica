package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
	"github.com/monstrong/proyektnaya-practica/src/pkg/handlers"
	"github.com/monstrong/proyektnaya-practica/src/pkg/render"
)

var portNumber = ":8080"
var session *scs.SessionManager
var app config.AppConfig

func main() {

	// изменить этот параметр на true когда сервер работает не на локал хосте
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache, ", err)
	}

	app.TemplateCache = tc
	app.UseCache = true

	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Starting application on port %s\n", portNumber)
	// _ = http.ListenAndServe(portNumber, nil)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}		
}
