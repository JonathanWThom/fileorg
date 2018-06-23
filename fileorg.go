package main

import (
	"fmt"
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

	/// create a map
	/// key is file extension
	/// value is array of filepaths
	/// create a folder for all of those
	/// move all the files there
	/// don't run while in this directory or it'll mess up this program
	filetypes := make(map[string][]string)
	for _, file := range files {
		name := file.Name()
		ext := filepath.Ext(name)
		filetypes[ext] = append(filetypes[ext], name)
	}

	for k, _ := range filetypes {
		dirname := trimLeftChar(k)
		os.Mkdir(dirname, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Printf("%#v", filetypes)
}

func trimLeftChar(s string) string {
	for i := range s {
		if i > 0 {
			return s[i:]
		}
	}
	return s[:0]
}

// TODO: Handle directories and files without extensions
