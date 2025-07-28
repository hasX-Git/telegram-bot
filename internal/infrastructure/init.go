package infrastructure

import (
	ai "tg-bot/internal/infrastructure/googleAI"
	db "tg-bot/internal/infrastructure/postgresDB"
	tgb "tg-bot/internal/infrastructure/telegoBot"
)

func Init() {
	ai.InitAI()
	db.ConnectToPostgresDB()
	tgb.InitTGBot()
}
