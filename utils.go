package main

import (
	"encoding/json"
	"io"
	"os"
)

func getRedirectURLs() (redirects map[string]string, err error) {
	file, err := os.Open("./.data/redirects.json")
	if err != nil {
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return
	}

	json.Unmarshal(data, &redirects)
	return
}
