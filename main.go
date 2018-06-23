package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	dirs, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dirs)
}
