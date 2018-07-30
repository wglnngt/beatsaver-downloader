package main

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
			process(song)
		}
	}
}
