package handlers

import (
	"net/http"

	"github.com/monstrong/proyektnaya-practica/src/pkg/config"
	"github.com/monstrong/proyektnaya-practica/src/pkg/models"
	"github.com/monstrong/proyektnaya-practica/src/pkg/others"
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

func (m *Repository) ServerPage(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	intMap := make(map[string]int)
	stringMap["test"] = "Hello"
	stringMap["time"] = others.CurrentTime()
	intMap["randomnum"] = others.Randomnum()

	// remoteIP := r.RemoteAddr
	// m.App.Session.Put(r.Context(), "remote_ip", remoteIP) - это в главной
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "server.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		IntMap: intMap,
	})
}


func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
	
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Journal(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "journal.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Resources(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "resources.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Team(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "team.page.tmpl", &models.TemplateData{})
}