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

	msg := tg.NewGetChatMember()
	msg.Username = "username"
	msg.UserID = 1234

	_, err = tg.GetChatMember(msg)
	if err != nil {
		log.Fatal(err)
	}
}
