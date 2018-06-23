package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		log.Fatal(err)
	}

	/// move all the files there
	filetypes := make(map[string][]string)
	for _, file := range files {
		name := file.Name()
		ext := filepath.Ext(name)
		if name[:1] != "." && !file.IsDir() {
			filetypes[ext] = append(filetypes[ext], name)
		}
	}

	for k, v := range filetypes {
		dirname := trimLeftChar(k)
		if dirname != "" {
			err := os.Mkdir(dirname, 0777)
			if err != nil {
				log.Fatal(err)
			}

			for _, file := range v {
				err := os.Rename(file, dirname+"/"+file)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}
}

// Could move this to a utils file
func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

// TODO: Handle directories and files without extensions
// TODO: Tests
// TODO: refactor into functions
