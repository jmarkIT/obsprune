package main

import (
	"fmt"
	"slices"
)

func main() {
	allFiles, err := getFiles("/Users/james/Documents/Notes")
	if err != nil {
		fmt.Println(err)
	}
	var attSlice []string
	for _, file := range allFiles {
		file.setAttachments()
		if file.attachments != nil {
			attSlice = append(attSlice, file.attachments...)
		}
	}
	allAttachments, err := getFiles("/Users/james/Documents/Notes/03-RESOURCES 📚/Attachments")
	if err != nil {
		fmt.Println(err)
	}
	for _, attachment := range allAttachments {
		if slices.Contains(attSlice, attachment.filename) {
			fmt.Println(attachment.path)
		}
	}

}
