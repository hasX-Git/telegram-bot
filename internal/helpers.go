package Internal

import (
	"errors"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	tu "github.com/mymmrac/telego/telegoutil"
	"gorm.io/gorm"
)

var fileHash sync.Map
var acc sync.Map

func findFile(ctx *th.Context, update telego.Update) *telego.SendDocumentParams {

	chatID := update.Message.Chat.ID
	hash := update.Message.Text
	var dbFile File

	result := DB.First(&dbFile, "hash = ?", hash)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "no such file in database"))
		} else {
			_, _ = ctx.Bot().SendMessage(ctx, tu.Message(tu.ID(chatID), "error occured when searching for file"))
		}
		return nil
	}

	filepath := "./files/" + dbFile.Filename

	file, err := os.Open(filepath)
	if err != nil {
		log.Println("error")
		return nil
	}

	document := tu.Document(
		tu.ID(chatID),
		tu.File(file),
	)

	return document
}

func findAcc(update telego.Update) string {
	AID := update.Message.Text
	var acc Account

	result := DB.Preload("PersonInfo").Preload("Trs").First(&acc, "aid = ?", AID)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "no such file in database"
		} else {
			return "error occured when searching for file"
		}
	}

	var accinfo string

	accinfo +=
		"First Name: " + acc.PersonInfo.Firstn + "\n" +
			"Last Name: " + acc.PersonInfo.Lastn + "\n" +
			"National ID: " + acc.PersonInfo.NID + "\n" +
			"Account ID " + acc.AID + "\n" +
			"Balance: " + strconv.Itoa(int(acc.Balance)) + "\n" +
			"Transactions:\n"

	for _, tr := range acc.Trs {
		accinfo +=
			"\tSum: " + strconv.Itoa(int(tr.Sum)) + "\n" +
				"\tTransaction ID: " + tr.TrID + "\n" +
				"\n"
	}

	return accinfo

}
