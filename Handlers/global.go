package Handler

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"google.golang.org/genai"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Ctx     context.Context
	Bot     *telego.Bot
	Updates <-chan telego.Update
	Bh      *th.BotHandler

	Client *genai.Client

	DB *gorm.DB
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

func ConnectToDB() {

	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	for i := 0; i < 10; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			log.Println("Connected")
			break
		}
		log.Println("Reconnecting...")
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatal("Conneciton to database failed")
	}

	if err = DB.AutoMigrate(&Account{}, &ClientInfo{}, &Transaction{}, &File{}); err != nil {
		log.Fatal("migration failed")
	}

}
