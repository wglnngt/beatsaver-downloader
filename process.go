package main

import (
	"fmt"
	"sort"
)

func process(song Song) {
	// Check for song existing
	i := sort.SearchStrings(existing, song.Key)
	if i < len(existing) && existing[i] == song.Key {
		fmt.Printf("%v already downloaded. Skipping...\n", song.Name)
		return
	}

	fmt.Printf("Downloading %v\n", song.Name)
	download(song.DownloadURL, song.Key, song.SHA1)
}
