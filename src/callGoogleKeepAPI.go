package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"net/http"
)

func main() {
	// GoogleKeep API URL
	GOOGLE_KEEP_API_URL := "https://keep.googleapis.com/v1/notes"

	payload := map[string]interface{} {
		"title": "test",
		"body": map[string]interface{} {
			"text": map[string]interface{} {
				"text": "sample",
			},
		},
	}

	json, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error occured in encoding json.")
	}
	body:= bytes.NewReader(json)

	res, err := http.Post(GOOGLE_KEEP_API_URL, "application/json", body)
	if err != nil {
		fmt.Println("Error occured in API request.")
	}
	fmt.Println("response status is " + res.Status)
}