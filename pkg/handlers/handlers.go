package handlers

import (
	"net/http"

	"github.com/deniskarpenko/go-web-app/pkg/config"
	"github.com/deniskarpenko/go-web-app/pkg/models"
	"github.com/deniskarpenko/go-web-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, gor *http.Request) {
	//perform some logic
	stringMap := make((map[string]string))
	stringMap["test"] = "Hello, again"

	//send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
