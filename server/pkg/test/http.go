package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

const testUrl = "api.test.com"

func PostRequest(jsonRequest map[string]interface{}) (*http.Request, error) {
	jsonBytes, err := json.Marshal(jsonRequest)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodPost, testUrl, bytes.NewReader(jsonBytes))
	return request, err
}

func InvalidPostRequest() (*http.Request, error) {
	return http.NewRequest(http.MethodPost, testUrl, strings.NewReader("notjson"))
}
