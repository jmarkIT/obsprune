package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	err := filepath.WalkDir("/Users/james/Documents/Notes", visit)
	fmt.Println(err)
}
