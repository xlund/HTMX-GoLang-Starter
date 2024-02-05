package html

import (
	"embed"
	"html/template"
	"io"
)

//go:embed *
var files embed.FS

func parse(file string) *template.Template {
	return template.Must(template.New("layout.html").ParseFS(files, "layout.html", "sidebar.html", file))
}

type HomeParams struct {
	Title string
}

func Home(w io.Writer, p HomeParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return parse("index.html").ExecuteTemplate(w, partial, p)
}

type UserParams struct {
	IsAuthenticated bool
	Name            string
	Email           string
}

func User(w io.Writer, p UserParams, partial string) error {
	if partial == "" {
		partial = "layout.html"
	}
	return parse("user.html").ExecuteTemplate(w, partial, p)
}
