package telegram

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const API_TOKEN_MASK = "<token>"
const API_METHOD_MASK = "<method>"

const API_URI = "https://api.telegram.org/bot" + API_TOKEN_MASK + "/" + API_METHOD_MASK

// Generate URI for API with token and method
func getTelegramUri(config BotConfig, apiMethod string) string {
	uri := strings.Replace(API_URI, API_METHOD_MASK, apiMethod, 1)
	uri = strings.Replace(uri, API_TOKEN_MASK, config.Token, 1)

	return uri
}

// Do a HTTP request
func doRequest(uri string, buffer io.Reader, requestType string) []byte {
	req, err := http.NewRequest("POST", uri, buffer)

	if err != nil {
		log.Fatal("Failed to create http.NewRequest", err)
	}

	req.Header.Set("content-type", requestType)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		log.Fatal("Failed to do request", err)
	}

	var bodyContent []byte
	res.Body.Read(bodyContent)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal("Failed to do read res.Body", err)
	}

	return body
}
