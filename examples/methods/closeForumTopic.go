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

	msg := tg.NewCloseForumTopic()
	msg.Username = "username"
	msg.MessageThreadID = 1234

	_, err = tg.CloseForumTopic(msg)
	if err != nil {
		log.Fatal(err)
	}
}
