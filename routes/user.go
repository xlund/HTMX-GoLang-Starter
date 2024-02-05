package routes

import (
	"go-starter/html"
	"net/http"
)

type CustomHandlerFunc func(w http.ResponseWriter, r *http.Request)

func (h CustomHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func User(w http.ResponseWriter, r *http.Request) {
	html.User(w, html.UserParams{
		IsAuthenticated: false,
		Name:            "John Doe",
		Email:           "john.doe@example.com",
	}, partial(r))
}

var UserHandler = CustomHandlerFunc(User)
