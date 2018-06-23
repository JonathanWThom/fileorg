package main

import (
	"fmt"
	"log"
	"os"
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

	for _, file := range files {
		fmt.Println(file.Name())
	}
}
