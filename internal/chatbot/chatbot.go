package chatbot

import (
	"context"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rtm-se/chatgptbot/internal/openai"
)

type ChatBot struct {
	Token   string
	Bot     *tgbotapi.BotAPI
	OpenBot *openai.OpenAiBot
}

func NewChatBot(tgken string, openToken string) (*ChatBot, error) {
	bot, err := tgbotapi.NewBotAPI(tgken)
	if err != nil {
		return nil, err
	}
	oc := openai.NewOpenAiBot(openToken)
	return &ChatBot{
		Token:   tgken,
		Bot:     bot,
		OpenBot: oc,
	}, nil
}

func (b *ChatBot) HandleMessage(ctx context.Context, update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	resp, err := b.OpenBot.Send(ctx, update.Message.Text)
	if err != nil {
		fmt.Printf("error sending openai message: %v", err)
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, resp.Choices[0].Text)
	_, er := b.Bot.Send(msg)
	if er != nil {
		log.Panicf("Error sending message %v", err)
	}

}
