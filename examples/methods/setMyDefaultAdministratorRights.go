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

	message := tg.NewSetMyDefaultAdministratorRights()

	_, err = tg.SetMyDefaultAdministratorRights(message)
	if err != nil {
		log.Fatal(err)
	}
}
