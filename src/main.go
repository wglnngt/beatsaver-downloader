package main

import (
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/schollz/progressbar"
)

var (
	sha1ver = "unknown"
	gitTag  string

	beatSaverURL = "https://beatsaver.com"
)

func main() {
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

	bar := progressbar.New(len(resp))
	for i := 0; i < len(resp); i++ {
		bar.Add(1)
		time.Sleep(time.Millisecond)
	}
}
