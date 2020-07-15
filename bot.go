package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UpdateT struct {
	Ok     bool            `json:"ok"`
	Result []UpdateResultT `json:"result"`
}

type UpdateResultT struct {
	UpdateId int                  `json:"update_id"`
	Message  UpdateResultMessageT `json:"message"`
}

type UpdateResultMessageT struct {
	MessageId int               `json:"message_id"`
	From      UpdateResultFromT `json:"from"`
	Chat      UpdateResultChatT `json:"chat"`
	Date      int               `json:"date"`
	Text      string            `json:"text"`
}

type UpdateResultFromT struct {
	Id        int    `json:"id"`
	IsBot     bool   `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Language  string `json:"language_code"`
}

type UpdateResultChatT struct {
	Id        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type SendMessageResponseT struct {
	Ok     bool                 `json:"ok"`
	Result UpdateResultMessageT `json:"result"`
}

const baseTelegramUrl = "https://api.telegram.org"
const telegramToken = "1265259222:AAHKgdF1ZuRyv7u-X0afNiWnl6JyGDdJFEQ"
const getUpdatesUri = "getUpdates"
const sendMessageUrl = "sendMessage"

const keywordStart = "/start"

func main() {

	update, err := getUpdates()
	if err != nil {
		fmt.Println(err.Error())

		return
	}

	for _, item := range update.Result {
		if item.Message.Text == "Привет" {
			text := "Привет, " + item.Message.From.FirstName + " " + item.Message.From.LastName
			sendMessage(item.Message.Chat.Id, text)
		}
	}
}

func getUpdates() (UpdateT, error) {
	url := baseTelegramUrl + "/bot" + telegramToken + "/" + getUpdatesUri
	response := getResponse(url)

	update := UpdateT{}
	err := json.Unmarshal(response, &update)
	if err != nil {
		return update, err
	}

	return update, nil
}

func sendMessage(chatId int, text string) (SendMessageResponseT, error) {
	url := baseTelegramUrl + "/bot" + telegramToken + "/" + sendMessageUrl
	url = url + "?chat_id=" + strconv.Itoa(chatId) + "&text=" + text
	response := getResponse(url)

	sendMessage := SendMessageResponseT{}
	err := json.Unmarshal(response, &sendMessage)
	if err != nil {
		return sendMessage, err
	}

	return sendMessage, nil
}

func getResponse(url string) []byte {
	response := make([]byte, 0)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

		return response
	}

	defer resp.Body.Close()

	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		response = append(response, bs[:n]...)

		if n == 0 || err != nil {
			break
		}
	}

	return response
}
