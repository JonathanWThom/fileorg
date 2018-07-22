// Fileorg will organize files into subdirectories based on file extension
package main

import (
	"flag"
	"fmt"
	"github.com/jonathanwthom/fileorg/utils"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dateFlag := flag.Bool("date", false, "Subdivide by date")
	flag.Parse()
	files := openDirectory(".")
	filetypes := organizeByFiletype(files)
	dirs := createSubdirectories(filetypes)
	if *dateFlag == true {
		createSubdirectoriesByDate(dirs)
	}
}

func openDirectory(path string) []os.FileInfo {
	dir, err := os.Open(path)
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

func createSubdirectories(filetypes map[string][]string) []string {
	var dirs []string
	for k, v := range filetypes {
		dirname := utils.TrimLeftChar(k)
		if dirname != "" {
			err := os.Mkdir(dirname, 0777)
			dirs = append(dirs, dirname)
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
	return dirs
}

func createSubdirectoriesByDate(dirs []string) {
	for _, dir := range dirs {
		files := openDirectory(dir)
		organizeByDate(dir, files)
	}
}

func organizeByDate(dir string, files []os.FileInfo) {
	for _, file := range files {
		modifiedAt := file.ModTime()
		month := modifiedAt.Month()
		year := modifiedAt.Year()
		dirname := fmt.Sprintf("%v/%v_%v", dir, month, year)

		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			err := os.Mkdir(dirname, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}

		err := os.Rename(dir+"/"+file.Name(), dirname+"/"+file.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
}
