package handlers

import (
	"net/http"

	"github.com/yasniel1408/go-course/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "home.page.html")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderPage(w, "about.page.html")
}
