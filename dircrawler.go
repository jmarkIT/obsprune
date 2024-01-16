package main

import (
	"io/fs"
	"path/filepath"
)

func getFiles(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}
