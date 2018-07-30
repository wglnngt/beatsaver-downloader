package main

import (
	"fmt"
)

func main() {
	createDir()
	loadExisting()

	fmt.Printf("%v", existing)
}

func downloadAll() {
	start := 0
	done := false

	for !done {
		songs, next, d := fetch(start)
		start = next
		done = d

		fmt.Printf("%v", songs)
	}
}
