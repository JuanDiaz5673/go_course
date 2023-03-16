package handlers

import (
	"net/http"
	"github.com/JuanDiaz5673/go-course/render"
)

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "home.page.tmpl")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
    render.RenderTemplate(w, "about.page.tmpl")
}

