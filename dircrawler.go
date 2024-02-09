package main

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func getFiles(root string) ([]ObsidianDoc, error) {
	var files []ObsidianDoc
	err := filepath.WalkDir(root, func(path string, info fs.DirEntry, err error) error {
		if !info.IsDir() {
			filenameSlice := strings.Split(path, string(os.PathSeparator))
			filename := filenameSlice[len(filenameSlice)-1]
			extensionSlice := strings.Split(path, ".")
			extension := extensionSlice[len(extensionSlice)-1]
			obsDoc := ObsidianDoc{
				path:      path,
				filename:  filename,
				extension: extension,
			}
			files = append(files, obsDoc)
		}
		return nil
	})
	return files, err
}
