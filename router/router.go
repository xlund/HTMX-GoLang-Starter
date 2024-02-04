package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", Home)
	handleStatic(r)
	return r
}

func handleStatic(r *chi.Mux) {
	prefix := "/static/"
	dir := http.Dir("html/static")
	fs := http.FileServer(dir)
	r.Handle(prefix, fs)
}

func partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
