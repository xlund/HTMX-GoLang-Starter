package main

import (
	"net/http"
	"todo-list/domain/todo"
	"todo-list/html"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	println("Server is running on port 8000")
	http.ListenAndServe("localhost:8000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	todoParams := html.TodoParams{
		PageTitle: "Todo List",
		Todos: []todo.Todo{
			*todo.New(todo.TodoTitle("Buy milk")),
			*todo.New(todo.TodoTitle("Send email")),
		},
	}
	html.Todo(w, todoParams, partial(r))
}

func partial(r *http.Request) string {
	return r.URL.Query().Get("partial")
}
