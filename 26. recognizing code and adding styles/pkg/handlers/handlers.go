package handlers

import (
	"net/http"

	"github.com/rajath002/RecognizingCodeAndAddingStyles/pkg/config"
	"github.com/rajath002/RecognizingCodeAndAddingStyles/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New Handlers : sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateDynamicCache(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateDynamicCache(w, "about.page.tmpl")
}
