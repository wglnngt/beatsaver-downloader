package main

// BeatmapInfo BeatSaver Beatmap Info
type BeatmapInfo struct {
	Key         string   `json:"key"`
	Name        string   `json:"name"`
	Hash        string   `json:"hash"`
	DownloadURL string   `json:"downloadURL"`
	Metadata    Metadata `json:"metadata"`
}

// Metadata Beatmap Metadata
type Metadata struct {
	SongName        string `json:"songName"`
	SongSubName     string `json:"songSubName"`
	SongAuthorName  string `json:"songAuthorName"`
	LevelAuthorName string `json:"levelAuthorName"`
}
