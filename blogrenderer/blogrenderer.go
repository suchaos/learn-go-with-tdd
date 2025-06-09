package blogrenderer

import (
	"fmt"
	"io"
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func Render(w io.Writer, post Post) error {
	_, err := fmt.Fprintf(w, "<h1>%s</h1><p>%s</p>", post.Title, post.Description)
	if err != nil {
		return err
	}

	_, err = fmt.Fprintf(w, "Tags: <ul>")
	if err != nil {
		return err
	}

	for _, tag := range post.Tags {
		_, err := fmt.Fprintf(w, "<li>%s</li>", tag)
		if err != nil {
			return err
		}
	}

	_, err = fmt.Fprintf(w, "</ul>")
	if err != nil {
		return err
	}

	return nil
}
