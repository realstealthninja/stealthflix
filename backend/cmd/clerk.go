package cmd

import (
	"io/fs"
	"log"
	"os"
	"strings"
)

var moviesFS fs.FS = os.DirFS("assets/movies/")

func DownloadedMovies() []string {
	files, err := fs.ReadDir(moviesFS, ".")

	var names []string

	if err != nil {
		log.Println("CLERK ERROR: " + err.Error())
	}

	for _, file := range files {
		names = append(names, file.Name())
	}

	return names
}

func MovieFileExists(name string) bool {

	files, err := fs.ReadDir(moviesFS, ".")
	if err != nil {
		return false
	}

	for _, file := range files {
		if strings.Contains(file.Name(), name) {
			return true
		}
	}
	return false
}
