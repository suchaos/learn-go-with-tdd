package blogrenderer

import (
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var aPost = Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "blog"},
	}

	t.Run("it renders the post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		got := buf.String()
		want := `<h1>hello world</h1><p>this is a description</p>Tags: <ul><li>go</li><li>blog</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
