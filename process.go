package main

import (
	"fmt"
	"sort"
)

func process(song Song) error {
	// Check for song existing
	i := sort.SearchStrings(existing, song.Key)
	if i < len(existing) && existing[i] == song.Key {
		fmt.Printf("[SKIPPING]    %v\n", song.Name)
		return nil
	}

	fmt.Printf("[DOWNLOADING] %v\n", song.Name)
	err := download(song.DownloadURL, song.Key, song.SHA1)
	if err != nil {
		return err
	}

	return nil
}
