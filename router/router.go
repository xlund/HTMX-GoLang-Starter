package router

import (
	"go-starter/html"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	return r
}

func home(w http.ResponseWriter, r *http.Request) {
	html.Home(w, html.HomeParams{
		Title: "Home",
	}, partial(r))
}

func partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
