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

	msg := tg.NewSetChatStickerSet()
	msg.Username = "username"
	msg.StickerSetName = "name"

	_, err = tg.SetChatStickerSet(msg)
	if err != nil {
		log.Fatal(err)
	}
}
