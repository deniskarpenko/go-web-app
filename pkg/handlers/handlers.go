package handlers

import (
	"net/http"

	"github.com/deniskarpenko/go-web-app/pkg/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html")
}

// About is the about page handler
func About(w http.ResponseWriter, gor *http.Request) {
	render.RenderTemplate(w, "about.page.html")
}
