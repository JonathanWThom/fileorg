// Fileorg will organize files into subdirectories based on file extension
package main

import (
	"github.com/jonathanwthom/fileorg/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	files := openDirectory()
	filetypes := organizeByFiletype(files)
	createSubdirectories(filetypes)
}

func openDirectory() []os.FileInfo {
	dir, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	files, err := dir.Readdir(-1)
	dir.Close()
	if err != nil {
		log.Fatal(err)
	}

	return files
}

func organizeByFiletype(files []os.FileInfo) map[string][]string {
	filetypes := make(map[string][]string)

	for _, file := range files {
		name := file.Name()
		ext := filepath.Ext(name)
		if name[:1] != "." && !file.IsDir() {
			filetypes[ext] = append(filetypes[ext], name)
		}
	}

	return filetypes
}

func createSubdirectories(filetypes map[string][]string) {
	for k, v := range filetypes {
		dirname := utils.TrimLeftChar(k)
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
