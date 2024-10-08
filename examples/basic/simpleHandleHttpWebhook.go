package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alisapro/telegram-bot-api/telegram"
)

func main() {
	fmt.Println("start at port:", "BotPortNumber")

	err := http.ListenAndServe("BotPortNumber", http.HandlerFunc(handleWebhook))
	if err != nil {
		log.Fatal(err)
	}
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	update, err := telegram.HandleUpdate(r)
	if err != nil {
		err = telegram.HandleUpdateError(w, err)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if update.Message != nil {
		fmt.Println(update.Message.Text)
	}
}
