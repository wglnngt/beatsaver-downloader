package main

import (
	"fmt"
)

func main() {
	createDir()
	loadExisting()

	downloadAll()
}

func downloadAll() {
	start := 0
	done := false

	for !done {
		songs, next, d := fetch(start)
		start = next
		done = d

		for _, song := range songs {
			err := process(song)
			if err != nil {
				fmt.Printf("Error downloading %v\n", song.Name)
			}
		}
	}
}
