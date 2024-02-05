package routes

import (
	"go-starter/html"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	html.Home(w, html.HomeParams{
		Title: "Home",
	}, partial(r))
}
