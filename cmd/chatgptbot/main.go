package main

import (
	"context"
	"fmt"
	"log"

	"github.com/rtm-se/chatgptbot/internal/chatbot"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	openaitoekn = ""
	tgbottoken  = ""
)

func main() {
	fmt.Println("It works!")

	bot, err := chatbot.NewChatBot(tgbottoken, openaitoekn)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.Bot.GetUpdatesChan(u)

	if err != nil {
		log.Fatal(err)
	}
	for update := range updates {
		bot.HandleMessage(ctx, update)
	}
}
