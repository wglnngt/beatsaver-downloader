package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
)

const dir string = "CustomSongs"

var existing []string

func createDir() {
	path := filepath.Join(".", dir)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
}

func loadExisting() {
	path := filepath.Join(".", dir)

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			existing = append(existing, f.Name())
		}
	}

	sort.Strings(existing)
}
