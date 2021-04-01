package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	path := "/Volumes/Samsung/TheMadeKnight"
	files, err := ioutil.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}
}
