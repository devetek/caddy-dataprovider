package utils

import (
	"log"
	"net/http"
	"time"
)

func Fetcher(url string, payload string) (*http.Response, error) {
	var start = time.Now()
	var client = &http.Client{}

	// init request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// do request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(start)
	if payload == "" {
		log.Printf("Fetcher took %s", elapsed)
	}

	return response, nil
}
