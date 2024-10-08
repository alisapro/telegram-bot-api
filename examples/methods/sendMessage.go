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

	msg := tg.NewSendMessage()
	msg.ChatID = 1234
	msg.ParseMode = tg.ModeMarkdown()
	msg.Text = "some text"

	_, err = tg.SendMessage(msg)
	if err != nil {
		log.Fatal(err)
	}
}
