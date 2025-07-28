package postgres

import (
	"errors"
	"log"
	"os"
	u "tg-bot/internal/domain"

	"gorm.io/gorm"
)

func SELECTfile(hash string) (*os.File, error) {
	var dbFile u.File

	result := DB.First(&dbFile, "hash = ?", hash)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("no such file in database")
		} else {
			log.Println("error occured when searching for file")
		}
		return nil, nil
	}

	filepath := "./files/" + dbFile.Filename

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("Could not open file")
		return nil, nil
	}

	return file, nil
}

func INSERTfile(file []byte, filename string) error {

	filepath := "./files/" + filename

	err := os.WriteFile(filepath, file, 0644)
	if err != nil {
		log.Println("error occured when downloading file")
		return err
	}

	var newFile u.File

	newFile.Filename = filename
	newFile.Hash = hash(filename)

	result := DB.Create(&newFile)

	if result.Error != nil {
		log.Println("error occured when updating file db")
		return result.Error
	}

	return nil
}

func SELECTacc(aid string) (*u.Account, error) {
	var acc u.Account

	result := DB.Preload("PersonInfo").Preload("Trs").First(&acc, "aid = ?", aid)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		} else {
			return nil, result.Error
		}
	}

	return &acc, nil
}
