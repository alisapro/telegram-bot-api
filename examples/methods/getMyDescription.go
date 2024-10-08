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

	getDescription := tg.NewGetMyDescription()

	_, err = tg.GetMyDescription(getDescription)
	if err != nil {
		log.Fatal(err)
	}
}
