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

	msg := tg.NewSendVideoNote()
	msg.ChatID = 1234
	msg.VideoNote = tg.FileURL("url")

	_, err = tg.SendVideoNote(msg)
	if err != nil {
		log.Fatal(err)
	}
}
