package telegram

import (
	"bytes"
	"net/url"
	"strconv"
)

// Send the message to Telegram chat using BotConfig
func SendMessage(config BotConfig, message string, parseMode string) string {
	apiMethod := "sendMessage"

	uri := getTelegramUri(config, apiMethod)

	params := setMessageParams(uri, strconv.Itoa(config.ChatId), message, parseMode)

	buffer := new(bytes.Buffer)
	buffer.WriteString(params.Encode())

	resData := doRequest(uri, buffer, "application/x-www-form-urlencoded")

	return string(resData)
}

// Set GET params for API URI
func setMessageParams(uri string, chat_id string, message string, parseMode string) (params url.Values) {
	params = url.Values{}
	params.Set("chat_id", chat_id)
	params.Set("parse_mode", parseMode)
	params.Set("text", message)

	return
}
