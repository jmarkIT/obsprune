package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func getFiles(root string) ([]ObsidianDoc, error) {
	var files []ObsidianDoc
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			filename := strings.Split(path, "/")
			obsDoc := ObsidianDoc{
				path:     path,
				filename: filename[len(filename)-1],
			}
			files = append(files, obsDoc)
		}
		return nil
	})
	return files, err
}
