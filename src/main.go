package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gosuri/uiprogress"
	"github.com/ttacon/chalk"
)

var (
	sha1ver = "unknown"
	gitTag  string

	printHelp   bool
	printVer    bool
	noSubdir    bool
	concurrency = uint(runtime.NumCPU())
	args        []string

	baseDir      = "."
	client       *http.Client
	userAgent    string
	beatSaverURL = "https://beatsaver.com"
)

func main() {
	SetDetails("BeatSaver Downloader "+gitTag, "https://github.com/lolPants/beatsaver-downloader")
	SetExample("./beatsaver-downloader [dir]")

	RegisterBoolFlag(&printVer, "v", "version", "Print version information.")
	RegisterUintFlag(&concurrency, "c", "concurrency", "Max number of jobs allowed to run at a time.")
	RegisterBoolFlag(&noSubdir, "n", "no-subdir", "Don't create CustomSongs/ subdirectory.")
	ParseFlags(&args)

	if len(os.Args[1:]) == 0 {
		PrintUsageAndExit()
		return
	}

	if printVer == true {
		if gitTag != "" {
			fmt.Println(gitTag)
		}

		fmt.Println(sha1ver)
		return
	}

	if printHelp == true {
		PrintUsageAndExit()
		return
	}

	if concurrency < 1 {
		printError(chalk.Bold.TextStyle("--concurrency") + " cannot be less than 1!")
	}

	if len(args) == 0 {
		PrintUsageAndExit()
		return
	}

	if noSubdir == true {
		baseDir = filepath.Join(args[0])
	} else {
		baseDir = filepath.Join(args[0], "CustomSongs")
	}

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
