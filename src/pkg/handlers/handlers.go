package handlers

import (
	"net/http"

	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
	"github.com/monstrong/proyektnaya-practica/src/pkg/render"
)

// the repository used by the handlers
var Repo *Repository

// is the new repository type
type Repository struct {
	App *config.AppConfig
}

// creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl")
}

func (m *Repository) Journal(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "journal.page.tmpl")
}

func (m *Repository) Resources(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "resources.page.tmpl")
}

func (m *Repository) Team(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "team.page.tmpl")
}

func (m *Repository) ServerPage(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "server.page.tmpl")
}