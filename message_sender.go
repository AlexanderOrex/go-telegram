package telegram

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const API_TOKEN_MASK = "<token>"
const API_METHOD_MASK = "<method>"

const API_URI = "https://api.telegram.org/bot" + API_TOKEN_MASK + "/" + API_METHOD_MASK

func SendMessage(config BotConfig, message string) error {
	apiMethod := "sendMessage"

	uri := strings.Replace(API_URI, API_METHOD_MASK, apiMethod, 1)
	uri = strings.Replace(uri, API_TOKEN_MASK, config.Token, 1)

	buffer := new(bytes.Buffer)
	params := url.Values{}
	params.Set("chat_id", strconv.Itoa(config.ChatId))
	params.Set("text", message)
	fmt.Println(params)
	buffer.WriteString(params.Encode())
	fmt.Println(uri)
	req, _ := http.NewRequest("POST", uri, buffer)
	req.Header.Set("content-type", "application/x-www-form-urlencoded")

	client := &http.Client{}

	res, err := client.Do(req)
	req.Body.Close()

	fmt.Println(res)

	return err
}
