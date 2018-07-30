package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func fetch(start int) ([]Song, int, bool) {
	url := fmt.Sprintf("https://beatsaver.com/api/songs/new/%v", start)

	bsClient := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "beatsaver-downloader")

	res, getErr := bsClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	response := BeatSaver{}
	jsonErr := json.Unmarshal(body, &response)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	num := len(response.Songs)
	next := num + start
	done := num == 0

	return response.Songs, next, done
}
