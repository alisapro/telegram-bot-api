package main

import (
	"log"

	"github.com/alisapro/telegram-bot-api/telegram"
	"github.com/alisapro/telegram-bot-api/types"
)

func main() {
	tg, err := telegram.New("BotToken")
	if err != nil {
		log.Fatal(err)
	}

	msg := tg.NewSetMessageReaction()
	msg.ChatID = 1234
	msg.MessageID = 1235

	reaction := types.ReactionType{
		ReactionTypeEmoji: types.ReactionTypeEmoji{
			Type:  "emoji",
			Emoji: "❤",
		},
		ReactionTypeCustomEmoji: types.ReactionTypeCustomEmoji{
			Type:        "custom_emoji",
			CustomEmoji: "😀",
		},
	}
	msg.Reaction = append(msg.Reaction, reaction)

	_, err = tg.SetMessageReaction(msg)
	if err != nil {
		log.Fatal(err)
	}
}
