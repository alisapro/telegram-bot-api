package main

import (
	"log"

	"github.com/alisapro/telegram-bot-api/telegram"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	message := tg.NewStopPoll()
	message.ChatID = 1234
	message.MessageID = 1234

	_, err = tg.StopPoll(message)
	if err != nil {
		log.Fatal(err)
	}
}
