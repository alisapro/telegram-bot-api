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

	msg := tg.NewSendVideo()
	msg.ChatID = 1234
	msg.Video = tg.FileURL("url")

	_, err = tg.SendVideo(msg)
	if err != nil {
		log.Fatal(err)
	}
}
