package main

// Song BeatSaver Song
type Song struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	DownloadURL string `json:"downloadUrl"`
	MD5         string `json:"hashMd5"`
}

// BeatSaver Response Object
type BeatSaver struct {
	Total int    `json:"total"`
	Songs []Song `json:"songs"`
}
