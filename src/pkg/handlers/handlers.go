package handlers

import (
	"net/http"
	"github.com/monstrong/proyektnaya-practica/src/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func Journal(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "journal.page.tmpl")
}

func Resources(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "resources.page.tmpl")
}

func Team(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "team.page.tmpl")
}

func ServerPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "server.page.tmpl")
}