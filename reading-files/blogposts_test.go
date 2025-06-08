package reading_files

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"
)

type StubFailingFS struct {
}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, I'm always failing")
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("it returns a list of posts", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hello.md":  {Data: []byte("hello world")},
			"hello2.md": {Data: []byte("hello world2")},
		}
		posts, err := NewPostsFromFS(fileSystem)
		if err != nil {
			t.Fatal(err)
		}
		if len(posts) != len(fileSystem) {
			t.Errorf("got %d posts, want %d posts", len(posts), len(fileSystem))
		}
	})

	t.Run("it returns an error when the file cannot be opened", func(t *testing.T) {
		stubFailingFS := StubFailingFS{}
		_, err := NewPostsFromFS(stubFailingFS)
		if err == nil {
			t.Error("did not get an error")
		}
	})

	t.Run("test title", func(t *testing.T) {
		fileSystem := fstest.MapFS{
			"hello.md":  {Data: []byte("Title: Post 1")},
			"hello2.md": {Data: []byte("Title: Post 1")},
		}
		posts, _ := NewPostsFromFS(fileSystem)

		assertPost(t, posts[0], Post{Title: "Post 1"})
	})

	t.Run("test content", func(t *testing.T) {
		const (
			firstBody = `Title: Post 1
Description: Description 1`
			secondBody = `Title: Post 2
Description: Description 2`
		)

		fileSystem := fstest.MapFS{
			"hello.md":  {Data: []byte(firstBody)},
			"hello2.md": {Data: []byte(secondBody)},
		}
		posts, _ := NewPostsFromFS(fileSystem)

		assertPost(t, posts[0], Post{Title: "Post 1", Description: "Description 1"})
	})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
