package telegrambot

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

var (
	Ctx    context.Context
	Client *genai.Client
)

const Prompt = "If the user seems to ask for some command, offer \"/help\" command ONLY. Otherwise, just respond without offering the command\n"
const Aimodel = "gemini-2.5-flash"

func InitAI() {
	var err error

	Ctx = context.Background()

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	apiKey := os.Getenv("GEMINI_API_KEY")

	Client, err = genai.NewClient(Ctx, &genai.ClientConfig{
		APIKey:  apiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal("error with genai client")
	}

}
