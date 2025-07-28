package bot

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
)

var (
	Ctx     context.Context
	Bot     *telego.Bot
	Updates <-chan telego.Update
	Bh      *th.BotHandler
)

func InitTGBot() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	botToken := os.Getenv("TOKEN")

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
}
