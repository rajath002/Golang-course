package handlers

import (
	"net/http"

	"github.com/rajath002/RecognizingCodeAndAddingStyles/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateDynamicCache(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplateDynamicCache(w, "about.page.tmpl")
}
