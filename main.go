package main

import (
	"fmt"
)

func main() {
	start := 0
	done := false

	for !done {
		songs, next, d := fetch(start)
		start = next
		done = d

		fmt.Printf("%v", songs)
	}
}
