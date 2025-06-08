package reading_files

import (
	"io/fs"
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var result []Post
	for _, entry := range dir {
		post, err := getPost(fileSystem, entry)
		if err != nil {
			return nil, err
		}
		result = append(result, post)
	}
	return result, nil
}

func getPost(fileSystem fs.FS, entry fs.DirEntry) (Post, error) {
	file, err := fileSystem.Open(entry.Name())
	if err != nil {
		return Post{}, err
	}
	defer file.Close()
	return newPost(file)
}
