package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type ObsidianDoc struct {
	path        string
	filename    string
	extension   string
	attachments []string
}

func (o *ObsidianDoc) setAttachments() {
	if o.extension != "md" && o.extension != "txt" {
		return
	}
	file, err := os.Open(o.path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	re, err := regexp.Compile(`!\[{2}([^\]]+)\]]`)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		foundAttachments := re.FindAllStringSubmatch(scanner.Text(), -1)

		for _, attachment := range foundAttachments {
			o.attachments = append(o.attachments, attachment[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func (o *ObsidianDoc) deleteFile() {
	os.Remove(o.path)
}
