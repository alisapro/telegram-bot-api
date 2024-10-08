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

	getCommands := tg.NewGetMyCommandsWithScope(tg.NewBotCommandScopeDefault())

	_, err = tg.GetMyCommands(getCommands)
	if err != nil {
		log.Fatal(err)
	}
}
