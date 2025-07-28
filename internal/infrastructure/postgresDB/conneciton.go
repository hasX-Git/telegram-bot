package postgres

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgresDB() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(".env file not found")
	}

	log.Println("bank db connection")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_SSLMODE"),
	)

	for i := 0; i < 10; i++ {
		if db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{}); err != nil {
			return db, nil
		}

		log.Println("Reconnecting...")
		time.Sleep(1 * time.Second)
	}
	// if err = DB.AutoMigrate(&u.Account{}, &u.Client{}, &u.Transaction{}, &u.File{}); err != nil {
	// 	log.Fatal("migration failed")
	// }

	return nil, errors.New("connection failer")
}
