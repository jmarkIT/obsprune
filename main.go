package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	usage := `usage: obspruner path_to_obsidian_vault path_to_obsidian_vault_attachments`
	if len(os.Args) != 3 {
		fmt.Println("Error: Incorrect positional arguments")
		fmt.Println(usage)
		os.Exit(1)
	}
	obsidianRoot := os.Args[1]
	attachmentsDir := os.Args[2]
	fmt.Printf("%s, %s", obsidianRoot, attachmentsDir)

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
		if !slices.Contains(attSlice, attachment.filename) {
			fmt.Printf("Are you sure you want to delete\n")
			fmt.Printf("%s\n> ", attachment.path)
			var confirmation string
			fmt.Scanln(&confirmation)
			if strings.ToLower(confirmation) == "y" {
				attachment.deleteFile()
			}
		}
	}

}
