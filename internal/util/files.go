package util

import (
	"log"
	"os"
)

func FilesFromDir(folder string) []string {
	files := []string{}
	dir, err := os.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range dir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files
}
