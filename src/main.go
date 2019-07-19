package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uiprogress"
)

var (
	sha1ver = "unknown"
	gitTag  string

	client       *http.Client
	userAgent    string
	beatSaverURL = "https://beatsaver.com"
)

func main() {
	userAgent = fmt.Sprintf("beatsaver-downloader/%v", sha1ver)
	client = &http.Client{
		Timeout: time.Second * 5,
	}

	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Writer = os.Stderr
	s.Suffix = " Downloading latest dump..."
	s.Start()

	resp, err := readDump()
	if err != nil {
		s.FinalMSG = "\t\t\t\t\t\n"
		s.Stop()

		panic(err)
	}

	s.FinalMSG = "Complete!\t\t\t\t\t\t\n"
	s.Stop()
	uiprogress.Start()

	count := len(resp)
	bar := uiprogress.AddBar(count).AppendCompleted()
	bar.PrependFunc(func(b *uiprogress.Bar) string {
		return fmt.Sprintf("%v / %v", b.Current(), count)
	})

	concurrency := runtime.NumCPU()
	sem := make(chan bool, concurrency)

	for _, bmap := range resp {
		sem <- true
		go func(bmap BeatmapInfo) {
			defer func() { <-sem }()

			saveMap(bmap)
			bar.Incr()
		}(bmap)
	}

	for i := 0; i < cap(sem); i++ {
		sem <- true
	}
}
