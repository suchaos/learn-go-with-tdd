package reading_files

import (
	"bufio"
	"io"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
)

func newPost(file io.Reader) (Post, error) {
	scanner := bufio.NewScanner(file)

	readLine := func() string {
		scanner.Scan()
		return scanner.Text()
	}
	titleLine := readLine()
	descriptionLine := readLine()

	title := titleLine[len(titleSeparator):]
	description := descriptionLine[len(descriptionSeparator):]

	return Post{Title: title, Description: description}, nil
}
