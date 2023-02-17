package openai

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	gogpt "github.com/sashabaranov/go-gpt3"
)

type OpenAiBot struct {
	client *gogpt.Client
}

func NewOpenAiBot(token string) *OpenAiBot {
	client := gogpt.NewClient(token)
	return &OpenAiBot{
		client: client,
	}
}

func (bot *OpenAiBot) setRequest(msg string, stream bool) gogpt.CompletionRequest {
	return gogpt.CompletionRequest{
		Model:     gogpt.GPT3TextDavinci003,
		MaxTokens: 2000,
		Prompt:    msg,
		Stream:    stream,
	}
}

func (bot *OpenAiBot) Send(ctx context.Context, msg string) (gogpt.CompletionResponse, error) {
	req := bot.setRequest(msg, false)
	response, err := bot.client.CreateCompletion(ctx, req)
	if err != nil {
		log.Fatal("Send OpenAiBot error: ", err)
	}
	return response, nil
}

func (bot *OpenAiBot) AquireStream(ctx context.Context, msg string) *gogpt.CompletionStream {
	req := bot.setRequest(msg, true)
	stream, err := bot.client.CreateCompletionStream(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	return stream
}

func (bot *OpenAiBot) CompileMessage(stream *gogpt.CompletionStream) string {
	defer stream.Close()
	message := ""
	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			fmt.Println("Stream finished")
			return message
		}

		if err != nil {
			fmt.Printf("Stream error: %v\n", err)
			return "error"
		}

		for chat := range response.Choices {
			message = message + response.Choices[chat].Text
		}

	}
}
