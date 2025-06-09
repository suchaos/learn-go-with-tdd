package blogrenderer

import (
	"bytes"
	"github.com/approvals/go-approval-tests"
	"io"
	"testing"
)

func TestRender(t *testing.T) {
	var aPost = Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "blog"},
	}

	postRenderer, err := NewPostRenderer()

	if err != nil {
		t.Fatal(err)
	}

	t.Run("it renders the post to HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := postRenderer.Render(&buf, aPost)

		if err != nil {
			t.Fatal(err)
		}

		approvals.VerifyString(t, buf.String())
	})
}

func BenchmarkRender(b *testing.B) {
	var aPost = Post{
		Title:       "hello world",
		Body:        "this is a post",
		Description: "this is a description",
		Tags:        []string{"go", "blog"},
	}
	postRenderer, err := NewPostRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		postRenderer.Render(io.Discard, aPost)
	}
}
