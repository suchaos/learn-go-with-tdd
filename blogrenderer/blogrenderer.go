package blogrenderer

import (
	"embed"
	"html/template"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

type PostRenderer struct {
	template *template.Template
}

var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func NewPostRenderer() (*PostRenderer, error) {
	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}

	return &PostRenderer{template: templ}, nil
}

func (r *PostRenderer) Render(w io.Writer, post Post) error {
	err := r.template.ExecuteTemplate(w, "blog.gohtml", post)
	if err != nil {
		return err
	}

	return nil
}
