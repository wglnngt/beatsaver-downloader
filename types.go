package main

// Song BeatSaver Song
type Song struct {
	ID          int    `json:"id"`
	Key         string `json:"key"`
	Name        string `json:"name"`
	DownloadURL string `json:"downloadUrl"`
	SHA1        string `json:"hashSha1"`
}

// BeatSaver Response Object
type BeatSaver struct {
	Total int    `json:"total"`
	Songs []Song `json:"songs"`
}

// BeatSaverAlt Response Object
type BeatSaverAlt struct {
	Total int             `json:"total"`
	Songs map[string]Song `json:"songs"`
}
