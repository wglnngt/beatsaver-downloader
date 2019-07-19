package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func readDump() ([]BeatmapInfo, error) {
	url := fmt.Sprintf("%v/api/download/dump/maps", beatSaverURL)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := make([]BeatmapInfo, 0)
	json.Unmarshal(body, &response)
	return response, nil
}
