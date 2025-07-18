package Helpers

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"google.golang.org/genai"
)

var (
	Ctx     context.Context
	Bot     *telego.Bot
	Updates <-chan telego.Update
	Client  *genai.Client
	Bh      *th.BotHandler
)

func InitServices() {
	var err error

	Ctx = context.Background()

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	botToken := os.Getenv("TOKEN")
	apiKey := os.Getenv("GEMINI_API_KEY")

	Bot, err = telego.NewBot(botToken, telego.WithDefaultDebugLogger())
	if err != nil {
		log.Fatal("error with bot telegoBot")
	}

	Updates, err = Bot.UpdatesViaLongPolling(Ctx, nil)
	if err != nil {
		log.Fatal("error with updates")
	}
	Bh, err = th.NewBotHandler(Bot, Updates)
	if err != nil {
		log.Fatal("error with bh")
	}

	Client, err = genai.NewClient(Ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal("error with genai client")
	}

}
