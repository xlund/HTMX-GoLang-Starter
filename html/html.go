package html

import (
	"embed"
	"html/template"
	"io"
	"todo-list/domain/todo"
)

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	return template.Must(template.New("layout.html").ParseFS(files, "layout.html", file))
}

type TodoParams struct {
	PageTitle string
	Todos     []todo.Todo
}

func Todo(w io.Writer, p TodoParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return parse("index.html").ExecuteTemplate(w, partial, p)
}

func Home(w io.Writer, p TodoParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return parse("home.html").ExecuteTemplate(w, partial, p)
}
