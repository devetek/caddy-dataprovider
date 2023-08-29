package utils

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func Fetcher(method string, url string, payload string) (*http.Response, error) {
	var start = time.Now()
	var client = &http.Client{}
	var body *bytes.Buffer = nil

	// set payload if not empty
	if payload != "" {
		body = bytes.NewBuffer([]byte(payload))
	}

	// init request
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// do request
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	elapsed := time.Since(start)
	log.Printf("Fetcher took %s", elapsed)

	return response, nil
}
